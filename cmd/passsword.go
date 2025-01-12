package cmd

import (
	"crypto/rand"
	"fmt"
	"math/big"

	mathRand "math/rand"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a random password",
	Long: `Generate a random password with the specified length.
	
	By default, the password will contain only letters.
	
	for example: password-gen generate --length 20 --digits --special --number 5
	`,

	Run: generatePassword,
}

func generatePassword(cmd *cobra.Command, args []string) {
	length, _ := cmd.Flags().GetInt("length")
	isDigits, _ := cmd.Flags().GetBool("digits")
	isSpecial, _ := cmd.Flags().GetBool("special")
	isNumber, _ := cmd.Flags().GetInt("number")

	lowerCase := "abcdefghijklmnopqrstuvwxyz"
	upperCase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits := "0123456789"
	special := "!@#$%^&*()-_=+[]{}|;:,.<>?/~`"

	minRequiredLength := 2 // Minimum for at least one uppercase and one lowercase
	if isDigits {
		minRequiredLength++
	}
	if isSpecial {
		minRequiredLength++
	}
	if length < minRequiredLength {
		fmt.Printf("Password length must be at least %d to include all specified character types.\n", minRequiredLength)
		return
	}

	var charSet string
	charSet += lowerCase + upperCase

	if isDigits {
		charSet += digits
	}
	if isSpecial {
		charSet += special
	}

	for i := 0; i < isNumber; i++ {
		var password []byte
		password = append(password, lowerCase[randInt(len(lowerCase))])
		password = append(password, upperCase[randInt(len(upperCase))])
		if isDigits {
			password = append(password, digits[randInt(len(digits))])
		}
		if isSpecial {
			password = append(password, special[randInt(len(special))])
		}
		// Fill the rest of the password randomly
		for j := len(password); j < length; j++ {
			password = append(password, charSet[randInt(len(charSet))])
		}

		// Shuffle the password
		shuffle(password)

		// Print the generated password
		fmt.Println(string(password))
	}
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().IntP("length", "l", 8, "Length of the password")
	generateCmd.Flags().BoolP("digits", "d", false, "Include digits in the password")
	generateCmd.Flags().BoolP("special", "s", false, "Include special characters in the password")
	generateCmd.Flags().IntP("number", "n", 1, "Number of passwords to generate")
}

// Helper function to generate a random integer
func randInt(max int) int {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		fmt.Println("Error generating random number:", err)
		return 0
	}
	return int(n.Int64())
}

// Helper function to shuffle a byte slice
func shuffle(slice []byte) {
	mathRand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}
