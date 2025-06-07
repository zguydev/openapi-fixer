package main

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/zguydev/openapi-fixer/pkg/fixup"
)

type EnumFixup struct {
	fixup.OpenAPIFixup
}

func (f *EnumFixup) Name() string {
	return "EnumFixup"
}

func (f *EnumFixup) Apply(doc *openapi3.T) error {
	fixes := []struct {
		fixName string
		fixFunc func(doc *openapi3.T) error
	}{
		{
			"f.fix_ModelIdsShared",
			f.fix_ModelIdsShared,
		},
		{
			"f.fix_VoiceIdsShared",
			f.fix_VoiceIdsShared,
		},
	}
	for _, fix := range fixes {
		if err := fix.fixFunc(doc); err != nil {
			return fmt.Errorf("%s: %w", fix.fixName, err)
		}
	}
	return nil
}

func (f *EnumFixup) fix_ModelIdsShared(doc *openapi3.T) error {
	target := "ModelIdsShared"

	schemaRef, ok := doc.Components.Schemas[target]
	if !ok {
		return fmt.Errorf("component schema %s not exists", target)
	}
	if schemaRef.Value == nil {
		return fmt.Errorf("component schema %s is nil", target)
	}

	schemaRef.Value.AnyOf = nil
	schemaRef.Value.Type = &openapi3.Types{"string"}
	return nil
}

func (f *EnumFixup) fix_VoiceIdsShared(doc *openapi3.T) error {
	target := "VoiceIdsShared"

	schemaRef, ok := doc.Components.Schemas[target]
	if !ok {
		return fmt.Errorf("component schema %s not exists", target)
	}
	if schemaRef.Value == nil {
		return fmt.Errorf("component schema %s is nil", target)
	}

	schemaRef.Value.AnyOf = nil
	schemaRef.Value.Type = &openapi3.Types{"string"}
	return nil
}
