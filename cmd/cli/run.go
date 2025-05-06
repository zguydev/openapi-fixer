package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/zguydev/openapi-fixer/internal/config"
	"github.com/zguydev/openapi-fixer/internal/fixer"
	"github.com/zguydev/openapi-fixer/internal/utils"
)

func run(cmd *cobra.Command, args []string) {
	fallbackLogger := utils.NewFallbackLogger()
	defer fallbackLogger.Sync() //nolint:errcheck

	configPath, err := cmd.Flags().GetString("config")
	if err != nil {
		fallbackLogger.Fatal("failed to get config flag", zap.Error(err))
	}

	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		fallbackLogger.Fatal("failed to load config",
			zap.Error(fmt.Errorf("config.LoadConfig: %w", err)))
	}

	logger, err := utils.NewLogger(cfg.Tool.Logger)
	if err != nil {
		fallbackLogger.Fatal("failed to init logger",
			zap.Error(fmt.Errorf("utils.NewLogger: %w", err)))
	}

	oafixer := fixer.NewOpenAPISpecFixer(cfg, logger)
	if err := oafixer.Fix(args[0], args[1]); err != nil {
		logger.Error("fix on spec failed",
			zap.Error(fmt.Errorf("oafixer.Fix: %w", err)))
		os.Exit(1)
	}
}
