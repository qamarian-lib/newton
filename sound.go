package newton

import (
	"crypto/rand"
	"github.com/qamarian-dtp/err"
)

// Function Sound () provides a random sound. The type of sound you want can be optionally
// supplied as an argument. For a consonant sound: use 'c', and for a vowel sound: use 'v'.
//
//	sound, err := soundPro ('c')
//
// If you are not bothered about the sound being consonant or vowel, do not supply any
// argument.
//
func Sound (soundType ... rune) (sound rune, e error) {
	// Input validation. ..1.. {
	if len (soundType) > 1 {
		e = err.New ("More than one sound type specified.", nil, nil)
		return
	}

	if len (soundType) == 1 && ((soundType [0] != 'c') && (soundType [0] != 'v')) {
		e = err.New ("Invalid sound type.", nil, nil)
		return
	}
	// ..1.. }

	sounds := []rune {}
	if len (soundType) == 1 && soundType [0] == 'c' {
		sounds = sound_cons
	} else if len (soundType) == 1 && soundType [0] == 'v' {
		sounds = sound_vowl
	} else {
		sounds = sound_all
	}

	randByte := make ([]byte, 1)
	_, errX := rand.Read (randByte)
	if errX != nil {
		e = err.New ("Unable to source a random byte, for deciding what sound " +
			"to choose.", nil, nil, errX)
		return
	}
	chosenSound := int (randByte [0]) % len (sounds)

	sound = sounds [chosenSound]
	return
}

var (
	sound_cons []rune = []rune (Cons)
	sound_vowl []rune = []rune (Vowl)
	sound_all  []rune
)
func init () {
	sound_all = append (sound_cons, sound_vowl...)
}
