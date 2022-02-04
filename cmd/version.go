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
	GitCommitSHA = "unknown"
	BuildDate    = "unknown"
)

type CLIVersionInfo struct {
	Version      string
	GitCommitSHA string
	BuildDate    string
	GoVersion    string
	Compiler     string
	Platform     string
}

func VersionInfo() *CLIVersionInfo {
	return &CLIVersionInfo{
		Version:      internal.Version,
		GitCommitSHA: GitCommitSHA,
		BuildDate:    BuildDate,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
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
