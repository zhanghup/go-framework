package action

import (
	"github.com/gin-gonic/gin"
	"github.com/zhanghup/go-framework/action/gql"
)

func RegisterGql(g *gql.Graphql) gin.HandlerFunc {
	return gql.Register(g)
}
