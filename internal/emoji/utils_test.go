package emoji

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestGetEmojiSlug(t *testing.T) {
  assert.Equal(t, "litter_in_bin_sign", getEmojiSlug("🚮"))
  assert.Equal(t, "low_battery", getEmojiSlug("🪫"))
}

func TestStringHasEmoji(t *testing.T) {
	assert.Equal(t, true, hasEmoji("🚮"))
	assert.Equal(t, false, hasEmoji("dude"))
	assert.Equal(t, true, hasEmoji("battery: 🪫"))
}
