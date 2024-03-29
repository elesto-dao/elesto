package keeper_test

import (
	"fmt"

	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/elesto-dao/elesto/v4/x/credential"
	"github.com/elesto-dao/elesto/v4/x/credential/keeper"
)

var (
	schemaOkCompact = []uint8{0x7b, 0x22, 0x24, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x22, 0x3a, 0x22, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x2d, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x64, 0x72, 0x61, 0x66, 0x74, 0x2d, 0x30, 0x37, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x23, 0x22, 0x2c, 0x22, 0x24, 0x69, 0x64, 0x22, 0x3a, 0x22, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x69, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x72, 0x65, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x32, 0x30, 0x32, 0x30, 0x2f, 0x31, 0x2e, 0x30, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x2d, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x22, 0x2c, 0x22, 0x24, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x3a, 0x7b, 0x22, 0x73, 0x6c, 0x75, 0x67, 0x22, 0x3a, 0x22, 0x72, 0x65, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x32, 0x30, 0x32, 0x30, 0x22, 0x2c, 0x22, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x3a, 0x22, 0x31, 0x2e, 0x30, 0x22, 0x2c, 0x22, 0x69, 0x63, 0x6f, 0x6e, 0x22, 0x3a, 0x22, 0xf0, 0x9f, 0x85, 0xa1, 0x22, 0x2c, 0x22, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x3a, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x2c, 0x22, 0x75, 0x72, 0x69, 0x73, 0x22, 0x3a, 0x7b, 0x22, 0x6a, 0x73, 0x6f, 0x6e, 0x4c, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x50, 0x6c, 0x75, 0x73, 0x22, 0x3a, 0x22, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x69, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x72, 0x65, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x32, 0x30, 0x32, 0x30, 0x2f, 0x31, 0x2e, 0x30, 0x2f, 0x6c, 0x64, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2d, 0x70, 0x6c, 0x75, 0x73, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x22, 0x2c, 0x22, 0x6a, 0x73, 0x6f, 0x6e, 0x4c, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x22, 0x3a, 0x22, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x69, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x72, 0x65, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x32, 0x30, 0x32, 0x30, 0x2f, 0x31, 0x2e, 0x30, 0x2f, 0x6c, 0x64, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x22, 0x2c, 0x22, 0x6a, 0x73, 0x6f, 0x6e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x22, 0x3a, 0x22, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x69, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x72, 0x65, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x32, 0x30, 0x32, 0x30, 0x2f, 0x31, 0x2e, 0x30, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x2d, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x22, 0x7d, 0x7d, 0x2c, 0x22, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x3a, 0x22, 0x52, 0x65, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x32, 0x30, 0x32, 0x30, 0x22, 0x2c, 0x22, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3a, 0x22, 0x52, 0x65, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x32, 0x30, 0x32, 0x30, 0x20, 0x2d, 0x20, 0x41, 0x20, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2d, 0x70, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x20, 0x6d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x73, 0x6d, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x20, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x20, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x22, 0x2c, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x22, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x2c, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x3a, 0x5b, 0x22, 0x40, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x22, 0x2c, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x2c, 0x22, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x22, 0x2c, 0x22, 0x69, 0x73, 0x73, 0x75, 0x61, 0x6e, 0x63, 0x65, 0x44, 0x61, 0x74, 0x65, 0x22, 0x2c, 0x22, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x5d, 0x2c, 0x22, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x22, 0x3a, 0x7b, 0x22, 0x40, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x22, 0x3a, 0x7b, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x5b, 0x22, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x2c, 0x22, 0x61, 0x72, 0x72, 0x61, 0x79, 0x22, 0x2c, 0x22, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x5d, 0x7d, 0x2c, 0x22, 0x69, 0x64, 0x22, 0x3a, 0x7b, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x22, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x2c, 0x22, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x22, 0x3a, 0x22, 0x75, 0x72, 0x69, 0x22, 0x7d, 0x2c, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x7b, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x5b, 0x22, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x2c, 0x22, 0x61, 0x72, 0x72, 0x61, 0x79, 0x22, 0x5d, 0x2c, 0x22, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x3a, 0x7b, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x22, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x7d, 0x7d, 0x2c, 0x22, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x22, 0x3a, 0x7b, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x5b, 0x22, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x2c, 0x22, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x5d, 0x2c, 0x22, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x22, 0x3a, 0x22, 0x75, 0x72, 0x69, 0x22, 0x2c, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x3a, 0x5b, 0x22, 0x69, 0x64, 0x22, 0x5d, 0x2c, 0x22, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x22, 0x3a, 0x7b, 0x22, 0x69, 0x64, 0x22, 0x3a, 0x7b, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x22, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x2c, 0x22, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x22, 0x3a, 0x22, 0x75, 0x72, 0x69, 0x22, 0x7d, 0x7d, 0x7d, 0x2c, 0x22, 0x69, 0x73, 0x73, 0x75, 0x61, 0x6e, 0x63, 0x65, 0x44, 0x61, 0x74, 0x65, 0x22, 0x3a, 0x7b, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x22, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x2c, 0x22, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x22, 0x3a, 0x22, 0x64, 0x61, 0x74, 0x65, 0x2d, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x7d, 0x2c, 0x22, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x22, 0x3a, 0x7b, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x22, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x2c, 0x22, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x22, 0x3a, 0x22, 0x64, 0x61, 0x74, 0x65, 0x2d, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x7d, 0x2c, 0x22, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x3a, 0x7b, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x22, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x2c, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x3a, 0x5b, 0x22, 0x69, 0x64, 0x22, 0x2c, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x2c, 0x22, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x5d, 0x2c, 0x22, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x22, 0x3a, 0x7b, 0x22, 0x69, 0x64, 0x22, 0x3a, 0x7b, 0x22, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x3a, 0x22, 0x49, 0x44, 0x22, 0x2c, 0x22, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3a, 0x22, 0x54, 0x68, 0x65, 0x20, 0x72, 0x65, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x49, 0x44, 0x22, 0x2c, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x22, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x2c, 0x22, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x22, 0x3a, 0x22, 0x75, 0x72, 0x69, 0x22, 0x7d, 0x2c, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x7b, 0x22, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x3a, 0x22, 0x54, 0x79, 0x70, 0x65, 0x22, 0x2c, 0x22, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3a, 0x22, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x20, 0x73, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x20, 0x62, 0x65, 0x3a, 0x20, 0x52, 0x65, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x32, 0x30, 0x32, 0x30, 0x20, 0x22, 0x2c, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x22, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x7d, 0x2c, 0x22, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x3a, 0x7b, 0x22, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x3a, 0x22, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x2c, 0x22, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3a, 0x22, 0x62, 0x61, 0x73, 0x65, 0x36, 0x34, 0x20, 0x65, 0x6e, 0x64, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x20, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x7a, 0x6c, 0x69, 0x62, 0x20, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x65, 0x64, 0x20, 0x62, 0x69, 0x74, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x2c, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x22, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x7d, 0x7d, 0x7d, 0x2c, 0x22, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x22, 0x3a, 0x7b, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x22, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x2c, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x3a, 0x5b, 0x22, 0x69, 0x64, 0x22, 0x2c, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x5d, 0x2c, 0x22, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x22, 0x3a, 0x7b, 0x22, 0x69, 0x64, 0x22, 0x3a, 0x7b, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x22, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x2c, 0x22, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x22, 0x3a, 0x22, 0x75, 0x72, 0x69, 0x22, 0x7d, 0x2c, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x7b, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x22, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x7d, 0x7d, 0x7d, 0x7d, 0x7d}
	vocabOkCompact  = []uint8{0x7b, 0x22, 0x40, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x22, 0x3a, 0x7b, 0x22, 0x77, 0x33, 0x63, 0x63, 0x72, 0x65, 0x64, 0x22, 0x3a, 0x22, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x77, 0x33, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x32, 0x30, 0x31, 0x38, 0x2f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x23, 0x22, 0x2c, 0x22, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2d, 0x69, 0x64, 0x22, 0x3a, 0x22, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x69, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x72, 0x65, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x32, 0x30, 0x32, 0x30, 0x2f, 0x31, 0x2e, 0x30, 0x2f, 0x6c, 0x64, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x23, 0x22, 0x2c, 0x22, 0x52, 0x65, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x32, 0x30, 0x32, 0x30, 0x22, 0x3a, 0x7b, 0x22, 0x40, 0x69, 0x64, 0x22, 0x3a, 0x22, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2d, 0x69, 0x64, 0x22, 0x7d, 0x2c, 0x22, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x3a, 0x7b, 0x22, 0x40, 0x69, 0x64, 0x22, 0x3a, 0x22, 0x77, 0x33, 0x63, 0x63, 0x72, 0x65, 0x64, 0x3a, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x2c, 0x22, 0x40, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x22, 0x3a, 0x7b, 0x22, 0x69, 0x64, 0x22, 0x3a, 0x7b, 0x22, 0x40, 0x69, 0x64, 0x22, 0x3a, 0x22, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2d, 0x69, 0x64, 0x3a, 0x69, 0x64, 0x22, 0x2c, 0x22, 0x40, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x22, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x54, 0x65, 0x78, 0x74, 0x22, 0x7d, 0x2c, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x7b, 0x22, 0x40, 0x69, 0x64, 0x22, 0x3a, 0x22, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2d, 0x69, 0x64, 0x3a, 0x74, 0x79, 0x70, 0x65, 0x22, 0x2c, 0x22, 0x40, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x22, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x54, 0x65, 0x78, 0x74, 0x22, 0x7d, 0x2c, 0x22, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x3a, 0x7b, 0x22, 0x40, 0x69, 0x64, 0x22, 0x3a, 0x22, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2d, 0x69, 0x64, 0x3a, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x2c, 0x22, 0x40, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x22, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x54, 0x65, 0x78, 0x74, 0x22, 0x7d, 0x7d, 0x7d, 0x2c, 0x22, 0x40, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x3a, 0x31, 0x2e, 0x31, 0x7d, 0x7d}
)

func (suite *KeeperTestSuite) Test_HandlePublicProposalChange() {
	testCases := []struct {
		name    string
		reqFn   func() govtypes.Content
		wantErr error
	}{
		{
			name: "PASS: ID and Proposal valid",
			reqFn: func() govtypes.Content {
				cd := &credential.CredentialDefinition{
					Id:           "did:cosmos:elesto:cd-1",
					PublisherId:  "did:cosmos:elesto:publisher",
					Schema:       schemaOkCompact,
					Vocab:        vocabOkCompact,
					Name:         "Credential Definition 1",
					Description:  "This is a sample credential",
					SupersededBy: "",
					IsActive:     true,
				}
				suite.keeper.SetCredentialDefinition(suite.ctx, cd)

				return credential.NewProposePublicCredentialID("TEST", "TEST", cd.Id)
			},
			wantErr: nil,
		},
		{
			name: "INVALID: Id does not exist",
			reqFn: func() govtypes.Content {
				return credential.NewProposePublicCredentialID("TEST", "TEST", "did:cosmos:elesto:cd-2")
			},
			wantErr: fmt.Errorf("proposal with %s id not found", "did:cosmos:elesto:cd-2"),
		},

		{
			name: "INVALID: ID already allowed",
			reqFn: func() govtypes.Content {
				cd := &credential.CredentialDefinition{
					Id:           "did:cosmos:elesto:cd-3",
					PublisherId:  "did:cosmos:elesto:publisher",
					Schema:       schemaOkCompact,
					Vocab:        vocabOkCompact,
					Name:         "Credential Definition 3",
					Description:  "This is a sample credential",
					SupersededBy: "",
					IsActive:     true,
				}
				suite.keeper.SetCredentialDefinition(suite.ctx, cd)
				suite.keeper.AllowPublicCredential(suite.ctx, cd.Id)

				return credential.NewProposePublicCredentialID("TEST", "TEST", cd.Id)
			},
			wantErr: fmt.Errorf("credential definition with id %s is already public", "did:cosmos:elesto:cd-3"),
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			req := tc.reqFn()
			err := keeper.NewPublicCredentialProposalHandler(suite.keeper)(suite.ctx, req)
			if tc.wantErr == nil {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
				suite.Require().Equal(tc.wantErr.Error(), err.Error())
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_HandleRemovePublicProposalChange() {
	testCases := []struct {
		name    string
		reqFn   func() govtypes.Content
		wantErr error
	}{
		{
			name: "PASS: ID and Proposal valid",
			reqFn: func() govtypes.Content {
				cd := &credential.CredentialDefinition{
					Id:           "did:cosmos:elesto:cd-1",
					PublisherId:  "did:cosmos:elesto:publisher",
					Schema:       schemaOkCompact,
					Vocab:        vocabOkCompact,
					Name:         "Credential Definition 1",
					Description:  "This is a sample credential",
					SupersededBy: "",
					IsActive:     true,
				}
				suite.keeper.SetCredentialDefinition(suite.ctx, cd)
				suite.keeper.AllowPublicCredential(suite.ctx, cd.Id)

				return credential.NewProposeRemovePublicCredentialID("TEST", "TEST", cd.Id)
			},
			wantErr: nil,
		},
		{
			name: "INVALID: Id does not exist",
			reqFn: func() govtypes.Content {
				return credential.NewProposeRemovePublicCredentialID("TEST", "TEST", "did:cosmos:elesto:cd-2")
			},
			wantErr: fmt.Errorf("proposal with %s id not found", "did:cosmos:elesto:cd-2"),
		},

		{
			name: "INVALID: ID is not allowed",
			reqFn: func() govtypes.Content {
				cd := &credential.CredentialDefinition{
					Id:           "did:cosmos:elesto:cd-3",
					PublisherId:  "did:cosmos:elesto:publisher",
					Schema:       schemaOkCompact,
					Vocab:        vocabOkCompact,
					Name:         "Credential Definition 3",
					Description:  "This is a sample credential",
					SupersededBy: "",
					IsActive:     true,
				}
				suite.keeper.SetCredentialDefinition(suite.ctx, cd)

				return credential.NewProposeRemovePublicCredentialID("TEST", "TEST", cd.Id)
			},
			wantErr: fmt.Errorf("credential definition with id %s is not public", "did:cosmos:elesto:cd-3"),
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			req := tc.reqFn()
			err := keeper.NewPublicCredentialProposalHandler(suite.keeper)(suite.ctx, req)
			if tc.wantErr == nil {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
				suite.Require().Equal(tc.wantErr.Error(), err.Error())
			}
		})
	}
}
