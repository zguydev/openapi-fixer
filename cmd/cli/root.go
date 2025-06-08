package cli

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "openapi-fixer input_file output_file --fixups fixups_path --config fixer_config",
	Short: "A powerful tool to fix OpenAPI spec to ensure compatibility with various code generators and tools",
	Args:  cobra.ExactArgs(2),
	Run:   run,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().String("config", ".openapi-fixer.yaml", "Path to fixer config")
	rootCmd.Flags().String("fixups", "./fixups/", "Path to fixups")
	_ = rootCmd.MarkFlagRequired("fixups")
}
