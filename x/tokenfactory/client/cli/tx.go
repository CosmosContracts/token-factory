package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/CosmWasm/token-factory/x/tokenfactory/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

var DeniedDenoms = [1]string{"juno"}

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		NewCreateDenomCmd(),
		NewMintCmd(),
		NewMintToCmd(),
		NewBurnCmd(),
		NewBurnFromCmd(),
		NewForceTransferCmd(),
		NewChangeAdminCmd(),
		NewModifyDenomMetadataCmd(),
	)

	return cmd
}

// NewCreateDenomCmd broadcast MsgCreateDenom
func NewCreateDenomCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-denom [subdenom] [flags]",
		Short: "create a new denom from an account",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			txf := tx.NewFactoryCLI(clientCtx, cmd.Flags()).WithTxConfig(clientCtx.TxConfig).WithAccountRetriever(clientCtx.AccountRetriever)

			msg := types.NewMsgCreateDenom(
				clientCtx.GetFromAddress().String(),
				args[0],
			)

			return tx.GenerateOrBroadcastTxWithFactory(clientCtx, txf, msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// NewMintCmd broadcast MsgMint
func NewMintCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mint [amount] [flags]",
		Short: "Mint a denom to your address. Must have admin authority to do so.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			txf := tx.NewFactoryCLI(clientCtx, cmd.Flags()).WithTxConfig(clientCtx.TxConfig).WithAccountRetriever(clientCtx.AccountRetriever)

			amount, err := sdk.ParseCoinNormalized(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgMint(
				clientCtx.GetFromAddress().String(),
				amount,
			)

			return tx.GenerateOrBroadcastTxWithFactory(clientCtx, txf, msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// NewMintToCmd broadcast MsgMintTo
func NewMintToCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mint-to [address] [amount] [flags]",
		Short: "Mint a denom to an address. Must have admin authority to do so.",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			txf := tx.NewFactoryCLI(clientCtx, cmd.Flags()).WithTxConfig(clientCtx.TxConfig).WithAccountRetriever(clientCtx.AccountRetriever)

			toAddr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			amount, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgMintTo(
				clientCtx.GetFromAddress().String(),
				amount,
				toAddr.String(),
			)

			return tx.GenerateOrBroadcastTxWithFactory(clientCtx, txf, msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// NewBurnCmd broadcast MsgBurn
func NewBurnCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "burn [amount] [flags]",
		Short: "Burn tokens from an address. Must have admin authority to do so.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			txf := tx.NewFactoryCLI(clientCtx, cmd.Flags()).WithTxConfig(clientCtx.TxConfig).WithAccountRetriever(clientCtx.AccountRetriever)

			amount, err := sdk.ParseCoinNormalized(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgBurn(
				clientCtx.GetFromAddress().String(),
				amount,
			)

			return tx.GenerateOrBroadcastTxWithFactory(clientCtx, txf, msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// NewBurnFromCmd broadcast MsgBurnFrom
func NewBurnFromCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "burn-from [address] [amount] [flags]",
		Short: "Burn tokens from an address. Must have admin authority to do so.",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			txf := tx.NewFactoryCLI(clientCtx, cmd.Flags()).WithTxConfig(clientCtx.TxConfig).WithAccountRetriever(clientCtx.AccountRetriever)

			fromAddr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			amount, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgBurnFrom(
				clientCtx.GetFromAddress().String(),
				amount,
				fromAddr.String(),
			)

			return tx.GenerateOrBroadcastTxWithFactory(clientCtx, txf, msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// NewForceTransferCmd broadcast MsgForceTransfer
func NewForceTransferCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "force-transfer [amount] [transfer-from-address] [transfer-to-address] [flags]",
		Short: "Force transfer tokens from one address to another address. Must have admin authority to do so.",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			txf := tx.NewFactoryCLI(clientCtx, cmd.Flags()).WithTxConfig(clientCtx.TxConfig).WithAccountRetriever(clientCtx.AccountRetriever)

			amount, err := sdk.ParseCoinNormalized(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgForceTransfer(
				clientCtx.GetFromAddress().String(),
				amount,
				args[1],
				args[2],
			)

			return tx.GenerateOrBroadcastTxWithFactory(clientCtx, txf, msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// NewChangeAdminCmd broadcast MsgChangeAdmin
func NewChangeAdminCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "change-admin [denom] [new-admin-address] [flags]",
		Short: "Changes the admin address for a factory-created denom. Must have admin authority to do so.",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			txf := tx.NewFactoryCLI(clientCtx, cmd.Flags()).WithTxConfig(clientCtx.TxConfig).WithAccountRetriever(clientCtx.AccountRetriever)

			msg := types.NewMsgChangeAdmin(
				clientCtx.GetFromAddress().String(),
				args[0],
				args[1],
			)

			return tx.GenerateOrBroadcastTxWithFactory(clientCtx, txf, msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// NewModifyDenomMetadataCmd broadcast a Bank Metadata modification transaction
func NewModifyDenomMetadataCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "modify-metadata [denom] [ticker-symbol] [description] [exponent] [flags]",
		Short: "Changes the base data for frontends to query the data of.",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			maxTickerLength := 6
			maxExponent := 18

			txf := tx.NewFactoryCLI(clientCtx, cmd.Flags()).WithTxConfig(clientCtx.TxConfig).WithAccountRetriever(clientCtx.AccountRetriever)

			fullDenom, ticker, desc := args[0], strings.ToUpper(args[1]), args[2]

			if !strings.HasPrefix(fullDenom, "factory/") {
				return fmt.Errorf("denom must start with factory/")
			}

			// Ticker Checks
			for _, prefix := range DeniedDenoms {
				if strings.Contains(strings.ToLower(ticker), prefix) {
					return fmt.Errorf("ticker contains a denied word: %s and is not allowed", prefix)
				}

				if strings.Contains(ticker, "/") {
					return fmt.Errorf("ticker cannot contain a / (slash)")
				}
			}

			// check if the length of ticker is greater than 6
			if len(ticker) > maxTickerLength {
				return fmt.Errorf("ticker cannot be greater than 6 characters")
			} else if len(ticker) == 0 {
				return fmt.Errorf("ticker cannot be empty")
			}

			// Description Length Checks
			if len(desc) > 255 {
				return fmt.Errorf("description cannot be greater than 255 characters: %d", len(desc))
			}

			deniedCharList := "@#$^*<>;()"
			if strings.ContainsAny(desc, deniedCharList) {
				return fmt.Errorf("desc cannot contain special characters: %s", deniedCharList)
			}

			// Exponent Checks
			exponent, err := strconv.ParseUint(args[3], 10, 32)
			if err != nil {
				return err
			}

			if exponent > uint64(maxExponent) {
				return fmt.Errorf("exponent cannot be greater than %d", maxExponent)
			}

			bankMetadata := banktypes.Metadata{
				Description: desc,
				Display:     fullDenom,
				Symbol:      ticker,
				Name:        fullDenom,
				DenomUnits: []*banktypes.DenomUnit{
					{
						Denom:    fullDenom,
						Exponent: 0, // must be 0 for the base denom
						Aliases:  []string{ticker},
					},
					{
						Denom:    ticker,
						Exponent: uint32(exponent),
						Aliases:  []string{fullDenom},
					},
				},
				Base: fullDenom,
			}

			msg := types.NewMsgSetDenomMetadata(
				clientCtx.GetFromAddress().String(),
				bankMetadata,
			)

			return tx.GenerateOrBroadcastTxWithFactory(clientCtx, txf, msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
