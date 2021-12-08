package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"gitlab-nomo.credissimo.net/nomo/cosmzone/x/suspend/types"

	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group suspend queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	// this line is used by starport scaffolding # 1

	cmd.AddCommand(QueryGetSuspendCmd())

	return cmd
}

func QueryGetSuspendCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Query the current suspend value",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)

			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			params := &types.QuerySuspendRequest{}
			res, err := queryClient.Suspend(cmd.Context(), params)

			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
