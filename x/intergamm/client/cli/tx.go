package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/quasarlabs/quasarnode/x/intergamm/types"
	"github.com/cosmos/cosmos-sdk/client"
)

var cmds []*cobra.Command

func addCommand(cmd *cobra.Command) {
	cmds = append(cmds, cmd)
}

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	cmd.AddCommand(cmds...)
	// this line is used by starport scaffolding # 1

	return cmd
}
