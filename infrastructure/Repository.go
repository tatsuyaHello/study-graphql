package infrastructure

import (
	"github.com/tatsuyaHello/study-graphql/graphQL_sample/infrastructure/event"
	"github.com/tatsuyaHello/study-graphql/graphQL_sample/infrastructure/user"
)

var NewUserRepository func() user.UserRepository = user.NewUserRepositoryMem

func UseUserRepositoryMem() {
	NewUserRepository = user.NewUserRepositoryMem
}

var NewEventRepository func() event.EventRepository = event.NewEventRepositoryMem

func UseEventRepositoryMem() {
	NewEventRepository = event.NewEventRepositoryMem
}
