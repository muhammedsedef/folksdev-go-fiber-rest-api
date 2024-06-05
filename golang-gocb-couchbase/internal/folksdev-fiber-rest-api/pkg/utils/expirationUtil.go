package utils

import "time"

//go:generate mockery --name=IExpirationUtil
type IExpirationUtil interface {
	GetTtl(isTestPackage bool, expectedTtlInDays int) time.Duration
}

type expirationUtil struct {
}

func NewExpirationUtil() IExpirationUtil {
	return &expirationUtil{}
}

func (util expirationUtil) GetTtl(isTestEntity bool, expectedTtlInDays int) time.Duration {
	if isTestEntity {
		return time.Hour * 3
	}
	return time.Hour * 24 * time.Duration(expectedTtlInDays)
}
