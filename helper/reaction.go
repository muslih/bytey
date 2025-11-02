package helper

import (
	"math/rand"
	"time"
)

// Bytey sometimes speaks up with a funny comment.
// These are random, lighthearted reactions for large or small files.
var reactions = []string{
	"Whoa, that’s a big one 💪",
	"Just a byte-sized snack 🍪",
	"This file’s been hitting the buffet 🍝",
	"Too heavy for email attachments 😅",
	"Much data, very wow 🐶",
	"I think your SSD just sighed 💨",
	"Baby file 👶 - so cute!",
	"Negative bytes? Time traveler alert ⏳",
}

// RandomReaction picks a random comment from Bytey’s mood board.
func RandomReaction() string {
	rand.Seed(time.Now().UnixNano())
	return reactions[rand.Intn(len(reactions))]
}
