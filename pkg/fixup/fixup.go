package fixup

import "github.com/getkin/kin-openapi/openapi3"

type OpenAPIFixup interface {
	Name() string
	Apply(doc *openapi3.T) error
}
