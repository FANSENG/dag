package main

import (
	"context"
	"time"

	"fs1n/dag/base"
)

func main() {
	graph := base.NewTaskGraph()
	nodeId2 := graph.AddNode(func(ctx context.Context) error {
		time.Sleep(time.Second * 3)
		return nil
	}, time.Second)
	nodeId3 := graph.AddNode(func(ctx context.Context) error {
		println("3")
		return nil
	}, time.Second)
	nodeId4 := graph.AddNode(func(ctx context.Context) error {
		println("4")
		return nil
	}, time.Second)
	nodeId5 := graph.AddNode(func(ctx context.Context) error {
		println("5")
		return nil
	}, time.Second)

	graph.AddEdge(nodeId5, nodeId2)
	graph.AddEdge(nodeId2, nodeId3)
	graph.AddEdge(nodeId3, nodeId4)
	_ = graph.Run(context.Background())
}
