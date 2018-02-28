package checksum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCcittXmodem(t *testing.T) {
	assert.Equal(t, uint16(0x31D5),
		CcittXmodem([]byte{0x10, 0x07, 0x06, 0x00, 0x49, 0x01, 0x1B, 0x08, 0x1F, 0x08, 0x01, 0x00, 0x00, 0x10, 0x08}))
	assert.Equal(t, uint16(0x29A6),
		CcittXmodem([]byte{0x10, 0x07, 0x00, 0x20, 0x49, 0x00, 0x00, 0x04, 0x83, 0x00, 0x21, 0xFF, 0xFF, 0x73,
			0x00, 0x00, 0x00, 0x00, 0x01, 0x01, 0x41, 0x96, 0x64, 0x02, 0x02, 0x41, 0x96, 0x64, 0x03, 0x0F,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0xBD, 0x82, 0x65, 0x10, 0x08}))
}

func TestCcittKermit(t *testing.T) {
	assert.Equal(t, uint16(0x8923),
		CcittKermit([]byte{0x10, 0x07, 0x06, 0x00, 0x49, 0x01, 0x1B, 0x08, 0x1F, 0x08, 0x01, 0x00, 0x00, 0x10, 0x08}))
}
