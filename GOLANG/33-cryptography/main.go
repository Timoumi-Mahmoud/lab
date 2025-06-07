package main

import (
	"crypto/md5"
	"fmt"
	"unicode"
)

type Cipher interface {
	Encryption(string) string
	Decryption(string) string
}

type cipher []int

func (c cipher) cipherAlgorithm(letters string, shift func(int, int) int) string {
	shiftedText := ""
	for _, letter := range letters {
		if !unicode.IsLetter(letter) {
			continue
		}
		shiftDist := c[len(shiftedText)%len(c)]

		fmt.Println("shiftDist:", shiftDist)
		s := shift(int(unicode.ToLower(letter)), shiftDist)
		switch {
		case s < 'a':
			s += 'z' - 'a' + 1
		case 'z' < s:
			s -= 'z' - 'a' + 1
		}
		shiftedText += string(s)
	}
	return shiftedText
}
func (c *cipher) Encryption(plainText string) string {
	return c.cipherAlgorithm(plainText, func(a, b int) int { return a + b })
}

func (c *cipher) Decryption(cipherText string) string {
	return c.cipherAlgorithm(cipherText, func(a, b int) int { return a - b })
}
func NewCaesar(key int) Cipher {
	return NewShift(key)
}
func NewShift(shift int) Cipher {
	if shift < -25 || 25 < shift || shift == 0 {
		return nil
	}
	c := cipher([]int{shift})
	return &c
}

func main() {
	fmt.Println("vim-go")
	c := NewCaesar(1)
	fmt.Println("Encrypt Key(01) abcd =>", c.Encryption("abcd"))
	fmt.Println("Decrypt Key(01) bcde =>", c.Decryption("bcde"))
	fmt.Println()
	fmt.Println("\n----------------Small Message----------------\n")
	message := []byte("Today web engineering has modern apps adhere to what is known as a single-page app (SPA) model.")

	fmt.Printf("Md5: %x\n\n", md5.Sum(message))
	fmt.Printf("Md5: %x\n\n", md5.Sum(message))

}

/*
Crypto: Secret, Graphy: writing
the process of altering messages in a way that their meaning is hidden from an enemy or opponent who might seize them.

[sender] encrypt -------Network----------> [receiver] decrypt

Plaintext
Ciphertext
Encryption
Decryption
Ciphers: the Encryption and Decryption algorithms are together knonw as ciphers.
Key

*/
