package utils

import "testing"

func TestUniqRand(t *testing.T) {

	vals := UniqRand(10, 100)
	for _, v := range vals {
		t.Logf("uniq values is :[%d]", v)
	}
}

func TestRand(t *testing.T) {
	vals := Rand(10, 100)
	for _, v := range vals {
		t.Logf("normal values is :[%d]", v)
	}
}

func BenchmarkRand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		vals := Rand(10, 100)
		for _, v := range vals {
			b.Logf("normal bench values is :[%d]", v)
		}
	}
}
