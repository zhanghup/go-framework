package test

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	a := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(a)
}
func TestDate(t *testing.T) {
	a := time.Now().Format("2006-01-02")
	fmt.Println(a)
}
func TestMonth(t *testing.T) {
	a := time.Now().Format("2006-01")
	fmt.Println(a)
}
