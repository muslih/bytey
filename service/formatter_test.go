package service

import (
	"bytey/domain"
	"testing"
)

func TestFormatter_FormatFormal(t *testing.T) {
	tests := []struct {
		name  string
		value int64
		want  string
	}{
		{"zero bytes", 0, "0B"},
		{"1 byte", 1, "1B"},
		{"1023 bytes", 1023, "1023B"},
		{"1 KB", 1024, "1KB"},
		{"1.5 KB", 1536, "1.5KB"},
		{"1 MB", 1024 * 1024, "1MB"},
		{"1.2 MB", 1258291, "1.2MB"}, // 1.2 * 1024 * 1024 = 1,258,291.2
	}

	f := &Formatter{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f.FormatFormal(domain.FileSize{Value: tt.value}); got != tt.want {
				t.Errorf("FormatFormal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatter_FormatPlayful(t *testing.T) {
	tests := []struct {
		name  string
		value int64
		// We can't test the exact output because of the random reaction,
		// but we can check the prefix and that it contains a reaction
		expectedPrefix string
	}{
		{"negative bytes", -1, "Negative bytes?"},
		{"zero bytes", 0, "0B — "},
		{"1 byte", 1, "1B — "},
		{"1 KB", 1024, "1KB — "},
	}

	f := &Formatter{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := f.FormatPlayful(domain.FileSize{Value: tt.value})
			if tt.name == "negative bytes" {
				if got != "Negative bytes? Are you reversing time, human? ⏳" {
					t.Errorf("FormatPlayful() = %v, want prefix %v", got, tt.expectedPrefix)
				}
			} else if len(got) <= len(tt.expectedPrefix) || got[:len(tt.expectedPrefix)] != tt.expectedPrefix {
				t.Errorf("FormatPlayful() = %v, want prefix %v", got, tt.expectedPrefix)
			}
		})
	}
}
