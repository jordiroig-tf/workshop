package cypher

const key = "BACKENDARCHITECTURE"

func Decrypt(message string) string {
	return total(message, decrypt)
}

func Encrypt(message string) string {
	return total(message, encrypt)
}

func total(message string, fcrypt crypt) string {
	result := ""
	j := 0
	for i := 0; i < len(message); i++ {
		keychar := key[j%len(key)]
		char := message[i]

		encryptedChar := char

		if isValidChar(char) {
			encryptedChar = fcrypt(char, keychar)
			j++
		}

		result = result + string(encryptedChar)
	}

	return result
}

type crypt func(uint8, uint8) uint8

func encrypt(char, keychar uint8) uint8 {
	return 'A' + ((char + keychar + 26) % 26)
}

func decrypt(char, keychar uint8) uint8 {
	return 'A' + ((char - keychar + 26) % 26)
}

func isValidChar(char uint8) bool {
	return char >= 'A' && char <= 'Z'
}
