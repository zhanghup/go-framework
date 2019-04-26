package gql

import "github.com/graphql-go/graphql"

type Field struct {
	Name              string                      `json:"name"` // used by graphlql-relay
	Type              output                      `json:"type"`
	Args              graphql.FieldConfigArgument `json:"args"`
	Resolve           graphql.FieldResolveFn      `json:"-"`
	DeprecationReason string                      `json:"deprecationReason"`
	Description       string                      `json:"description"`
	field             *graphql.Field              `json:"-"` // graphql 原生属性
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

	this.field.Args = this.Args
}

func (this *Field) parseResolve() {
	//resolve := func(p graphql.ResolveParams) (interface{}, error) {
	//	return this.Resolve(ResolveParams(p))
	//}
	this.field.Resolve = this.Resolve
}

func (this *Field) getField() *graphql.Field {
	if this.field != nil {
		return this.field
	}
	this.newField()
	this.parseArgument()
	this.parseResolve()
	this.field.Type = this.Type.output()

	return this.field
}
