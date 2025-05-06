package fixer

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
	"go.uber.org/zap"

	"github.com/zguydev/openapi-fixer/internal/config"
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

func (o *OpenAPISpecFixer) Fix(inputSpecPath, outSpecPath, rulesPath string) error {
	doc, err := o.loadSpec(inputSpecPath)
	if err != nil {
		return fmt.Errorf("o.loadSpec: %w", err)
	}

	rules, err := o.loadRules(rulesPath)
	if err != nil {
		return fmt.Errorf("o.loadRules: %w", err)
	}

	if err := o.applyFixes(doc, rules); err != nil {
		return fmt.Errorf("o.applyFixes: %w", err)
	}

	if err := o.exportSpec(outSpecPath, doc); err != nil {
		return fmt.Errorf("o.exportSpec: %w", err)
	}
	return nil
}

func (o *OpenAPISpecFixer) applyFixes(doc *openapi3.T, rules []FixRule) error {
	for _, rule := range rules {
		o.logger.Info("applying OpenAPI fix rule",
			zap.String("rule", rule.Name()))
		if err := rule.Apply(doc); err != nil {
			o.logger.Error("failed to apply OpenAPI fix rule",
				zap.Error(err), zap.String("rule", rule.Name()))
			return fmt.Errorf("rule.Apply: %w", err)
		}
	}
	return nil
}
