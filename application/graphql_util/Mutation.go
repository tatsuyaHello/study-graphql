package graphql_util

import (
	"github.com/graphql-go/graphql"
	"github.com/tatsuyaHello/study-graphql/graphQL_sample/application/graphql_util/fields"
)

var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"createUser":  fields.CreateUserField,
		"createEvent": fields.CreateEventField,
	},
})
