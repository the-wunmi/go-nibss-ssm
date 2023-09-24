package main

import (
	"fmt"
	SSM "go-nibss-ssm/ssm"
)

func main() {
	SSM.GenerateKeyPair("public.key", "private.key", "adewunmi", "1234567890")
	encryptedStr := SSM.EncryptMessage("./public.key", "hello world")
	fmt.Printf("encrypted value: %s", encryptedStr)
	decryptedStr := SSM.DecryptMessage("./private.key", "1234567890", encryptedStr)
	fmt.Printf("decrypted value: %s", decryptedStr)
}
