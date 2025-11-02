package bytey

import (
	"bytey/domain"
	"bytey/service"
)

// ByteyFacade is the main entry point of the package.
// It exposes friendly methods to talk to Bytey — your playful size formatter.
type ByteyFacade struct {
	formatter *service.Formatter
}

// NewBytey initializes a new Bytey instance.
func NewBytey() *ByteyFacade {
	return &ByteyFacade{
		formatter: service.NewFormatter(),
	}
}

// Talk returns a human-readable size string with a playful random comment.
// Example: "1.0MB — File’s been hitting the buffet 🍝"
func (b *ByteyFacade) Talk(size int64) string {
	return b.formatter.FormatPlayful(domain.FileSize{Value: size})
}

// Size returns a formal, clean human-readable size string.
// Example: "1.0MB"
func (b *ByteyFacade) Size(size int64) string {
	return b.formatter.FormatFormal(domain.FileSize{Value: size})
}

// Whisper returns a quiet, gentle version of Size with a random emoji.
// Perfect for aesthetic output.
func (b *ByteyFacade) Whisper(size int64) string {
	result := b.formatter.FormatFormal(domain.FileSize{Value: size})
	emojis := []string{"✨", "🌸", "💫", "🌙", "🍃"}
	return result + " " + emojis[int(size)%len(emojis)]
}

// Scream returns an exaggerated, dramatic output —
// great for large files or when Bytey’s feeling extra expressive.
func (b *ByteyFacade) Scream(size int64) string {
	formatted := b.formatter.FormatPlayful(domain.FileSize{Value: size})
	return "🚨 BYTEY ALERT 🚨 " + formatted
}
