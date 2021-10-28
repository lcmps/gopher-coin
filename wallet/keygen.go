package wallet

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func GenerateKeyHash() (pKey, sKey string) {
	privateKey, _ := crypto.GenerateKey()

	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Printf("Private key: %s\n", hexutil.Encode(privateKeyBytes)[2:])

	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Printf("Public key:\t %s\n", hexutil.Encode(publicKeyBytes)[4:])

	return hexutil.Encode(publicKeyBytes)[4:], hexutil.Encode(privateKeyBytes)[2:]
}
