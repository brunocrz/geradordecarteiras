package geradordecarteiras

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func Gerador() {
	for i := 1; i < 3; i++ {

		// Gerar uma chave privada aleatória usando o método GenerateKey a partir do pacote crypto.
		privateKey, err := crypto.GenerateKey()
		if err != nil {
			log.Fatal(err)
		}

		// Converte a chave privada para bytes usando o método FromECDSA do pacote crypto/ecdsa.
		privateKeyBytes := crypto.FromECDSA(privateKey)

		// Converte a chave privada em uma sequência hexadecimal
		//usando o método Encode do pacote hexutil e imprime a nova Chave Privada.
		fmt.Println("Chave Privada", i, ":", hexutil.Encode(privateKeyBytes))

		// Gera uma chave pública a partir da Chave Privada usando o método Public do pacote crypto.
		publicKey := privateKey.Public()

		// Verifica o tipo de chave Pública e em seguida converte em uma sequência hexadeciomal
		// Usando o mesmo processo da chave Privada e imprime a Chave Pública.
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			log.Fatal("Não é possível afirmar o tipo da publicKey, não é do tipo *ecdsa.PublicKey")
		}

		publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
		fmt.Println("Chave Publica", i, ":", hexutil.Encode(publicKeyBytes))

		// Gera um endereço de carteira usando o método PubkeyToAddress do pacote crypto
		// que aceita a chave pública ECDSA, retorna um endereço e armazena na variavel address.
		address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
		// Imprimindo o endereço junto com uma mensagem "Endereço:"
		fmt.Println("Endereço", i, ":", address)
		fmt.Println(" ")
	}
}
