package da

import (
	"encoding/json"
	"errors"

	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/core/store"
	"cosmossdk.io/depinject"
	sdkerrors "cosmossdk.io/errors"
	v1 "github.com/Wondertan/da/modules/da/v1"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	sdkclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/module"
	staking "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	gwruntime "github.com/grpc-ecosystem/grpc-gateway/runtime"
	celestia "github.com/rollkit/celestia-openrpc"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ConsensusVersion defines the current da module consensus version.
const ConsensusVersion = 1

var (
	_ module.AppModuleBasic = AppModule{}
	_ module.HasGenesis     = AppModule{}

	_ appmodule.AppModule = AppModule{}
)

type AppModuleBasic struct{}

// Name returns the da module's name.
func (AppModuleBasic) Name() string {
	return ModuleName
}

// RegisterLegacyAminoCodec registers the da module's types on the LegacyAmino codec.
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the da module.
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *gwruntime.ServeMux) {
}

// RegisterInterfaces registers interfaces and implementations of the da module.
func (AppModuleBasic) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
}

// DefaultGenesis returns default genesis state as raw bytes for the da
// module.
func (ab AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(&v1.GenesisState{
		LatestCounterpartyHeight: 1,
	})
}

// ValidateGenesis performs genesis state validation for the da module.
func (ab AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, config sdkclient.TxEncodingConfig, bz json.RawMessage) error {
	var data v1.GenesisState
	if err := cdc.UnmarshalJSON(bz, &data); err != nil {
		return sdkerrors.Wrapf(err, "failed to unmarshal %s genesis state",ModuleName)
	}

	if data.LatestCounterpartyHeight == 0 {
		errors.New("latest counterparty height is zero")
	}

	return nil
}

// AppModule implements an application module for the da module.
type AppModule struct {
	AppModuleBasic

	keeper Keeper
	da     celestia.Client
}

// NewAppModule creates a new AppModule object
func NewAppModule(keeper Keeper, da celestia.Client) AppModule {
	return AppModule{keeper: keeper, da: da}
}

// IsOnePerModuleType implements the depinject.OnePerModuleType interface.
func (am AppModule) IsOnePerModuleType() {}

// IsAppModule implements the appmodule.AppModule interface.
func (am AppModule) IsAppModule() {}

// InitGenesis performs genesis initialization for the da module. It returns
// no validator updates.
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, bz json.RawMessage) {
	var gs v1.GenesisState
	cdc.MustUnmarshalJSON(bz, &gs)
	// TODO: set the initial height
}

// ExportGenesis returns the exported genesis state as raw bytes for the da
// module.
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	height, err := am.keeper.LatestDataCommitmentHeight(ctx)
	if err != nil {
		panic(err)
	}

	gs := &v1.GenesisState{
		LatestCounterpartyHeight: height,
	}

	return cdc.MustMarshalJSON(gs)
}

func init() {
	appmodule.Register(
		&v1.Module{},
		appmodule.Provide(ProvideModule),
	)
}

type ModuleInputs struct {
	depinject.In

	Cdc           codec.BinaryCodec
	StoreService  store.KVStoreService
	StakingKeeper staking.Keeper
	Opts          servertypes.AppOptions
}

type ModuleOutputs struct {
	depinject.Out

	DAKeeper Keeper
	Module   appmodule.AppModule
}

func ProvideModule(in ModuleInputs) ModuleOutputs {
	keeper := NewKeeper(
		in.Cdc,
		in.StoreService,
		in.StakingKeeper,
	)

	client, ok := in.Opts.Get("celestia-client").(celestia.Client)
	if !ok {
		panic("panic for now")
	}
	module := NewAppModule(keeper, client)

	return ModuleOutputs{DAKeeper: keeper, Module: module}
}
