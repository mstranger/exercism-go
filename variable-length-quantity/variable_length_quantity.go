package variablelengthquantity

import "fmt"

// EncodeVarint encodes list of ints as series of bytes.
func EncodeVarint(input []uint32) []byte {
	output := make([]byte, 0)
	for _, x := range input {
		t := x >> 7
		chunk := append([]byte{}, byte(x&0x7f))
		for t > 0 {
			nb := byte((t & 0x7f) | 0x80)
			chunk = append([]byte{nb}, chunk...)
			t = t >> 7
		}
		output = append(output, chunk...)
	}

	return output
}

// DecodeVarint decodes bytes into a list of ints.
func DecodeVarint(input []byte) ([]uint32, error) {
	output := make([]uint32, 0)
	var n uint32
	var i int
	for _, b := range input {
		v := uint32(b & 0x7f)
		last := (b & 0x80) == 0
		n = (n << 7) | v
		if last == true {
			output = append(output, n)
			i = 0
			n = 0
		} else {
			i++
		}
	}

	if i > 0 {
		return output, fmt.Errorf("invalid input")
	}

	return output, nil
}
