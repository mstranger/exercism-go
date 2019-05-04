package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

// PrivateKey generates a private key.
func PrivateKey(p *big.Int) *big.Int {
	// need random from the range [2, p)
	r, _ := rand.Int(rand.Reader, new(big.Int).Sub(p, big.NewInt(2)))
	r.Add(r, big.NewInt(2))
	return r
}

// PublicKey generates a public key.
func PublicKey(private, p *big.Int, g int64) *big.Int {
	// A = g**a mod p
	return new(big.Int).Exp(big.NewInt(g), private, p)
}

// NewPair returns a public and a private keys.
func NewPair(p *big.Int, g int64) (private, public *big.Int) {
	private = PrivateKey(p)
	public = PublicKey(private, p, g)
	return
}

// SecretKey calculates a secret key.
func SecretKey(private1, public2, p *big.Int) *big.Int {
	// s = B**a mod p
	return new(big.Int).Exp(public2, private1, p)
}
