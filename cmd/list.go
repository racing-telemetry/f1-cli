package cmd

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/racing-telemetry/f1-cli/internal/packets"
	"github.com/racing-telemetry/f1-cli/internal/text/printer"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:          "list",
		Short:        "List F1 packet types.",
		Long:         "List F1 packet types.",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := ui.Init(); err != nil {
				return printer.Error(err)
			}

			defer ui.Close()

			p := widgets.NewParagraph()
			p.Text = packets.Pretty()
			p.Border = false
			p.SetRect(0, 0, 128, 128)

			ui.Render(p)

			for e := range ui.PollEvents() {
				if e.Type == ui.KeyboardEvent {
					break
				}
			}

			return nil
		},
	})
}
