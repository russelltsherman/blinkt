package cli

import (
	"github.com/russelltsherman/blinkt/lib"
	"github.com/spf13/cobra"
)

// CPULoadCommand returns a cobra command
func CPULoadCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cpuload",
		Short: "a visual representation of cpu load",
		Run: func(cmd *cobra.Command, args []string) {
			brightness := 0.1
			blinkt := lib.NewBlinkt(brightness)
			blinkt.Setup()
			lib.Delay(100)

		},
	}
	return cmd
}

// https://github.com/pimoroni/blinkt/blob/master/examples/cpu_load.py
