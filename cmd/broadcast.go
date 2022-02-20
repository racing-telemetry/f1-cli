package cmd

import (
	"github.com/racing-telemetry/f1-dump/cmd/flags"
	"github.com/racing-telemetry/f1-dump/internal/stream"
	"github.com/racing-telemetry/f1-dump/internal/text/emoji"
	"github.com/racing-telemetry/f1-dump/internal/text/printer"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

var broadcastCmd = &cobra.Command{
	Use:          "broadcast",
	Short:        "Start broadcasting",
	Long:         `Start broadcasting`,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		instant, _ := cmd.Flags().GetBool("instant")
		flags, err := flags.Build(cmd)
		if err != nil {
			return printer.Error(err)
		}

		b, err := stream.NewBroadcaster(flags)
		if err != nil {
			return printer.Error(err)
		}

		// wait exit signal, ctrl+c to early exit
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)

		go func() {
			<-c

			b.Stop()
			printer.Print(emoji.Rocket, "%d packs have been published", b.Stats.RecvCount())

			os.Exit(0)
		}()

		if len(flags.Packs) != 0 {
			printer.Print(emoji.RoundPushpin, "These packets are being ignored: %s", flags.Packs.Pretty())
		}

		printer.Print(emoji.Sparkless, "Broadcast started at %s:%d, press Ctrl+C to stop", flags.Host, flags.Port)
		err = b.Start(instant)
		if err != nil {
			return printer.Error(err)
		}

		b.Stop()
		printer.Print(emoji.Rocket, "%d packet have been published", b.Stats.RecvCount())

		return nil
	},
}

func init() {
	flags.Add(broadcastCmd)

	broadcastCmd.Flags().BoolP("instant", "i", false, "Broadcast all packets instantly. (Not Thread Safe)")

	rootCmd.AddCommand(broadcastCmd)
}
