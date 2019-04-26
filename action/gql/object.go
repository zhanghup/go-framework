package gql

import "github.com/graphql-go/graphql"

type Object struct {
	Name        string             `json:"name"`
	Interfaces  interface{}        `json:"interfaces"`
	Fields      map[string]*Field  `json:"fields"`
	IsTypeOf    graphql.IsTypeOfFn `json:"isTypeOf"`
	Description string             `json:"description"`

	self *graphql.Object
}

func (this *Object) output() graphql.Output {
	return graphql.NewObject(graphql.ObjectConfig{
		Name:        this.Name,
		Fields:      this.Fields,
		Interfaces:  this.Interfaces,
		IsTypeOf:    this.IsTypeOf,
		Description: this.Description,
	})
}
