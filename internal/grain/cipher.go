package grain

import "github.com/Fovir-GitHub/grain-128aeadv2-go/internal/utils"

func (g *Grain128AEADV2) Encrypt(plaintext []byte) []byte {
	g.Init()
	m := utils.Byte2Bits(plaintext)
	L := len(m)
	ciphertext := make([]int, L)
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

	c, _ := utils.Bits2Byte(ciphertext)
	return c
}

func (g *Grain128AEADV2) Decrypt(ciphertext []byte) []byte {
	g.Init()
	c := utils.Byte2Bits(ciphertext)
	L := len(c)
	plaintext := make([]int, L)
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

	p, _ := utils.Bits2Byte(plaintext)
	return p
}
