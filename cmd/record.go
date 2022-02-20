package cmd

import (
	"errors"
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/racing-telemetry/f1-dump/cmd/flags"
	"github.com/racing-telemetry/f1-dump/internal/stream"
	"github.com/racing-telemetry/f1-dump/internal/text/emoji"
	"github.com/racing-telemetry/f1-dump/internal/text/printer"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

var cmdRecord = &cobra.Command{
	Use:          "record",
	Short:        "Start recording packets from UDP source.",
	Long:         `Start recording packets from UDP source.`,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		flags, err := flags.Build(cmd)
		if err != nil {
			return printer.Error(err)
		}

		rc, err := stream.NewRecorder(flags)
		if err != nil {
			return fmt.Errorf("%s\n%s", printer.Error(errors.New("recorder can't create")), printer.Error(err))
		}

		// wait exit signal, ctrl+c to exit
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)

		go func() {
			<-c

			rc.Stop()

			f, err := rc.Save()
			if err != nil {
				printer.PrintError(err.Error())
			} else {
				stat, err := f.Stat()
				if err != nil {
					printer.PrintError(err.Error())
				}

				printer.Print(emoji.File, "File saved to %s (size: %s)", f.Name(), humanize.Bytes(uint64(stat.Size())))

				err = f.Close()
				if err != nil {
					printer.PrintError(err.Error())
				}
			}

			printer.Print(emoji.Rocket, "Received Packets: %d", rc.Stats.RecvCount())
			printer.Print(emoji.Construction, "Lost Packets: %d", rc.Stats.ErrCount())

			os.Exit(0)
		}()

		if len(flags.Packs) != 0 {
			printer.Print(emoji.RoundPushpin, "These packets are being ignored: %s", flags.Packs.Pretty())
		}

		printer.Print(emoji.Sparkless, "Record started at %s:%d, press Ctrl+C to stop", flags.Host, flags.Port)
		rc.Start()

		return nil
	},
}

func init() {
	flags.Add(cmdRecord)

	rootCmd.AddCommand(cmdRecord)
}
