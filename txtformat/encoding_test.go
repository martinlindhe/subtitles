package txtformat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLooksLikeLatin1(t *testing.T) {
	assert.Equal(t, true, looksLikeLatin1([]byte("hall\xe5")))
	assert.Equal(t, false, looksLikeLatin1([]byte("hallå")))
}
