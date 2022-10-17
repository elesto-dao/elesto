package types

import (
	"reflect"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestGetInflationRate(t *testing.T) {
	tests := []struct {
		name   string
		height int64
		want   sdk.Dec
	}{
		{
			name:   "epoch-1",
			height: 10,
			want:   sdk.NewDec(100),
		},
		{
			name:   "epoch-2",
			height: BlocksPerEpoch,
			want:   sdk.NewDec(50),
		},
		{
			name:   "epoch-3",
			height: BlocksPerEpoch * 2,
			want:   sdk.MustNewDecFromStr("25"),
		},
		{
			name:   "epoch-4",
			height: BlocksPerEpoch * 3,
			want:   sdk.MustNewDecFromStr("12.5"),
		},
		{
			name:   "epoch-5",
			height: BlocksPerEpoch * 4,
			want:   sdk.MustNewDecFromStr("6.24"),
		},
		{
			name:   "epoch-6",
			height: BlocksPerEpoch * 5,
			want:   sdk.MustNewDecFromStr("3.12"),
		},
		{
			name:   "epoch-7",
			height: BlocksPerEpoch * 6,
			want:   sdk.MustNewDecFromStr("2"),
		},
		{
			name:   "epoch-8",
			height: BlocksPerEpoch * 7,
			want:   sdk.MustNewDecFromStr("1.99"),
		},
		{
			name:   "epoch-9",
			height: BlocksPerEpoch * 8,
			want:   sdk.MustNewDecFromStr("1.99"),
		},
		{
			name:   "epoch-10",
			height: BlocksPerEpoch * 9,
			want:   sdk.MustNewDecFromStr("1.92"),
		},
		{
			name:   "epoch-11",
			height: BlocksPerEpoch * 10,
			want:   sdk.NewDec(0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetInflationRate(tt.height); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetInflation() = %v, want %v", got, tt.want)
			}
		})
	}
}
