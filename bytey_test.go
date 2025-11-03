package bytey_test

import (
	"testing"

	"github.com/muslih/bytey"
)

func TestByteyFacade(t *testing.T) {
	tests := []struct {
		name  string
		value int64
		fn    func(*bytey.ByteyFacade, int64) string
		// We'll just check the prefix for methods that include random elements
		wantPrefix string
	}{
		{"Talk - 1 byte", 1, (*bytey.ByteyFacade).Talk, "1B — "},
		{"Size - 1 byte", 1, (*bytey.ByteyFacade).Size, "1B"},
		{"Whisper - 1 byte", 1, (*bytey.ByteyFacade).Whisper, "1B "},
		{"Scream - 1 byte", 1, (*bytey.ByteyFacade).Scream, "🚨 BYTEY ALERT 🚨 1B — "},
		{"Talk - 1 KB", 1024, (*bytey.ByteyFacade).Talk, "1KB — "},
		{"Size - 1 KB", 1024, (*bytey.ByteyFacade).Size, "1KB"},
		{"Whisper - 1 KB", 1024, (*bytey.ByteyFacade).Whisper, "1KB "},
		{"Scream - 1 KB", 1024, (*bytey.ByteyFacade).Scream, "🚨 BYTEY ALERT 🚨 1KB — "},
	}

	b := bytey.NewBytey()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fn(b, tt.value)
			if len(got) < len(tt.wantPrefix) || got[:len(tt.wantPrefix)] != tt.wantPrefix {
				t.Errorf("%s() = %v, want prefix %v", tt.name, got, tt.wantPrefix)
			}
		})
	}
}

func TestNegativeSize(t *testing.T) {
	b := bytey.NewBytey()
	t.Run("Negative size", func(t *testing.T) {
		if got := b.Talk(-1); got != "Negative bytes? Are you reversing time, human? ⏳" {
			t.Error("Expected negative size to return specific error message")
		}
	})
}
