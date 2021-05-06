package chord

import "hash/crc32"

var HashIdSize = float64(32)

func GenerateHashId(id string) uint32 {
	return crc32.ChecksumIEEE([]byte(id))
}
