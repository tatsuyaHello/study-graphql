package graphql_util

import (
	"github.com/graphql-go/graphql"
	"github.com/tatsuyaHello/study-graphql/graphQL_sample/application/graphql_util/fields"
)

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"user":      fields.UserField,
		"userList":  fields.UserListField,
		"event":     fields.EventField,
		"eventList": fields.EventListField,
	},
})
