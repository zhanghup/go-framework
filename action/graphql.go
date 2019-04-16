package action

import (
	"github.com/gin-gonic/gin"
	"github.com/zhanghup/go-framework/action/graphql"
)

func RegisterGql(g *graphql.Graphql) gin.HandlerFunc {
	return graphql.Register(g)
}
