package service

import (
	"github.com/tatsuyaHello/study-graphql/graphQL_sample/domain/model/event"
	"github.com/tatsuyaHello/study-graphql/graphQL_sample/infrastructure"
)

func FindEventById(userId string) (*event.Event, error) {
	repository := infrastructure.NewEventRepository()
	return repository.FindById(userId)
}
