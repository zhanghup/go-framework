package gql

//type Query map[string]*Field
//type Mutation map[string]*Field
//type Subscription map[string]*Field
//type Graphql struct {
//	Query        Query
//	Mutation     Mutation
//	Subscription Subscription
//}
//
//type Object struct {
//	FuncName        string             `json:"name"`
//	Interfaces      interface{}        `json:"interfaces"`
//	Fields          Fields             `json:"fields"`
//	IsTypeOf        graphql.IsTypeOfFn `json:"isTypeOf"`
//	FuncDescription string             `json:"description"`
//
//	graphqlObject *graphql.Object
//}
//
//func (this *Object) instance() graphql.Output {
//	this.check()
//	return this.graphqlObject
//}
//
//func (this *Object) check() {
//	if this.graphqlObject != nil {
//		return
//	}
//	fields := graphql.Fields{}
//	for k, v := range this.Fields {
//		fields[k] = v.getField()
//	}
//
//	this.graphqlObject = graphql.NewObject(graphql.ObjectConfig{
//		Name:        this.FuncName,
//		Interfaces:  this.Interfaces,
//		Fields:      fields,
//		IsTypeOf:    this.IsTypeOf,
//		Description: this.FuncDescription,
//	})
//}
