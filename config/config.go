package config

// ProcessingNode ProcessingNode
type ProcessingNode struct {
	NodeType int
	NodeName string
	NodeAddr string
}

// DefaultProcessingNode DefaultProcessingNode
func DefaultProcessingNode() ProcessingNode {
	n := ProcessingNode{
		NodeType: 1,
		NodeName: "",
		NodeAddr: "",
	}

	return n
}
