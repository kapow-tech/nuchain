package audit

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "nuchain/api/nuchain/audit"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod:      "GetData",
					Use:            "get-data [datatype] [id]",
					Short:          "Query get-data",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "datatype"}, {ProtoField: "id"}},
				},

				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "PushData",
					Use:            "push-data [actor] [data-type] [id] [data] [timestamp]",
					Short:          "Send a push-data tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "actor"}, {ProtoField: "dataType"}, {ProtoField: "id"}, {ProtoField: "data"}, {ProtoField: "timestamp"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
