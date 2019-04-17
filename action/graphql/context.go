package graphql

import "github.com/graphql-go/graphql"

type Graphql struct {
	Query        map[string]*Field
	Mutation     map[string]*Field
	Subscription map[string]*Field
}

type Object struct {
	Name        string             `json:"name"`
	Interfaces  interface{}        `json:"interfaces"`
	Fields      map[string]*Field  `json:"fields"`
	IsTypeOf    graphql.IsTypeOfFn `json:"isTypeOf"`
	Description string             `json:"description"`
}

