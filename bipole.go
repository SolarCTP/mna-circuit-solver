package main

type Bipole interface {
	Component
	SetNodes(nodeA Node, nodeB Node)
}
