package cli

import (
	"github.com/russelltsherman/blinkt/lib"
	"github.com/spf13/cobra"
)

// GraphCommand returns a cobra command
func GraphCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "graph",
		Short: "",
		Run: func(cmd *cobra.Command, args []string) {
			brightness := 0.1
			blinkt := lib.NewBlinkt(brightness)
			blinkt.Setup()
			lib.Delay(100)

		},
	}
	return cmd
}

// https://github.com/pimoroni/blinkt/blob/master/examples/graph.py
