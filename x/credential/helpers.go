package credential

import (
	"bytes"
	"encoding/json"
	"os"
	"strings"
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
