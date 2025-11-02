package domain

import (
	"math"
)

var symbols = []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}

// FileSize represents the core domain entity of Bytey.
// It encapsulates the raw file size (in bytes) and provides
// human-friendly conversion logic.
type FileSize struct {
	Value int64
}

// Index returns the exponent index for the unit (KB, MB, GB...).
// For example:
// - 1 byte returns 0 (B)
// - 1024 bytes returns 1 (KB)
// - 1048576 bytes returns 2 (MB)
func (f FileSize) Index() float64 {
	if f.Value <= 0 {
		return 0
	}
	return math.Floor(math.Log(float64(f.Value)) / math.Log(1024))
}

// HumanValue returns the converted value (e.g., 1024 → 1.0 KB).
func (f FileSize) HumanValue() float64 {
	i := f.Index()
	return float64(f.Value) / math.Pow(1024, math.Floor(i))
}

// Symbol returns the appropriate symbol (e.g., KB, MB, GB).
func (f FileSize) Symbol() string {
	i := int(f.Index())
	if i >= len(symbols) {
		i = len(symbols) - 1
	}
	return symbols[i]
}
