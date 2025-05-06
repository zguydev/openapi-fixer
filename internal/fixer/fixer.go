package fixer

import (
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

func (o *OpenAPISpecFixer) Fix(inputSpecPath, outSpecPath string) error {
	return nil
}
