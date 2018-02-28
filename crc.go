package checksum

func Simple(data []byte) uint64 {
	sum := uint64(0)
	for _, b := range data {
		sum += uint64(b)
	}
	checksum := 256 - sum%256
	if checksum == 256 {
		return 0
	} else {
		return checksum
	}
}

func Xor(data []byte) byte {
	checksum := byte(0)
	for _, b := range data {
		checksum ^= b
	}
	return checksum
}

func polyBigEndian(data []byte, poly, initial uint16) uint16 {
	crc := initial
	for _, b := range data {
		crc ^= uint16(b) << 8
		for n := 0; n < 8; n++ {
			if crc&0x8000 == 0x8000 {
				crc = (crc << 1) ^ poly
			} else {
				crc <<= 1
			}
		}
	}
	return crc
}

func polyLittleEndian(data []byte, poly, initial uint16) uint16 {
	crc := initial
	for _, b := range data {
		c := b
		for n := 0; n < 8; n++ {
			if (crc^uint16(c))&0x0001 == 0x0001 {
				crc = (crc >> 1) ^ poly
			} else {
				crc = crc >> 1
			}
			c = c >> 1
		}
	}
	return crc
}

func CcittXmodem(data []byte) uint16 {
	return polyBigEndian(data, 0x1021, 0x0000)
}

func CcittKermit(data []byte) uint16 {
	crc := polyLittleEndian(data, 0x8408, 0x0000)
	return (crc&0x00FF)<<8 | crc>>8
}

func Crc16BigEndian(data []byte) uint16 {
	return polyBigEndian(data, 0x8005, 0xFFFF)
}

func Crc16LittleEndian(data []byte) uint16 {
	return polyLittleEndian(data, 0xA001, 0xFFFF)
}
