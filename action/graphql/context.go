package graphql

import "github.com/graphql-go/graphql"

type Graphql struct {
	Query        map[string]*Field
	Mutation     map[string]*Field
	Subscription map[string]*Field
}

type Config struct {
	Name        string      `json:"name"`
	Interfaces  interface{} `json:"interfaces"`
	Fields      interface{} `json:"fields"`
	Description string      `json:"description"`
}
type FieldResolve graphql.FieldResolveFn
type FieldArgument graphql.FieldConfigArgument
type Field struct {
	Name              string         `json:"name"` // used by graphlql-relay
	Type              graphql.Output `json:"type"`
	Args              FieldArgument  `json:"args"`
	Resolve           FieldResolve   `json:"-"`
	DeprecationReason string         `json:"deprecationReason"`
	Description       string         `json:"description"`
	field             *graphql.Field `json:"-"`
}

func (this *Field) getField() *graphql.Field {
	if this.field != nil {
		return this.field
	}
	this.field = &graphql.Field{
		Name:              this.Name,
		Type:              this.Type,
		Args:              graphql.FieldConfigArgument(this.Args),
		Resolve:           graphql.FieldResolveFn(this.Resolve),
		DeprecationReason: this.DeprecationReason,
		Description:       this.Description,
	}
	return this.field
}
