// private utility functions
package emoji

import (
    "github.com/forPelevin/gomoji"
	"strings"
	//"fmt"
)

// input must be one emoji, in string format
// replaces '-' with "_"
func getEmojiSlug(emj string) string {
	info, err := gomoji.GetInfo(emj)
	if err != nil {
		// internal error, input wasn't an emoji
		//fmt.Errorf("Error: Internal Error, input was not emoji %v", emj)
	}

	return strings.Replace(info.Slug, "-", "_", -1)
}

// check if string has emoji
// return err if emoji invalid
func hasEmoji(emj string) bool {
	return gomoji.ContainsEmoji(emj)
}