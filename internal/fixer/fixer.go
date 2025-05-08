package fixer

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
	"go.uber.org/zap"

	"github.com/zguydev/openapi-fixer/internal/config"
	"github.com/zguydev/openapi-fixer/pkg/fixup"
)

type OpenAPISpecFixer struct {
	cfg    *config.FixerConfig
	logger *zap.Logger
	loader *openapi3.Loader
}

func NewOpenAPISpecFixer(cfg *config.FixerConfig, logger *zap.Logger) *OpenAPISpecFixer {
	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true

	return &OpenAPISpecFixer{
		cfg:    cfg,
		logger: logger,
		loader: loader,
	}
}

func (o *OpenAPISpecFixer) Fix(inputSpecPath, outSpecPath, fixupsPath string) error {
	doc, err := o.loadSpec(inputSpecPath)
	if err != nil {
		return fmt.Errorf("o.loadSpec: %w", err)
	}

	fixups, err := o.loadFixups(fixupsPath)
	if err != nil {
		return fmt.Errorf("o.loadFixups: %w", err)
	}

	if err := o.applyFixups(doc, fixups); err != nil {
		return fmt.Errorf("o.applyFixups: %w", err)
	}

	if err := o.exportSpec(outSpecPath, doc); err != nil {
		return fmt.Errorf("o.exportSpec: %w", err)
	}
	return nil
}

func (o *OpenAPISpecFixer) applyFixups(doc *openapi3.T, fixups []fixup.OpenAPIFixup) error {
	for _, fixup := range fixups {
		o.logger.Info("applying OpenAPI fixup",
			zap.String("fixup", fixup.Name()))
		if err := fixup.Apply(doc); err != nil {
			o.logger.Error("failed to apply OpenAPI fixup",
				zap.Error(err), zap.String("fixup", fixup.Name()))
			return fmt.Errorf("fixup.Apply: %w", err)
		}
	}
	return nil
}
