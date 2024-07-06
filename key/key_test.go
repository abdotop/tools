package key

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeyGen(t *testing.T) {
	for i := 0; i < 10; i++ {
		password := Generate(12, true, false)
		fmt.Println(password)
		assert.NotEmpty(t, password)
	}
}
