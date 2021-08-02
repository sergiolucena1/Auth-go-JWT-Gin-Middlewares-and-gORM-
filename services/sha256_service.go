package services

import (
	"crypto/sha256"
	"fmt"
)

//encriptando o password( PASSANDO NOSSA SENHA PARA SHA 256)
func SHA256Encoder(s string)string{
	str := sha256.Sum256([]byte(s))   // pacote do go parea encriptar

	return fmt.Sprintf("%x", str)
}
