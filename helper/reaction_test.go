package helper

import (
	"strings"
	"testing"
)

func TestRandomReaction(t *testing.T) {
	// Test that RandomReaction returns a valid reaction
	reaction := RandomReaction()

	// Check that the reaction is not empty
	if reaction == "" {
		t.Error("RandomReaction() should not return empty string")
	}

	// Check that the reaction is one of the known reactions
	found := false
	for _, validReaction := range reactions {
		if reaction == validReaction {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("RandomReaction() returned unexpected reaction: %s", reaction)
	}
}

func TestRandomReactionVariety(t *testing.T) {
	// Test that RandomReaction returns different values over multiple calls
	// This is probabilistic, but with 8 reactions, we should see variety
	reactions := make(map[string]bool)

	// Call RandomReaction multiple times
	for i := 0; i < 50; i++ {
		reaction := RandomReaction()
		reactions[reaction] = true
	}

	// We should see at least 2 different reactions (very low bar due to randomness)
	if len(reactions) < 2 {
		t.Error("RandomReaction() should return variety over multiple calls")
	}
}

func TestReactionsSliceNotEmpty(t *testing.T) {
	// Test that the reactions slice has content
	if len(reactions) == 0 {
		t.Error("reactions slice should not be empty")
	}
}

func TestReactionsSliceContent(t *testing.T) {
	// Test that each reaction in the slice is not empty
	for i, reaction := range reactions {
		if reaction == "" {
			t.Errorf("reactions[%d] should not be empty", i)
		}
	}
}

func TestReactionsSliceExpectedContent(t *testing.T) {
	// Test that the reactions slice contains the expected number of reactions
	expectedCount := 8

	if len(reactions) != expectedCount {
		t.Errorf("Expected %d reactions, got %d", expectedCount, len(reactions))
	}

	// Check that each reaction contains expected keywords (more robust than exact matching)
	expectedKeywords := []string{
		"big one",           // for "Whoa, that's a big one 💪"
		"byte-sized snack",  // for "Just a byte-sized snack 🍪"
		"buffet",            // for "This file's been hitting the buffet 🍝"
		"email attachments", // for "Too heavy for email attachments 😅"
		"Much data",         // for "Much data, very wow 🐶"
		"SSD just sighed",   // for "I think your SSD just sighed 💨"
		"Baby file",         // for "Baby file 👶 - so cute!"
		"Time traveler",     // for "Negative bytes? Time traveler alert ⏳" (capital T)
	}

	for _, keyword := range expectedKeywords {
		found := false
		for _, reaction := range reactions {
			if strings.Contains(reaction, keyword) {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("No reaction found containing keyword: %s", keyword)
			t.Logf("Available reactions: %v", reactions)
		}
	}
}
