package cli

import (
	"github.com/russelltsherman/blinkt/lib"
	"github.com/spf13/cobra"
)

// ClearCommand returns a cobra command
func ClearCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "clear",
		Short: "turn off all pizels",
		Run: func(cmd *cobra.Command, args []string) {
			brightness := 0.1
			blinkt := lib.NewBlinkt(brightness)
			blinkt.SetClearOnExit(true)
			blinkt.Setup()
			blinkt.Clear()
			blinkt.Show()
		},
	}
	return cmd
}
