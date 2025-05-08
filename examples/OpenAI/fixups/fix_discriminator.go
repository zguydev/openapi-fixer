package main

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
)

type DiscriminatorFixup struct{}

func (f *DiscriminatorFixup) Name() string {
	return "FixDiscriminator"
}

func (f *DiscriminatorFixup) Apply(doc *openapi3.T) error {
	fixes := []struct {
		fixName string
		fixFunc func(doc *openapi3.T) error
	}{
		{
			"f.Fix_ChatCompletionRequestMessage",
			f.Fix_ChatCompletionRequestMessage,
		},
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

func (f *DiscriminatorFixup) Fix_ChatCompletionRequestMessage(doc *openapi3.T) error {
	target := "ChatCompletionRequestMessage"
	discriminatorProperty := "role"
	mapping := map[string]string{
		"developer": "#/components/schemas/ChatCompletionRequestDeveloperMessage",
		"system":    "#/components/schemas/ChatCompletionRequestSystemMessage",
		"user":      "#/components/schemas/ChatCompletionRequestUserMessage",
		"assistant": "#/components/schemas/ChatCompletionRequestAssistantMessage",
		"tool":      "#/components/schemas/ChatCompletionRequestToolMessage",
		"function":  "#/components/schemas/ChatCompletionRequestFunctionMessage",
	}

	schemaRef, ok := doc.Components.Schemas[target]
	if !ok {
		return fmt.Errorf("component schema %s not exists", target)
	}
	if schemaRef.Value == nil {
		return fmt.Errorf("component schema %s is nil", target)
	}
	f.addDiscriminator(schemaRef, discriminatorProperty, mapping)
	return nil
}

func (f *DiscriminatorFixup) Fix_CreateChatCompletionRequest(doc *openapi3.T) error {
	target := "CreateChatCompletionRequest"
	discriminatorProperty := "type"
	mapping := map[string]string{
		"text":        "#/components/schemas/ResponseFormatText",
		"json_schema": "#/components/schemas/ResponseFormatJsonSchema",
		"json_object": "#/components/schemas/ResponseFormatJsonObject",
	}

	schemaRef, ok := doc.Components.Schemas[target]
	if !ok {
		return fmt.Errorf("component schema %s not exists", target)
	}
	if schemaRef.Value == nil {
		return fmt.Errorf("component schema %s is nil", target)
	}

	allOf := schemaRef.Value.AllOf
	for _, schemaRef := range allOf {
		if !(schemaRef.Value != nil && len(schemaRef.Value.Properties) != 0) {
			continue
		}
		responseFormat, ok := schemaRef.Value.Properties["response_format"]
		if !ok {
			continue
		}
		f.addDiscriminator(responseFormat, discriminatorProperty, mapping)
		return nil
	}
	return fmt.Errorf("element to fixup not found")
}

func (f *DiscriminatorFixup) addDiscriminator(schemaRef *openapi3.SchemaRef,
	discriminatorProperty string, mapping map[string]string) {
	schema := schemaRef.Value
	if schema.Discriminator == nil {
		schema.Discriminator = &openapi3.Discriminator{}
	}
	schema.Discriminator.PropertyName = discriminatorProperty
	schema.Discriminator.Mapping = mapping
}
