package service

import (
	"fmt"

	"bytey/domain"
	"bytey/helper"
)

// Formatter handles the transformation of a FileSize
// into a human-readable or playful representation.
type Formatter struct{}

// NewFormatter returns a new instance of the Formatter service.
func NewFormatter() *Formatter {
	return &Formatter{}
}

// FormatFormal produces a clean, human-readable size string
// (e.g. "1.2MB") without any extra humor or commentary.
func (f *Formatter) FormatFormal(size domain.FileSize) string {
	if size.Value < 10 {
		return fmt.Sprintf("%dB", size.Value)
	}

	human := size.HumanValue()
	format := "%.0f"
	// Only use decimal for non-whole numbers
	if human != float64(int64(human)) && human < 10 {
		format = "%.1f"
	}

	return fmt.Sprintf(format+"%s", human, size.Symbol())
}

// FormatPlayful produces a human-readable size string
// but with Bytey’s random reactions for extra personality.
func (f *Formatter) FormatPlayful(size domain.FileSize) string {
	if size.Value < 0 {
		return "Negative bytes? Are you reversing time, human? ⏳"
	}

	formal := f.FormatFormal(size)
	reaction := helper.RandomReaction()
	return fmt.Sprintf("%s — %s", formal, reaction)
}
