package da

import (
	gwruntime "github.com/grpc-ecosystem/grpc-gateway/runtime"
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/core/store"
	"cosmossdk.io/depinject"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
	modulev1 "github.com/Wondertan/da/modules/da/v1"
)

// ConsensusVersion defines the current da module consensus version.
const ConsensusVersion = 1

var (
	_ module.AppModuleBasic = AppModule{}
	// _ module.HasGenesis     = AppModule{}

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

// AppModule implements an application module for the da module.
type AppModule struct {
	AppModuleBasic

	keeper Keeper
}

// NewAppModule creates a new AppModule object
func NewAppModule(keeper Keeper) AppModule {
	return AppModule{keeper: keeper}
}

// IsOnePerModuleType implements the depinject.OnePerModuleType interface.
func (am AppModule) IsOnePerModuleType() {}

// IsAppModule implements the appmodule.AppModule interface.
func (am AppModule) IsAppModule() {}

func init() {
	appmodule.Register(
		&modulev1.Module{},
		appmodule.Provide(ProvideModule),
	)
}

type ModuleInputs struct {
	depinject.In

	Cdc          codec.BinaryCodec
	StoreService store.KVStoreService
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
	)
	module := NewAppModule(keeper)

	return ModuleOutputs{DAKeeper: keeper, Module: module}
}
