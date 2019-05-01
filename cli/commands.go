package cli

import (
	"github.com/spf13/cobra"
)

// AddCommands add the commands from subcommand groups to the root command
func AddCommands(cmd *cobra.Command) {
	cmd.AddCommand(
		BinaryClockCommand(),
		CandleCommand(),
		ClearCommand(),
		CPULoadCommand(),
		GraphCommand(),
		LarsonCommand(),
		PulseCommand(),
		RevolvingCommand(),
		SequenceCommand(),
		SolidCommand(),
		VersionCommand(),
	)
}
