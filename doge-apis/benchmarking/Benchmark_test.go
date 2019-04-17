package benchmarking

import (
	"doge-apis/apis"
	"testing"
)

func BenchmarkHandleFormular(b *testing.B) {
	for n := 0; n < b.N; n++ {
		// Insert you function here
		apis.HandleFormular()
	}
}
