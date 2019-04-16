package graphql

import (
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func Register(g *Graphql) gin.HandlerFunc {
	fn := func(field map[string]*Field, name string) *graphql.Object {
		fs := graphql.Fields{}
		for k, v := range field {
			fs[k] = v.getField()
		}
		return graphql.NewObject(graphql.ObjectConfig{
			Name:   name,
			Fields: fs,
		})
	}

	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query:        fn(g.Query, "Query"),
		Mutation:     fn(g.Mutation, "Mutation"),
		Subscription: fn(g.Subscription, "Subscription"),
	})

	h := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   true,
		Playground: true,
	})
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
