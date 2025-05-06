package cli

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "openapi-fixer input output --rules rules_path --config fixer_config",
	Short: "Fix an OpenAPI spec to selectively comply for example with code generators",
	Args:  cobra.ExactArgs(2),
	Run:   run,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().String("config", "openapi-fixer.yaml", "Path to fixer config")
	rootCmd.Flags().String("rules", "./rules/", "Path to fixer rules")
	_ = rootCmd.MarkFlagRequired("config")
	_ = rootCmd.MarkFlagRequired("rules")
}
