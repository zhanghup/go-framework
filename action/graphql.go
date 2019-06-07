package action

import (
	"github.com/zhanghup/go-framework/action/gql"
	"github.com/zhanghup/go-framework/pkg/gin"
)

func RegisterGql(g *gql.Graphql) gin.HandlerFunc {
	return gql.Register(g)
}
