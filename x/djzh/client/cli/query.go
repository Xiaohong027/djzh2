package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
/*	sdk "github.com/cosmos/cosmos-sdk/types"*/

	"github.com/changtong1996/djzh/x/djzh/internal/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	// Group djzh queries under a subcommand
	djzhQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	djzhQueryCmd.AddCommand(
		flags.GetCommands(
	// TODO: Add query Cmds
			GetCmdArticle(queryRoute, cdc),
			GetCmdVoteNum(queryRoute, cdc),
		)...,
	)

	return djzhQueryCmd
}

// TODO: Add Query Commands
func GetCmdArticle(queryRoute string, cdc *codec.Codec) *cobra.Command {
		return &cobra.Command{
			Use:     "article [article_id]",
			Short:   "Query a article by article_id",
			Args:    cobra.ExactArgs(1),
			RunE:    func(cmd *cobra.Command, args []string) error{
				cliCtx := context.NewCLIContext().WithCodec(cdc)
				article_id := args[0]

				res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/getarticle/%s", queryRoute, article_id), nil)
				if err != nil {
					fmt.Printf("could not get article - %s \n", article_id)
					return nil
				}

				var out types.Article
				cdc.MustUnmarshalJSON(res, &out)
				return cliCtx.PrintOutput(out)
			},
		}
}


func GetCmdVoteNum(queryRoute string, cdc *codec.Codec) *cobra.Command {
		return &cobra.Command{
			Use:     "vote number of article [name]",
			Short:   "get the vote number of a article",
			Args:    cobra.ExactArgs(1),
			RunE:    func(cmd *cobra.Command, args []string) error{
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			name := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/whois/%s", queryRoute, name), nil)
			if err != nil {
				fmt.Printf("could not get  - %s \n", name)
				return nil
			}
			var out types.Article
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
			},
		}
}