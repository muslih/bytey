package domain

import (
	"testing"
)

func TestFileSize_Index(t *testing.T) {
	tests := []struct {
		name  string
		value int64
		want  float64
	}{
		{"zero bytes", 0, 0},
		{"1 byte", 1, 0},
		{"1023 bytes", 1023, 0},
		{"1 KB", 1024, 1},
		{"1 MB", 1024 * 1024, 2},
		{"1.5 MB", int64(1024 * 1024 * 1.5), 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := FileSize{Value: tt.value}
			if got := fs.Index(); got != tt.want {
				t.Errorf("Index() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFileSize_HumanValue(t *testing.T) {
	tests := []struct {
		name  string
		value int64
		want  float64
	}{
		{"1 byte", 1, 1},
		{"1 KB", 1024, 1},
		{"1.5 KB", 1024 + 512, 1.5},
		{"1 MB", 1024 * 1024, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := FileSize{Value: tt.value}
			got := fs.HumanValue()
			if got != tt.want {
				t.Errorf("HumanValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFileSize_Symbol(t *testing.T) {
	tests := []struct {
		name  string
		value int64
		want  string
	}{
		{"bytes", 500, "B"},
		{"kilobytes", 1024, "KB"},
		{"megabytes", 1024 * 1024, "MB"},
		{"gigabytes", 1024 * 1024 * 1024, "GB"},
		{"terabytes", 1024 * 1024 * 1024 * 1024, "TB"},
		{"petabytes", 1024 * 1024 * 1024 * 1024 * 1024, "PB"},
		{"exabytes", 1024 * 1024 * 1024 * 1024 * 1024 * 1024, "EB"},
		{"beyond exabytes", 1 << 62, "EB"}, // Should cap at EB
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := FileSize{Value: tt.value}
			if got := fs.Symbol(); got != tt.want {
				t.Errorf("Symbol() = %v, want %v", got, tt.want)
			}
		})
	}
}
