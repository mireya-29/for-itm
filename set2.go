// Placeholder
package main

import (
"fmt"
)

func main() {
	// Placeholder
}

// Shift letter
//
// Shift a letter right by the given number.
// Wrap the letter around if it reaches the end of the alphabet.
//
// Examples
// shiftLetter("A", 0) -> "A"
// shiftLetter("A", 2) -> "C"
// shiftLetter("Z", 1) -> "A"
// shiftLetter("X", 5) -> "C"
// shiftLetter(" ", _) -> " "
//
// Note: we use the underscore `_` to denote the presence of a value that is present but irrelevant.
//
// Params:
// - letter, a single uppercase English letter, or a space
// - shift, the number by which to shift the letter
//
// Returns:
// - the letter, shifted appropriately, if a letter; otherwise a space.
func shiftLetter(letter string, shift int) string {
        alphabetSequence := map[string]int{
         "A": 1,
         "B": 2,
         "C": 3,
         "D": 4,
         "E": 5,
         "F": 6,
         "G": 7,
         "H": 8,
         "I": 9,
         "J": 10,
         "K": 11,
         "L": 12,
         "M": 13,
         "N": 14,
         "O": 15,
         "P": 16,
         "Q": 17,
         "R": 18,
         "S": 19,
         "T": 20,
         "U": 21,
         "V": 22,
         "W": 23,
         "X": 24,
         "Y": 25,
         "Z": 26,
        }
        currentPosition := alphabetSequence[letter]
        shiftedLetter := (currentPosition + shift - 1) % 26 + 1
        for l, v := range alphabetSequence{
                if v == shiftedLetter{
                return l}
        }
        return ""
}
//for test runs
  //func main() {
        //myLetter := shiftLetter("A", 2)
        //fmt.Println("Shifted Letter: ", myLetter)
//}




// Caesar cipher
//
// Apply a shift number to a string of uppercase English letters and spaces.
//
// Params:
// - message, a string of uppercase English letters and spaces
// - shift, the number by which to shift the letters
//
// Returns
// - the message, shifted appropriately
func caesarCipher(message string, shift int) string {
	func caesarCipher(message string, shift int) string {
        answer := ""
        for _, letter := range message {
                if letter >= 'A' && letter <= 'Z' {
                        shiftedLetter := (letter-'A'+rune(shift))%26 + 'A'
                        answer += string(shiftedLetter)
                } else {
                        answer += string(letter)
                }
        }
        return answer
}
//for test runs
//func main() {
       // myMessage := "HELLO WORLD"
      //  shift := 3
      //  newMessage := caesarCipher(myMessage, shift)
      //  fmt.Println("New Message:", newMessage)
}


// Shift by letter
//
// Shift a letter to the right using the number equivalent of another letter.
// The shift letter is any letter from A to Z, where A represents 0, B represents 1, ..., Z represents 25.
//
// Examples
// shiftByLetter("A", "A") -> "A"
// shiftByLetter("A", "C") -> "C"
// shiftByLetter("B", "K") -> "L"
// shiftByLetter(" ", _) -> " "
//
// Params:
// - letter, a single uppercase English letter, or a space
// - letterShift, a single uppercase English letter
//
// Returns:
// - the letter, shifted appropriately
func shiftByLetter(letter string, letterShift string) string {
        alphabetSequence := map[string]int{
         "A": 0,
         "B": 1,
         "C": 2,
         "D": 3,
         "E": 4,
         "F": 5,
         "G": 6,
         "H": 7,
         "I": 8,
         "J": 9,
         "K": 10,
         "L": 11,
         "M": 12,
         "N": 13,
         "O": 14,
         "P": 15,
         "Q": 16,
         "R": 17,
         "S": 18,
         "T": 19,
         "U": 20,
         "V": 21,
         "W": 22,
         "X": 23,
         "Y": 24,
         "Z": 25,
        }
        
        if letter == " "{
	return " "
        }

        firstLetterPosition := alphabetSequence[letter]
        secondLetterPosition := alphabetSequence[letterShift]

        shiftedLetter := (firstLetterPosition + secondLetterPosition ) % 26

        for l, v := range alphabetSequence{
                if v == shiftedLetter{
                return l
        }
        return ""
}
//for test runs
  //func main() {
       // myLetter := shiftByLetter("B", "K")
      //  fmt.Println("Shifted Letter: ", myLetter)
//}




// Vigenere cipher
//
// Encrypt a message using a keyphrase instead of a static number.
// Every letter in the message is shifted by the number represented by the respective letter in the key.
// Spaces are ignored.
//
// Example
// vigenereCipher("A C", "KEY") -> "K A"
//
// If needed, the keyphrase is extended to match the length of the key.
// If the key is "KEY" and the message is "LONGTEXT", the key will be extended to "KEYKEYKE".
//
// Params:
// - message, a string of uppercase english letters and/or spaces
// - key, a string of uppercase English letters with no spaces. Will not exceed the length of the message.
//
// Returns:
// - the message, shifted appropriately
func vigenereCipher(message string, key string) string {
	// Replace this with your code
	return ""
}

// Scytale cipher
//
// Encrypts a message by simulating a scytale cipher.
//
// A scytale is a cylinder around which you can wrap a long strip of
//
//	parchment that contained a string of apparent gibberish. The parchment,
//	when read using the scytale, would reveal a message due to every nth
//	letter now appearing next to each other, revealing a proper message.
//
// This encryption method is obsolete and should never be used to encrypt
//
//	data in production settings.
//
// You may read more about the method here:
//
//	https://en.wikipedia.org/wiki/Scytale
//
// You may follow this algorithm to implement a scytale-style cipher:
//  1. Take a message to be encoded and a "shift" number.
//     For this example, we will use "INFORMATION_AGE" as
//     the message and 3 as the shift.
//  2. Check if the length of the message is a multiple of
//     the shift. If it is not, add additional underscores
//     to the end of the message until it is.
//     In this example, "INFORMATION_AGE" is already a multiple of 3,
//     so we will leave it alone.
//  3. This is the tricky part. Construct the encoded message.
//     For each index i in the encoded message, use the character at the index
//     (i // shift) + (len(message) // shift) * (i % shift) of the raw message.
//     If this number doesn't make sense, you can play around with the cipher at
//     https://dencode.com/en/cipher/scytale to try to understand it.
//  4. Return the encoded message. In this case,
//     the output should be "IMNNA_FTAOIGROE".
//
// Example
// scytaleCipher("INFORMATION_AGE", 3) -> "IMNNA_FTAOIGROE"
// scytaleCipher("INFORMATION_AGE", 4) -> "IRIANMOGFANEOT__"
// scytaleCipher("ALGORITHMS_ARE_IMPORTANT", 8) -> "AOTSRIOALRH_EMRNGIMA_PTT"
//
// Params:
// - message, a string of uppercase English letters and underscores. Underscores represent spaces.
// - shift, a positive integer that does not exceed the length of the message.
//
// Returns:
// - the message, encoded appropriately.
func scytaleCipher(message string, shift int) string {
	// Replace this with your code
	return ""
}

// Scytale decipher
//
// Decrypts a message that was originally encrypted with the `scytaleCipher` function above.
//
// Example:
// scytaleDecipher('IMNNA_FTAOIGROE', 3) -> 'INFORMATION_AGE'
// scytaleDecipher('AOTSRIOALRH_EMRNGIMA_PTT', 8) -> 'ALGORITHMS_ARE_IMPORTANT'
// scytaleDecipher('IRIANMOGFANEOT__', 4) -> 'INFORMATION_AGE_'
//
// Params:
// - message, a string of uppercase English letters and underscores. Underscores represent spaces.
// - shift, a positive integer that does not exceed the length of the message.
//
// Returns:
// - The message, decoded appropriately.
func scytaleDecipher(message string, shift int) string {
	// Replace this with your code
	return ""
}
