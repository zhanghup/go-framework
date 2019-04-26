package gql

import "github.com/graphql-go/graphql"

type output interface {
	output() graphql.Output
}
