package chord

func InitializeNetwork(id string) Node {
	node := NewNode(id)
	node.ServeGrpc()
	return *node
}
