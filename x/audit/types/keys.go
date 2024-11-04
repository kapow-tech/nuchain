package types

import (
	ormapi "cosmossdk.io/api/cosmos/orm/v1alpha1"

	"nuchain/api/nuchain/audit"
)

const (
	// ModuleName defines the module name
	ModuleName = "audit"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_audit"
)

var (
	ParamsKey   = []byte("p_audit")
	AuditSchema = &ormapi.ModuleSchemaDescriptor{
		SchemaFile: []*ormapi.ModuleSchemaDescriptor_FileEntry{
			{
				Id:            1,
				ProtoFileName: audit.File_nuchain_audit_state_proto.Path(),
			},
		},
	}
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
