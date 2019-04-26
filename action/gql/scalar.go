package gql

import "github.com/graphql-go/graphql"

type Scalar graphql.Scalar

func (this *Scalar) output() graphql.Output {
	s := graphql.Scalar(*this)
	return &s
}
func newScalar(s *graphql.Scalar) *Scalar {
	ss := Scalar(*s)
	return &ss
}
