package da

import (
	"context"
	"encoding/binary"

	"cosmossdk.io/core/store"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/runtime"

	staking "github.com/cosmos/cosmos-sdk/x/staking/keeper"
)

const ModuleName = "da"

type Keeper struct {
	cdc           codec.BinaryCodec
	storeService  store.KVStoreService
	stakingKeeper staking.Keeper
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	stakingKeeper staking.Keeper,
) Keeper {
	return Keeper{
		cdc:           cdc,
		storeService:  storeService,
		stakingKeeper: stakingKeeper,
	}
}

func (k Keeper) HighestHeightWithMajority(ctx context.Context, heights map[uint64][][]byte) (uint64, error) {
	// TODO: Obviously this has to take into account the stake of the validators
	var last uint64
	for height := range heights {
		if height > last {
			last = height
		}
	}
	return last, nil
}

func (k Keeper) LatestDataCommitmentHeight(ctx context.Context) (uint64, error) {
	store := k.storeService.OpenKVStore(ctx)
	iter := storetypes.KVStoreReversePrefixIterator(runtime.KVStoreAdapter(store), []byte{DataCommitmentPrefix})
	defer iter.Close()
	if !iter.Valid() {
		return 0, iter.Error()
	}
	return ParseDataCommitmentKey(iter.Key()), iter.Error()
}

func (k Keeper) OldestDataCommitmentHeight(ctx context.Context) (uint64, error) {
	store := k.storeService.OpenKVStore(ctx)
	iter := storetypes.KVStorePrefixIterator(runtime.KVStoreAdapter(store), []byte{DataCommitmentPrefix})
	defer iter.Close()
	if !iter.Valid() {
		return 0, iter.Error()
	}
	return ParseDataCommitmentKey(iter.Key()), iter.Error()
}

func (k Keeper) GetDataCommitment(ctx context.Context, height uint64) ([]byte, error) {
	store := k.storeService.OpenKVStore(ctx)
	return store.Get(DataCommitmentKey(height))
}

func (k Keeper) HasDataCommitment(ctx context.Context, height uint64) (bool, error) {
	store := k.storeService.OpenKVStore(ctx)
	return store.Has(DataCommitmentKey(height))
}

func (k Keeper) SetDataCommitment(ctx context.Context, height uint64, dataCommitment []byte) error {
	store := k.storeService.OpenKVStore(ctx)
	return store.Set(DataCommitmentKey(height), dataCommitment)
}

func (k Keeper) DeleteDataCommitment(ctx context.Context, height uint64) error {
	store := k.storeService.OpenKVStore(ctx)
	return store.Delete(DataCommitmentKey(height))
}

const DataCommitmentPrefix = byte(0x01)

func DataCommitmentKey(height uint64) []byte {
	return binary.BigEndian.AppendUint64([]byte{DataCommitmentPrefix}, height)
}

func ParseDataCommitmentKey(key []byte) uint64 {
	return binary.BigEndian.Uint64(key[1:])
}
