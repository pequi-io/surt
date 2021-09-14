package clamav

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	clamdAddress  = "127.0.0.1:3310"
	clamdtNetwork = "tcp"
)

func TestNewClamavClient(t *testing.T) {
	_, err := New(clamdtNetwork, clamdAddress)
	assert.Nil(t, err)
}

func TestNewInvalidClamavClient(t *testing.T) {
	_, err := New(clamdAddress, clamdAddress)
	assert.NotNil(t, err)
}
