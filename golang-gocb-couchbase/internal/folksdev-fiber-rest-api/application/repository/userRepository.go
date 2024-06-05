package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/couchbase/gocb/v2"
	"golang-gocb-couchbase/configuration"
	"golang-gocb-couchbase/internal/folksdev-fiber-rest-api/domain"
	"golang-gocb-couchbase/internal/folksdev-fiber-rest-api/pkg/utils"
	"strings"
	"time"
)

type IUserRepository interface {
	Upsert(ctx context.Context, user *domain.User, ttlInDays int) error
	GetById(ctx context.Context, id string) (*domain.User, error)
	Get(ctx context.Context) ([]*domain.User, error)
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
}

var couchbaseDefaultTimeoutDuration = time.Second * 5

type userRepository struct {
	FolksdevCluster    *gocb.Cluster
	FolksdevUserBucket *gocb.Bucket
	expirationUtil     utils.IExpirationUtil
}

func NewUserRepository(cluster *gocb.Cluster, expirationUtil utils.IExpirationUtil) IUserRepository {
	return &userRepository{
		FolksdevCluster:    cluster,
		FolksdevUserBucket: cluster.Bucket(configuration.FolksdevUserBucket),
		expirationUtil:     expirationUtil,
	}
}

func (r *userRepository) Upsert(ctx context.Context, user *domain.User, ttlInDays int) error {

	_, err := r.FolksdevUserBucket.DefaultCollection().Upsert(user.Id,
		user,
		&gocb.UpsertOptions{
			Expiry: r.expirationUtil.GetTtl(false, ttlInDays),
		},
	)

	if err != nil {
		fmt.Printf("ctx: %#v - userRepository.Upsert ERROR: %#v\n", ctx, err.Error())
		return errors.New("INTERNAL SERVER ERROR")
	}

	return nil
}

func (r *userRepository) GetById(ctx context.Context, id string) (*domain.User, error) {
	var user domain.User

	queryResult, err := r.FolksdevUserBucket.DefaultCollection().Get(id,
		&gocb.GetOptions{Timeout: time.Second * 1})

	if err != nil {
		if errors.Is(err, gocb.ErrDocumentNotFound) {
			fmt.Printf("ctx: %#v - userRepository.GetById INFO: %#v\n", ctx, err.Error())
			return nil, nil
		}

		fmt.Printf("ctx: %#v - userRepository.GetById ERROR: %#v\n", ctx, err.Error())
		return nil, errors.New(fmt.Sprintf("There was an error while getting userId: %s - ERROR: %#v", id, err.Error()))
	}

	if err = queryResult.Content(&user); err != nil {
		fmt.Printf("ctx: %#v - userRepository.GetById Content ERROR: %#v\n", ctx, err.Error())
		return nil, errors.New("INTERNAL SERVER ERROR")
	}

	return &user, nil
}

func (r *userRepository) Get(ctx context.Context) ([]*domain.User, error) {

	queryStr := strings.ReplaceAll("SELECT u.* FROM `{{bucket}}` u", "{{bucket}}", configuration.FolksdevUserBucket)

	fmt.Printf("ctx: %#v - userRepository.Get build queryStr: %s\n", ctx, queryStr)

	queryResult, err := r.FolksdevCluster.Query(queryStr,
		&gocb.QueryOptions{
			Timeout:  couchbaseDefaultTimeoutDuration,
			Adhoc:    true,
			Readonly: true,
		})

	if err != nil {
		if errors.Is(err, gocb.ErrDocumentNotFound) {
			fmt.Printf("ctx: %#v - userRepository.Get INFO: Document not found ERROR: %#v\n", ctx, err.Error())
			return nil, nil
		}

		fmt.Printf("ctx: %#v - userRepository.Get ERROR: %#v\n", ctx, err.Error())
		return nil, errors.New("INTERNAL SERVER ERROR")
	}

	defer queryResult.Close()

	var users []*domain.User

	for queryResult.Next() {
		var user domain.User
		if err := queryResult.Row(&user); err == nil {
			users = append(users, &user)
		}
	}

	if len(users) == 0 {
		fmt.Printf("ctx: %#v - userRepository.Get INFO: Not Found Data\n", ctx)
		return nil, nil
	}

	return users, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	queryStr := strings.ReplaceAll("SELECT u.* FROM `{{bucket}}` u WHERE u.email = $email LIMIT 1",
		"{{bucket}}",
		configuration.FolksdevUserBucket)

	fmt.Printf("ctx: %#v - userRepository.GetByEmail INFO: Build query: %s\n", ctx, queryStr)

	queryResult, err := r.FolksdevCluster.Query(queryStr,
		&gocb.QueryOptions{
			Timeout:         couchbaseDefaultTimeoutDuration,
			Adhoc:           true,
			Readonly:        true,
			NamedParameters: map[string]interface{}{"email": email},
		})

	if err != nil {
		if errors.Is(err, gocb.ErrDocumentNotFound) {
			fmt.Printf("ctx: %#v - userRepository.GetByEmail INFO: Document not found ERROR: %#v\n", ctx, err.Error())
			return nil, nil
		}

		fmt.Printf("ctx: %#v - userRepository.GetByEmail ERROR: %#v\n", ctx, err.Error())
		return nil, errors.New("INTERNAL SERVER ERROR")
	}

	defer queryResult.Close()

	var user *domain.User

	for queryResult.Next() {
		err = queryResult.Row(&user)
		if err != nil {
			fmt.Printf("ctx: %#v - userRepository.GetByEmail QueryResult.Row ERROR: %#v\n", ctx, err)
			return nil, errors.New("INTERNAL SERVER ERROR")
		}
	}

	return user, nil
}
