package main

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
)

type NumberTypeFixup struct{}

func (f *NumberTypeFixup) Name() string {
	return "NumberTypeFixup"
}

func (f *NumberTypeFixup) Apply(doc *openapi3.T) error {
	fixes := []struct {
		fixName string
		fixFunc func(doc *openapi3.T) error
	}{
		{
			"f.Fix_CreateChatCompletionRequest",
			f.Fix_CreateChatCompletionRequest,
		},
	}
	for _, fix := range fixes {
		if err := fix.fixFunc(doc); err != nil {
			return fmt.Errorf("%s: %w", fix.fixName, err)
		}
	}
	return nil
}

func (f *NumberTypeFixup) Fix_CreateChatCompletionRequest(doc *openapi3.T) error {
	target := "CreateChatCompletionRequest"

	schemaRef, ok := doc.Components.Schemas[target]
	if !ok {
		return fmt.Errorf("component schema %s not exists", target)
	}
	if schemaRef.Value == nil {
		return fmt.Errorf("component schema %s is nil", target)
	}

	allOf := schemaRef.Value.AllOf

	for _, schemaRef := range allOf {
		if schemaRef.Value == nil || len(schemaRef.Value.Properties) == 0 {
			continue
		}
		seed, ok := schemaRef.Value.Properties["seed"]
		if !ok {
			continue
		}
		seedValue := seed.Value
		if seedValue == nil {
			continue
		}
		seedValue.Type = &openapi3.Types{"number"}
		return nil
	}
	return fmt.Errorf("element to fixup not found")
}
