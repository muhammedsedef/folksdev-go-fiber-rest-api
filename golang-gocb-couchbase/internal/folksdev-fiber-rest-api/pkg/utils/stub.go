package utils

import "golang-gocb-couchbase/internal/folksdev-fiber-rest-api/domain"

func GetUserStub() []*domain.User {
	return []*domain.User{
		{
			Id:        "1",
			FirstName: "Muhammed",
			LastName:  "Sedef",
			Email:     "muhammetsedef34@gmail.com",
			Password:  "1234",
			Age:       26,
		},
		{
			Id:        "2",
			FirstName: "Çağrı",
			LastName:  "Dursun",
			Email:     "cagrı@gmail.com",
			Password:  "123456",
			Age:       35,
		},
		{
			Id:        "3",
			FirstName: "Ayşe",
			LastName:  "Test",
			Email:     "aysetest@gmail.com",
			Password:  "123456",
			Age:       11,
		},
	}
}
