package walletrpc

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"
)

// NewPaymentID64 generates a 64 bit payment ID (hex encoded).
// With 64 bit IDs, there is a non-negligible chance of a collision
// if they are randomly generated. It is up to recipients generating
// them to sanity check for uniqueness.
//
// 1 million IDs at 64-bit (simplified): 1,000,000^2 / (2^64 * 2) = ~1/36,893,488 so
// there is a 50% chance a collision happens around 5.06 billion IDs generated.
func NewPaymentID64() string {
	buf := make([]byte, 8)
	rand.Read(buf)
	return hex.EncodeToString(buf)
}

// NewPaymentID256 generates a 256 bit payment ID (hex encoded).
func NewPaymentID256() string {
	buf := make([]byte, 32)
	rand.Read(buf)
	return hex.EncodeToString(buf)
}

// XMRToDecimal converts a raw atomic XMR balance to
// human readable string.
func XMRToDecimal(xmr uint64) string {
        if xmr == 0 {
                return "0"
        }
        str0 := fmt.Sprintf("%013d", xmr)
        l := len(str0)
        decimal := str0[:l-12]
        float := strings.TrimRight(str0[l-12:], "0")
        if len(float) == 0 {
                return decimal
        }
        return decimal + "." + float
}

func XMRToDecimal(xmr uint64) string {
	str0 := fmt.Sprintf("%013d", xmr)
	l := len(str0)
	return str0[:l-12] + "." + str0[l-12:]
}

// XMRToFloat64 converts raw atomic XMR to a float64
func XMRToFloat64(xmr uint64) float64 {
	return float64(xmr) / 1e12
}
