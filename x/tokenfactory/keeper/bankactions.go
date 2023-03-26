package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/CosmWasm/token-factory/x/tokenfactory/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

func (k Keeper) mintTo(ctx sdk.Context, amount sdk.Coin, mintTo string) error {
	// verify that denom is an x/tokenfactory denom
	_, _, err := types.DeconstructDenom(amount.Denom)
	if err != nil {
		return err
	}

	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(amount))
	if err != nil {
		return err
	}

	addr, err := sdk.AccAddressFromBech32(mintTo)
	if err != nil {
		return err
	}

	if k.bankKeeper.BlockedAddr(addr) {
		return fmt.Errorf("failed to mint to blocked address: %s", addr)
	}

	return k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName,
		addr,
		sdk.NewCoins(amount))
}

func (k Keeper) burnFrom(ctx sdk.Context, amount sdk.Coin, burnFrom string) error {
	// verify that denom is an x/tokenfactory denom
	_, _, err := types.DeconstructDenom(amount.Denom)
	if err != nil {
		return err
	}

	addr, err := sdk.AccAddressFromBech32(burnFrom)
	if err != nil {
		return err
	}

	if k.bankKeeper.BlockedAddr(addr) {
		return fmt.Errorf("failed to burn from blocked address: %s", addr)
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx,
		addr,
		types.ModuleName,
		sdk.NewCoins(amount))
	if err != nil {
		return err
	}

	return k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(amount))
}

func (k Keeper) forceTransfer(ctx sdk.Context, amount sdk.Coin, fromAddr string, toAddr string) error {
	// verify that denom is an x/tokenfactory denom
	_, _, err := types.DeconstructDenom(amount.Denom)
	if err != nil {
		return err
	}

	fromSdkAddr, err := sdk.AccAddressFromBech32(fromAddr)
	if err != nil {
		return err
	}

	toSdkAddr, err := sdk.AccAddressFromBech32(toAddr)
	if err != nil {
		return err
	}

	if k.bankKeeper.BlockedAddr(toSdkAddr) {
		return fmt.Errorf("failed to force transfer to blocked address: %s", toSdkAddr)
	}

	return k.bankKeeper.SendCoins(ctx, fromSdkAddr, toSdkAddr, sdk.NewCoins(amount))
}

func (k Keeper) toggleTransfer(ctx sdk.Context, denom string, enabledStatus bool) error {
	// verify that denom is an x/tokenfactory denom
	_, _, err := types.DeconstructDenom(denom)
	if err != nil {
		return err
	}

	var updatedCoins []*banktypes.SendEnabled

	coins := k.sendKeeper.GetParams(ctx).SendEnabled

	for _, coin := range coins {
		if coin.Denom == denom {
			// TODO: if its the same status, error here?
			if enabledStatus {
				continue
			}
			coin.Enabled = enabledStatus
		}

		updatedCoins = append(updatedCoins, coin)
	}

	k.sendKeeper.SetParams(ctx, banktypes.Params{SendEnabled: updatedCoins})

	return nil
}
