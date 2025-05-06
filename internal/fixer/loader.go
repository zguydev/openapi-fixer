package fixer

import (
	"fmt"
	"os"
	"path/filepath"
	"plugin"

	"github.com/getkin/kin-openapi/openapi3"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

func (o *OpenAPISpecFixer) loadSpec(specPath string) (*openapi3.T, error) {
	inputFileData, err := os.ReadFile(specPath)
	if err != nil {
		o.logger.Error("failed to read OpenAPI spec file",
			zap.Error(err))
		return nil, fmt.Errorf("os.ReadFile: %w", err)
	}

	doc, err := o.loader.LoadFromData(inputFileData)
	if err != nil {
		o.logger.Error("failed to load OpenAPI spec",
			zap.Error(err))
		return nil, fmt.Errorf("loader.LoadFromData: %w", err)
	}
	return doc, nil
}

func (o *OpenAPISpecFixer) loadRules(rulesPath string) ([]FixRule, error) {
	var rules []FixRule

	err := filepath.Walk(rulesPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			o.logger.Error("error accessing rules path",
				zap.Error(err), zap.String("path", path))
			return fmt.Errorf("func filepath.Walk(%q): %w", path, err)
		}
		if info.IsDir() || filepath.Ext(path) != ".so" {
			return nil
		}

		p, err := plugin.Open(path)
		if err != nil {
			o.logger.Error("failed to load plugin",
				zap.Error(err), zap.String("path", path))
			return fmt.Errorf("plugin.Open(%q): %w", path, err)
		}

		sym, err := p.Lookup("Rule")
		if err != nil {
			o.logger.Error("plugin does not export Rule",
				zap.Error(err), zap.String("path", path))
			return fmt.Errorf("(%q)p.Lookup: %w", path, err)
		}

		rule, ok := sym.(FixRule)
		if !ok {
			o.logger.Error("plugin has wrong type",
				zap.Error(err), zap.String("path", path))
			return fmt.Errorf("(%q)sym.(FixRule): %w", path, err)
		}

		rules = append(rules, rule)
		return nil
	})
	if err != nil {
		o.logger.Error("failed to walk path for finding plugins",
			zap.Error(err), zap.String("rulesPath", rulesPath))
		return nil, fmt.Errorf("filepath.Walk: %w", err)
	}
	return rules, nil
}

func (o *OpenAPISpecFixer) exportSpec(outSpecPath string, doc *openapi3.T) error {
	f, err := os.Create(outSpecPath)
	if err != nil {
		o.logger.Error("failed to create output spec file",
			zap.Error(err))
		return fmt.Errorf("os.Create: %w", err)
	}
	defer f.Close()

	encoder := yaml.NewEncoder(f)
	encoder.SetIndent(2)
	defer encoder.Close()

	if err := encoder.Encode(doc); err != nil {
		o.logger.Error("failed to encode output spec file",
			zap.Error(err))
		return fmt.Errorf("encoder.Encode: %w", err)
	}
	return nil
}
