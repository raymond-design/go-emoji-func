package emoji

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestGetEmojiDefs(t *testing.T) {
  assert.Equal(t, "🚮", GetEmojiDefinitions()["litter-in-bin-sign"])
  assert.Equal(t, "🪫", GetEmojiDefinitions()["low-battery"])
}

func TestStringToEmoji(t *testing.T) {
	assert.Equal(t, "🚮", stringToEmoji("litter-in-bin-sign"))
	assert.Equal(t, "🪫", stringToEmoji("low-battery"))
}

func TestConvFuncBack(t *testing.T) {
	assert.Equal(t, "1️⃣", ConvFuncBack("Ukeycap_1"))
	assert.Equal(t, "🪫🪫", ConvFuncBack("Ulow-batteryUlow-battery"))
	assert.Equal(t, "🪫1️⃣🪫🚮", ConvFuncBack("Ulow-batteryUkeycap_1Ulow-batteryUlitter-in-bin-sign"))
}