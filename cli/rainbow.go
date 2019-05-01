package cli

import (
	colorsys "github.com/cckuailong/colorsys-go"
	"github.com/russelltsherman/blinkt/lib"
	"github.com/spf13/cobra"
)

// RainbowCommand returns a cobra command
func RainbowCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pulse",
		Short: "",
		Run: func(cmd *cobra.Command, args []string) {
			brightness := 0.1
			blinkt := lib.NewBlinkt(brightness)
			blinkt.Setup()
			lib.Delay(100)

			h := 0
			for {
				h++
				if h > 360 {
					h = 0
				}

				r, g, b := colorsys.Hsv2Rgb(float64(h), 1.0, brightness)
				// fmt.Printf("R: %v G: %v B: %v\n", r, g, b)
				blinkt.Clear()
				blinkt.SetAll(int(r), int(g), int(b))
				blinkt.Show()
				lib.Delay(100)
			}

		},
	}
	return cmd
}

// https://github.com/pimoroni/blinkt/blob/master/examples/rainbow.py
