package crypto

import (
	"crypto/sha256"
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	// cipher key
	key := "111111"

	hash := sha256.New()
	hash.Write([]byte(key))
	sha256key := hash.Sum(nil)

	// plaintext
	pt := "This is a secretLoreLore utn mollis cursuque maximulaLorem ipsum , et venenatis neque maximus utn mollis cllis cura"

	c, _ := AESEncrypt(pt, sha256key)
	fmt.Println("sha256key", sha256key)
	// plaintext
	//fmt.Println(pt)

	// ciphertext
	fmt.Println("密文：", c)
	fmt.Println("-----------------------------------------------------------------------------------------------------")
	// decrypt
	decrypt, _ := AESDecrypt("f5bb872a08ef929e6744d117|18eb2d87f88705e1f2252e6a426dc96ca6f6b4f8bb1ee9b9c32fab63b105e51665266fe40720daa0bad1cf49a8bb64cdd9471fddfa1a63a6cd4c511c9fb8ec42dca02072a58d8908adfed346564208ed3c2fb956642aeb0df8bde8f923885c49ee5eb31eaa3ada304de25d377431c4da7437e2d20e07e1bd2969951f9675d411c48965", sha256key)
	fmt.Println("明文：", decrypt)
}
