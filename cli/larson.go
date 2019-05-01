package cli

import (
	"github.com/russelltsherman/blinkt/lib"
	"github.com/spf13/cobra"
)

// LarsonCommand returns a cobra command
func LarsonCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "larson",
		Short: "larson scanner sequence",
		Run: func(cmd *cobra.Command, args []string) {
			brightness := 0.1
			blinkt := lib.NewBlinkt(brightness)
			blinkt.Setup()
			lib.Delay(100)

			r := 150
			g := 0
			b := 0

			for {
				for pixel := 0; pixel <= 7; pixel++ {
					blinkt.Clear()
					blinkt.SetPixel(pixel, r, g, b)
					blinkt.Show()
					lib.Delay(100)
				}
				for pixel := 7; pixel >= 0; pixel-- {
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
