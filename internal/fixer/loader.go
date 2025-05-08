package fixer

import (
	"fmt"
	"os"
	"os/exec"
	"plugin"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"

	"github.com/zguydev/openapi-fixer/pkg/fixup"
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

func (o *OpenAPISpecFixer) loadFixups(fixupsPath string) ([]fixup.OpenAPIFixup, error) {
	info, err := os.Stat(fixupsPath)
	if err != nil {
		o.logger.Error("error accessing fixups path",
			zap.Error(err), zap.String("path", fixupsPath))
		return nil, fmt.Errorf("os.Stat: %w", err)
	}

	switch {
	case info.IsDir():
		o.logger.Info("fixups path specified as directory, trying to build and load as Go plugin")

		soPath := strings.TrimRight(fixupsPath, "/") + ".so"
		cmd := exec.Command("go", "build", "-buildmode=plugin", "-o", soPath, fixupsPath)
		cmd.Env = append(os.Environ(), "GO111MODULE=on")

		output, err := cmd.CombinedOutput()
		if err != nil {
			o.logger.Error("failed to build plugin",
				zap.Error(err),
				zap.ByteString("output", output))
			return nil, fmt.Errorf("failed to build plugin from fixups path %q: %w", fixupsPath, err)
		}
		o.logger.Info("built fixups to plugin", zap.String("plugin", soPath))
		fixupsPath = soPath
		fallthrough

	case !info.IsDir() && strings.HasSuffix(fixupsPath, ".so"):
		o.logger.Info("loading fixups as Go plugin")

		p, err := plugin.Open(fixupsPath)
		if err != nil {
			o.logger.Error("failed to load fixups path as plugin",
				zap.Error(err), zap.String("path", fixupsPath))
			return nil, fmt.Errorf("plugin.Open: %w", err)
		}

		sym, err := p.Lookup("Fixups")
		if err != nil {
			o.logger.Error("fixups plugin does not export Fixups variable",
				zap.Error(err), zap.String("path", fixupsPath))
			return nil, fmt.Errorf("p.Lookup: %w", err)
		}

		fixups, ok := sym.(*[]fixup.OpenAPIFixup)
		if !ok {
			o.logger.Error("fixups plugin has wrong value type",
				zap.String("path", fixupsPath))
			return nil, fmt.Errorf("sym.([]fixup.OpenAPIFixup): %w", err)
		}
		return *fixups, nil

	default:
		o.logger.Error("unknown fixups path type",
			zap.String("path", fixupsPath))
		return nil, fmt.Errorf("unknown fixups path type")
	}
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
