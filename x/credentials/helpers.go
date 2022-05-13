package credentials

import (
	"bytes"
	"encoding/json"
	"os"
	"sort"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/gogo/protobuf/proto"
)

// IsEmpty checks if a string is empty, it trims spaces before checking for empty string
func IsEmpty(v string) bool {
	return strings.TrimSpace(v) == ""
}

// CompactJSON read a JSON from a file and return a compact version of it
func CompactJSON(filePath string) (compact string, err error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return
	}
	cp := bytes.NewBuffer([]byte{})
	if err = json.Compact(cp, data); err != nil {
		return
	}
	compact = cp.String()
	return
}

// LoadProtoJSON read JSON data from a file and un-marshall it to a struct
func LoadProtoJSON(cdc codec.Codec, filePath string, to proto.Message) (err error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return
	}
	return cdc.UnmarshalJSON(data, to)
}

// StringUnion perform union, distinct amd sort operation between two slices
// duplicated element in list are removed
func StringUnion(a, b []string) []string {
	if len(b) == 0 {
		return a
	}
	m := make(map[string]struct{})
	for _, item := range a {
		m[item] = struct{}{}
	}
	for _, item := range b {
		if _, ok := m[item]; !ok {
			m[item] = struct{}{}
		}
	}
	u := make([]string, 0, len(m))
	for k := range m {
		u = append(u, k)
	}
	sort.Strings(u)
	return u
}
