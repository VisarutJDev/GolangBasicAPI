package unittest

import (
	"doge-apis/apis"
	"testing"
)

func TestHelloThere(t *testing.T) {
	test := apis.HelloThere()
	if test == "" {
		t.Error("Expected Hello, world!")
	}
}
