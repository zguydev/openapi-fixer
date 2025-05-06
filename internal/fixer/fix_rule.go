package fixer

import "github.com/getkin/kin-openapi/openapi3"

type FixRule interface {
	Name() string
	Apply(doc *openapi3.T) error
}
