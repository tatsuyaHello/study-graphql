package service

import (
	"github.com/tatsuyaHello/study-graphql/graphQL_sample/domain/model/user"
	"github.com/tatsuyaHello/study-graphql/graphQL_sample/infrastructure"
)

func FindUserById(userId string) (*user.User, error) {
	repository := infrastructure.NewUserRepository()
	return repository.FindById(userId)
}
