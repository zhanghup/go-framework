package graphql

import "github.com/graphql-go/graphql"

type FieldResolve func(p ResolveParams) (interface{}, error)
type FieldArgument map[string]*ArgumentConfig
type ArgumentConfig graphql.ArgumentConfig
type ResolveParams graphql.ResolveParams

type Field struct {
	Name              string         `json:"name"` // used by graphlql-relay
	Type              graphql.Output `json:"type"`
	Args              FieldArgument  `json:"args"`
	Resolve           FieldResolve   `json:"-"`
	DeprecationReason string         `json:"deprecationReason"`
	Description       string         `json:"description"`
	field             *graphql.Field `json:"-"` // graphql 原生属性
}

func (this *Field) newField() {
	if this.field != nil {
		return
	}
	this.field = new(graphql.Field)
	this.field.Name = this.Name
	this.field.DeprecationReason = this.DeprecationReason
	this.field.Description = this.Description
}

func (this *Field) parseArgument() {
	fieldArgument := graphql.FieldConfigArgument{}
	for k, v := range this.Args {
		av := graphql.ArgumentConfig(*v)
		fieldArgument[k] = &av
	}
	this.field.Args = fieldArgument
}

func (this *Field) parseResolve() {
	resolve := func(p graphql.ResolveParams) (interface{}, error) {
		return this.Resolve(ResolveParams(p))
	}
	this.field.Resolve = resolve
}

func (this *Field) getField() *graphql.Field {
	if this.field != nil {
		return this.field
	}
	this.newField()
	this.parseArgument()
	this.parseResolve()

	this.field.Type = this.Type

	return this.field
}
