package salt

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
)

// 生成salt
// 输入是加密长度 返回值是加密后的盐.
func GenerateRandomSalt(saltSize int) string {
	var salt = make([]byte, saltSize)

	_, err := rand.Read(salt)

	if err != nil {
		panic(err)
	}

	return string(salt)
}

// 使用salt 和 SHA-512
// 参数第一个是明文 第二个是盐
// 返回值是加密后的密码.
func HashPassword(password string, salt string) string {
	// Convert password string to byte slice
	var passwordBytes = []byte(password)

	// Create sha-512 hasher
	var sha512Hasher = sha512.New()

	// Append salt to password
	passwordBytes = append(passwordBytes, salt...)

	// Write password bytes to the hasher
	sha512Hasher.Write(passwordBytes)

	// Get the SHA-512 hashed password
	var hashedPasswordBytes = sha512Hasher.Sum(nil)

	// Convert the hashed password to a hex string
	var hashedPasswordHex = hex.EncodeToString(hashedPasswordBytes)

	return hashedPasswordHex
}

// 验证： 第一个参数是加密后密码 第二个参数是明文 第三个参数是盐 返回值布尔值.
func PasswordsMatch(hashedPassword, currPassword string,
	salt string) bool {
	var currPasswordHash = HashPassword(currPassword, salt)

	return hashedPassword == currPasswordHash
}
