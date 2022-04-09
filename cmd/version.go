package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/racing-telemetry/f1-dump/internal"
	"github.com/racing-telemetry/f1-dump/internal/text/printer"
	"github.com/spf13/cobra"
	"runtime"
)

var (
	commit = "unknown"
	date   = "unknown"
)

type CLIVersionInfo struct {
	Version   string `json:"version"`
	Commit    string `json:"commit"`
	Date      string `json:"date"`
	GoVersion string `json:"go"`
	Compiler  string `json:"compiler"`
	Platform  string `json:"platform"`
}

func VersionInfo() *CLIVersionInfo {
	return &CLIVersionInfo{
		Version:   internal.Version,
		Commit:    commit,
		Date:      date,
		GoVersion: runtime.Version(),
		Compiler:  runtime.Compiler,
		Platform:  fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:          "version",
		Short:        "Prints the CLI version",
		Long:         "Prints the CLI version",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			bytes, err := json.Marshal(VersionInfo())
			if err != nil {
				return printer.Error(errors.New("failed to marshal version info"))
			}

			fmt.Println(string(bytes))
			return nil
		},
	})
}
