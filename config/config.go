package config

// ProcessingNode ProcessingNode
type ProcessingNode struct {
	NodeType int
	NodeName string
	NodeAddr string
}

// DefaultProcessingNode DefaultProcessingNode
func DefaultProcessingNode(n *ProcessingNode) {
	n = &ProcessingNode{
		NodeType: 1,
		NodeName: "",
		NodeAddr: "",
	}

	return
}
