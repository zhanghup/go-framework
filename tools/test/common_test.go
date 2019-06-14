package test

import (
	"fmt"
	"github.com/zhanghup/go-framework/tools"
	"testing"
)

func TestSha256(t *testing.T) {
	fmt.Println(tools.Password("Aa123456","jdk"))
}
