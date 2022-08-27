package testhelper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewErrWriter(t *testing.T) {
	var expected = fmt.Errorf("FAILED")
	_, err := NewErrWriter(expected).Write([]byte("something"))
	assert.EqualError(t, err, expected.Error())
}
