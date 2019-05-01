package cli

import (
	"github.com/russelltsherman/blinkt/lib"
	"github.com/spf13/cobra"
)

// BinaryClockCommand returns a cobra command
func BinaryClockCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "binaryclock",
		Short: "a binary clock",
		Run: func(cmd *cobra.Command, args []string) {
			brightness := 0.1
			blinkt := lib.NewBlinkt(brightness)
			blinkt.Setup()
			lib.Delay(100)

		},
	}
	return cmd
}

// https://github.com/pimoroni/blinkt/blob/master/examples/binary_clock.py
// https://github.com/pimoroni/blinkt/blob/master/examples/binary_clock_meld.py
