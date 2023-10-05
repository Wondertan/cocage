package module

import (
	"encoding/binary"

	"cosmossdk.io/core/store"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const ModuleName = "da"

type Keeper struct {
	cdc          codec.BinaryCodec
	storeService store.KVStoreService
}

func NewKeeper(cdc codec.BinaryCodec, storeService store.KVStoreService) Keeper {
	return Keeper{cdc: cdc, storeService: storeService}
}

func (k Keeper) LatestCommitmentHeight(ctx sdk.Context) (int64, error) {
	store := k.storeService.OpenKVStore(ctx)
	iter := storetypes.KVStoreReversePrefixIterator(runtime.KVStoreAdapter(store), []byte{DataCommitmentPrefix})
	defer iter.Close()
	if !iter.Valid() {
		return 0, iter.Error()
	}
	return ParseDataCommitmentKey(iter.Key()), iter.Error()
}

func (k Keeper) OldestCommitmentHeight(ctx sdk.Context) (int64, error) {
	store := k.storeService.OpenKVStore(ctx)
	iter := storetypes.KVStorePrefixIterator(runtime.KVStoreAdapter(store), []byte{DataCommitmentPrefix})
	defer iter.Close()
	if !iter.Valid() {
		return 0, iter.Error()
	}
	return ParseDataCommitmentKey(iter.Key()), iter.Error()
}

func (k Keeper) GetDataCommitment(ctx sdk.Context, height int64) ([]byte, error) {
	store := k.storeService.OpenKVStore(ctx)
	return store.Get(DataCommitmentKey(height))
}

func (k Keeper) HasDataCommitment(ctx sdk.Context, height int64) (bool, error) {
	store := k.storeService.OpenKVStore(ctx)
	return store.Has(DataCommitmentKey(height))
}

func (k Keeper) SetDataCommitment(ctx sdk.Context, height int64, dataCommitment []byte) error {
	store := k.storeService.OpenKVStore(ctx)
	return store.Set(DataCommitmentKey(height), dataCommitment)
}

const DataCommitmentPrefix = byte(0x01)

func DataCommitmentKey(height int64) []byte {
	return binary.BigEndian.AppendUint64([]byte{DataCommitmentPrefix}, uint64(height))
}

func ParseDataCommitmentKey(key []byte) int64 {
	return int64(binary.BigEndian.Uint64(key[1:]))
}