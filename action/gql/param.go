package gql

import "github.com/graphql-go/graphql"

type ResolveParams graphql.ResolveParams
type FieldResolve func(p ResolveParams) (interface{}, error)
type Params map[string]*Param
type Param struct {
	Type         output     `json:"type"`
	DefaultValue interface{} `json:"defaultValue"`
	Description  string      `json:"description"`
	self         graphql.ArgumentConfig
}
