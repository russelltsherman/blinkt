package cli

import (
	"github.com/russelltsherman/blinkt/lib"
	"github.com/spf13/cobra"
)

// SequenceCommand returns a cobra command
func SequenceCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sequence",
		Short: "light up in sequence",
		Run: func(cmd *cobra.Command, args []string) {
			brightness := 0.1
			blinkt := lib.NewBlinkt(brightness)
			blinkt.SetClearOnExit(true)
			blinkt.Setup()
			lib.Delay(100)

			r := 150
			g := 0
			b := 0

			for {
				for pixel := 0; pixel < 8; pixel++ {
					blinkt.Clear()
					blinkt.SetPixel(pixel, r, g, b)
					blinkt.Show()
					lib.Delay(100)
				}
			}
		},
	}
	return cmd
}
