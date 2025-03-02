package keeper

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/archway-network/archway/x/rewards/types"
)

// SetFlatFee checks if a contract has metadata set and stores the given flat fee to be associated with that contract
func (k Keeper) SetFlatFee(ctx sdk.Context, senderAddr sdk.AccAddress, feeUpdate types.FlatFee) error {
	// Check if the contract metadata exists
	contractInfo := k.GetContractMetadata(ctx, feeUpdate.MustGetContractAddress())
	if contractInfo == nil {
		return types.ErrMetadataNotFound
	}
	if contractInfo.OwnerAddress != senderAddr.String() {
		return errorsmod.Wrap(types.ErrUnauthorized, "flat_fee can only be set or changed by the contract owner")
	}

	if feeUpdate.FlatFee.Amount.IsZero() {
		k.state.FlatFee(ctx).RemoveFee(feeUpdate.MustGetContractAddress())
	} else {
		if contractInfo.RewardsAddress == "" {
			return errorsmod.Wrap(types.ErrMetadataNotFound, "flat_fee can only be set when rewards address has been configured")
		}
		k.state.FlatFee(ctx).SetFee(feeUpdate.MustGetContractAddress(), feeUpdate.FlatFee)
	}

	types.EmitContractFlatFeeSetEvent(
		ctx,
		feeUpdate.MustGetContractAddress(),
		feeUpdate.FlatFee,
	)
	return nil
}

// GetFlatFee retreives the flat fee stored for a given contract
func (k Keeper) GetFlatFee(ctx sdk.Context, contractAddr sdk.AccAddress) (sdk.Coin, bool) {
	fee, found := k.state.FlatFee(ctx).GetFee(contractAddr)
	if !found {
		return sdk.Coin{}, false
	}

	return fee, true
}

// CreateFlatFeeRewardsRecords creates a rewards record for the flatfees of the given contract
func (k Keeper) CreateFlatFeeRewardsRecords(ctx sdk.Context, contractAddress sdk.AccAddress, flatfees sdk.Coins) {
	rewardsRecordState := k.state.RewardsRecord(ctx)
	calculationHeight, calculationTime := ctx.BlockHeight(), ctx.BlockTime()

	metadata := k.GetContractMetadata(ctx, contractAddress)
	rewardsAddr := sdk.MustAccAddressFromBech32(metadata.RewardsAddress)

	rewardsRecordState.CreateRewardsRecord(rewardsAddr, flatfees, calculationHeight, calculationTime)
}
