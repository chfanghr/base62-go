package base62_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const str = `base62`

func TestEncoding_EncodeBytes(t *testing.T) {
	assert.Equal(t, "UiP9AV6Y", StdEncoding.EncodeBytes([]byte(str)))
	assert.Equal(t, "uIp9av6y", InvertedEncoding.EncodeBytes([]byte(str)))
}

func TestEncoding_DecodeToBytes(t *testing.T) {
	assert.Equal(t, []byte(str), StdEncoding.DecodeToBytes("UiP9AV6Y"))
	assert.Equal(t, []byte(str), InvertedEncoding.DecodeToBytes("uIp9av6y"))
}
