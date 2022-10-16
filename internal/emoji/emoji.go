// main emoji checking logic
// convert into unicode
package emoji

import (
	"github.com/forPelevin/gomoji"
	"strings"
	//"fmt"
)

// check if emoji in func header and conv into unicode equivalent
// input is the function header:
// ex 1: 
// func 🦋() {
//
// ex 2:
// func (test *Test) 🦋() (string, error){
//
func ConvFuncHeader(def string) {

}

// conv already-converted emoji header back into original emoji
// Ukeycap_1 needs to be converted back into 1️⃣ 
// Ukeycap_1Ukeycap_1 needs to be converted back into 1️⃣1️⃣ 
func ConvFuncBack(def string) string {
	strings.Replace(def, "U", " ", -1)
	strings.Replace(def, "_", "-", -1)
	emojis := strings.Fields(def)

	original := ""
	for _, x := range emojis {
		original += stringToEmoji(x)
	}

	return original
}

// must pass string in valid slug form
func stringToEmoji(emj string) string {
	return GetEmojiDefinitions()[emj]
}

// get globally defined emoji hashmap
func GetEmojiDefinitions() map[string]string {
	emojilist := make(map[string]string)
	emojis := gomoji.AllEmojis()
	for _, x:= range emojis {
		emojilist[x.Slug] = x.Character
		//fmt.Printf("%v%v\n", x.Slug, x.Character)
	}  
	return emojilist
}