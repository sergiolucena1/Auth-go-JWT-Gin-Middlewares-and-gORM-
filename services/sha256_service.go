package services

import (
	"crypto/sha256"
	"fmt"
)

//encriptando o password
func SHA256Encoder(s string)string{
	str := sha256.Sum256([]byte(s))   // pacote do go parea encriptar

	return fmt.Sprintf("%x", str)
}
