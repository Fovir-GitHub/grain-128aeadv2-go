package grain

import "github.com/Fovir-GitHub/grain-128aeadv2-go/internal/utils"

// Init initializes LFSR and NFSR for encryption and decryption.
//
// It also returns states of LFSR and NFSR after loading and initialziation in hex format string.
func (g *Grain128AEADV2) Init() (loadedLFSR, loadedNFSR, initLFSR, initNFSR string) {
	// Load key and nonce into LFSR and NFSR.
	loadedLFSR, loadedNFSR = g.loadLFSRNFSR()

	// Initialize LFSR and NFSR.
	for t := range 512 {
		lfsrFeedback := g.lfsrFeedback()
		nfsrFeedback := g.nfsrFeedback()
		yt := g.preOutput()
		s0 := g.LFSR[0]
		g.shiftLFSR(t, lfsrFeedback, yt)
		g.shiftNFSR(t, s0, nfsrFeedback, yt)
	}

	initLFSR, _ = utils.Bits2Hex(g.LFSR)
	initNFSR, _ = utils.Bits2Hex(g.NFSR)

	return loadedLFSR, loadedNFSR, initLFSR, initNFSR
}

// loadLFSRNFSR loads LFSR and NFSR using key and nonce.
//
// LFSR = Nonce + 31 ones + a zero
//
// NFSR = Key
func (g *Grain128AEADV2) loadLFSRNFSR() (loadedLFSR, loadedNFSR string) {
	g.LFSR = make([]int, 0, FeedbackRegisterBitLength)
	g.LFSR = append(g.LFSR, g.nonce...)
	for range 31 {
		g.LFSR = append(g.LFSR, 1)
	}
	g.LFSR = append(g.LFSR, 0)

	g.NFSR = make([]int, FeedbackRegisterBitLength)
	copy(g.NFSR, g.key)

	loadedLFSR, _ = utils.Bits2Hex(g.LFSR)
	loadedNFSR, _ = utils.Bits2Hex(g.NFSR)
	return loadedLFSR, loadedNFSR
}

// lfsrFeedback calculates the LFSR feedback value L(S_t).
func (g *Grain128AEADV2) lfsrFeedback() int {
	s := g.LFSR
	feedback := s[0] ^ s[7] ^ s[38] ^ s[70] ^ s[81] ^ s[96]
	return feedback
}

// nfsrFeedback calculates the NFSR feedback value F(B_t).
func (g *Grain128AEADV2) nfsrFeedback() int {
	b := g.NFSR
	feedback := (b[0] ^
		b[26] ^
		b[56] ^
		b[91] ^
		b[96] ^
		(b[3] & b[67]) ^
		(b[11] & b[13]) ^
		(b[17] & b[18]) ^
		(b[27] & b[59]) ^
		(b[40] & b[48]) ^
		(b[61] & b[65]) ^
		(b[68] & b[84]) ^
		(b[22] & b[24] & b[25]) ^
		(b[70] & b[78] & b[82]) ^
		(b[88] & b[92] & b[93] & b[95]))
	return feedback
}

// h is the functiton h(x).
func (g *Grain128AEADV2) h() int {
	s := g.LFSR
	b := g.NFSR
	x := []int{
		b[12],
		s[8],
		s[13],
		s[20],
		b[95],
		s[42],
		s[60],
		s[79],
		s[94],
	}
	hx := ((x[0] & x[1]) ^
		(x[2] & x[3]) ^
		(x[4] & x[5]) ^
		(x[6] & x[7]) ^
		(x[0] & x[4] & x[8]))
	return hx
}

// preOutput calculates the value of y_t.
func (g *Grain128AEADV2) preOutput() int {
	hx := g.h()
	s93 := g.LFSR[93]
	A := []int{2, 15, 36, 45, 64, 73, 89}
	b := g.NFSR
	sumB := 0

	for _, j := range A {
		sumB ^= b[j]
	}

	yt := hx ^ s93 ^ sumB
	return yt
}

// shiftLFSR updates LFSR based on `t`.
func (g *Grain128AEADV2) shiftLFSR(t, lfsrFeedback, yt int) {
	k := g.key
	s127 := 0

	// Update `s127` based on `t <= 319`, `t <= 383`, or `t <= 511` (otherwise).
	if t <= 319 {
		s127 = lfsrFeedback ^ yt
	} else if t <= 383 {
		s127 = lfsrFeedback ^ yt ^ k[t-256]
	} else {
		s127 = lfsrFeedback
	}

	g.LFSR = append(g.LFSR[1:], s127)
}

// shiftNFSR updates NFSR based on `t`.
func (g *Grain128AEADV2) shiftNFSR(t, s0, nfsrFeedback, yt int) {
	k := g.key
	b127 := 0

	// Update b127 based on `t <= 319`, `t <= 383`, or `t <= 511` (otherwise).
	if t <= 319 {
		b127 = s0 ^ nfsrFeedback ^ yt
	} else if t <= 383 {
		b127 = s0 ^ nfsrFeedback ^ yt ^ k[t-320]
	} else {
		b127 = s0 ^ nfsrFeedback
	}

	g.NFSR = append(g.NFSR[1:], b127)
}
