package cmd

import (
	"github.com/racing-telemetry/f1-dump/internal/text/emoji"
	"github.com/racing-telemetry/f1-dump/internal/text/printer"
	"github.com/racing-telemetry/f1-dump/pkg/broadcaster"
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
		flags, err := buildFlags(cmd)
		if err != nil {
			return printer.Error(err)
		}

		b, err := broadcaster.NewBroadcaster(flags.UDPAddr())
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

		printer.Print(emoji.Sparkless, "Broadcast started at %s:%d, press Ctrl+C to stop", flags.host, flags.port)
		err = b.Start(flags.file)
		if err != nil {
			return printer.Error(err)
		}

		b.Stop()
		printer.Print(emoji.Rocket, "%d packs have been published", b.Stats.RecvCount())

		return nil
	},
}

func init() {
	addFlags(broadcastCmd)

	rootCmd.AddCommand(broadcastCmd)
}
