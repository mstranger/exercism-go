package railfence

import "testing"

func testCases(op func(string, int) string, cases []testCase, t *testing.T) {
	for _, tc := range cases {
		if actual := op(tc.message, tc.rails); actual != tc.expected {
			t.Fatalf("FAIL: %s\nExpected: %q\nActual: %q", tc.description, tc.expected, actual)
		}
		t.Logf("PASS: %s", tc.description)
	}
}

func TestEncode(t *testing.T) { testCases(Encode, encodeTests, t) }
func TestDecode(t *testing.T) { testCases(Decode, decodeTests, t) }

// self write
func BenchmarkEncode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, test := range encodeTests {
			Encode(test.message, test.rails)
		}
	}
}

// self write
func BenchmarkDecode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, test := range encodeTests {
			Decode(test.expected, test.rails)
		}
	}
}
