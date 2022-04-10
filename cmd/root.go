package cmd

import (
	"github.com/racing-telemetry/f1-cli/internal/opts"
	"github.com/spf13/cobra"
	"os"
)

var root = &cobra.Command{
	Use:   "f1dump",
	Short: "Dump F1 data",
	Long:  `A helper CLI for dumping F1 data`,
}

func Execute() {
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	root.PersistentFlags().BoolVarP(&opts.Verbose, "verbose", "v", false, "verbose output")
}
