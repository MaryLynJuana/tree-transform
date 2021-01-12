package main

import (
	"log"
	"bytes"
	"../../serialization"
	"../../trees"
)

var baseUrl = "http://localhost:8000"

func main() {
	tree := serialization.ReadTree("client-tree.json")
	trees.WalkTree(&tree, trees.TransformTree)
	var body bytes.Buffer
	serialization.SerializeTree(tree, &body)
    if err := SendRequest(baseUrl + "/tree", body); err != nil {
    	log.Fatalf("Error sending request: %s", err)
    }
}