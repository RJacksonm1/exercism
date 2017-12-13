package raindrops

import (
	"reflect"
	"testing"
)

func TestFactors(t *testing.T) {
	for _, test := range factorTests {
		if actual := Factors(test.input); !reflect.DeepEqual(actual, test.expected) && test.input != 1 {
			t.Errorf("Factors(%d) = %d, expected %d.", test.input, actual, test.expected)
		}
	}
}

func TestConvert(t *testing.T) {
	for _, test := range tests {
		if actual := Convert(test.input); actual != test.expected {
			t.Errorf("Convert(%d) = %q, expected %q.",
				test.input, actual, test.expected)
		}
	}
}

func BenchmarkConvert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			Convert(test.input)
		}
	}
}
