package grain

import "github.com/Fovir-GitHub/grain-128aeadv2-go/internal/utils"

// Encrypt encrypts a given plaintext, and returns corresponding ciphertext.
func (g *Grain128AEADV2) Encrypt(plaintext []byte) []byte {
	// Initialize again.
	g.Init()

	// Transform byte to bits.
	m := utils.Byte2Bits(plaintext)
	L := len(m)
	ciphertext := make([]int, L)

	// Calculate the ciphertext.
	i := 0
	for t := 512; t < 512+2*L; t++ {
		yt := g.preOutput()
		if t%2 == 0 {
			ciphertext[i] = yt ^ m[i]
			i++
		}

		lfsrFeedback := g.lfsrFeedback()
		nfsrFeedback := g.nfsrFeedback()
		s0 := g.LFSR[0]
		g.shiftLFSR(t, lfsrFeedback, yt)
		g.shiftNFSR(t, s0, nfsrFeedback, yt)
	}

	// Transform btis to byte (ignore errors).
	c, _ := utils.Bits2Byte(ciphertext)
	return c
}

// Decrypt decrypts a given ciphertext, and returns corresponding plaintext.
func (g *Grain128AEADV2) Decrypt(ciphertext []byte) []byte {
	// Initailize again.
	g.Init()

	// Transform byte to bits.
	c := utils.Byte2Bits(ciphertext)
	L := len(c)
	plaintext := make([]int, L)

	// Calculate the plaintext.
	i := 0
	for t := 512; t < 512+2*L; t++ {
		yt := g.preOutput()
		if t%2 == 0 {
			plaintext[i] = yt ^ c[i]
			i++
		}
		lfsrFeedback := g.lfsrFeedback()
		nfsrFeedback := g.nfsrFeedback()
		s0 := g.LFSR[0]
		g.shiftLFSR(t, lfsrFeedback, yt)
		g.shiftNFSR(t, s0, nfsrFeedback, yt)
	}

	// Transform btis to byte (ignore errors).
	p, _ := utils.Bits2Byte(plaintext)
	return p
}
