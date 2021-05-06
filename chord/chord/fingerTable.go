package chord

import (
	"math"
)

const (
	float64Two = float64(2)
)

func (node *Node) GenerateNewFingerTable() {
	fingerTable := make(FingerTable, int(HashIdSize))
	for i := float64(0); i < HashIdSize; i++ {
		fingerTable[int(i)] = FingerEntry{uint32(math.Mod((node.HashIdFloat64 + math.Pow(float64Two, i)), math.Pow(float64Two, HashIdSize))), node}
	}
	node.FingerTable = fingerTable
}
