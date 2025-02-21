package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZipcode(t *testing.T) {
	_, err1 := NewCEP("13180000")
	_, err2 := NewCEP("1318000")
	_, err3 := NewCEP("13180-000")
	_, err4 := NewCEP("")

	assert.NoError(t, err1)
	assert.Error(t, err2, "invalid zipcode")
	assert.Error(t, err3, "invalid zipcode")
	assert.Error(t, err4, "invalid zipcode")
}
