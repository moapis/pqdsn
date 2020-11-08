package pqdsn

import "testing"

func BenchmarkBuilder_addString(t *testing.B) {
	const (
		k = "key"
		s = "foo's bar is a long string that's nice"
	)

	b := &builder{escape: true}

	for i := 0; i < t.N; i++ {
		b.addString(k, s)
	}
}
