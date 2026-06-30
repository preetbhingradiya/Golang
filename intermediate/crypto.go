package intermediate

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func Crypt() {

	password := "password123"

	//Hexadecimal encoding of the password ( 256 )
	// hash256 := sha256.Sum256([]byte(password))
	// fmt.Println("SHA-256 Hash:", hash256)
	// hex256 := hex.EncodeToString(hash256[:])
	// fmt.Println("Hexadecimal Representation:", hex256)

	// //Hexadecimal encoding of the password ( 512 )
	// hash512 := sha512.Sum512([]byte(password))
	// fmt.Println("SHA-512 Hash:", hash512)
	// hex512 := hex.EncodeToString(hash512[:])
	// fmt.Println("Hexadecimal Representation:", hex512)

	//Hash with salt
	salt := generateSalt()
	saltedPassword := password + salt
	hashSalted := sha256.Sum256([]byte(saltedPassword))
	hexSalted := hex.EncodeToString(hashSalted[:])
	fmt.Println("Hexadecimal Representation of Salted Hash:", hexSalted)

	verify := verifyPassword(password, salt, hexSalted)
	fmt.Println("Password verification result:", verify)
}

// Generate random salt
func generateSalt() string {
	salt := make([]byte, 16)
	fmt.Println("SALT : ", salt)
	rand.Read(salt)
	fmt.Println("SALT : ", salt)
	return hex.EncodeToString(salt)
}

func verifyPassword(inputPassword, storedSalt, storedHash string) bool {
	// Recreate hash using stored salt
	salted := inputPassword + storedSalt
	hash := sha256.Sum256([]byte(salted))
	newHash := hex.EncodeToString(hash[:])

	// Compare
	return newHash == storedHash
}
