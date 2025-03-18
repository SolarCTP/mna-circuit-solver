package main

type NodeID uint64

type Node struct {
	ID NodeID
}

func NewNode(id NodeID) *Node {
	return &Node{
		ID: id,
	}
}
