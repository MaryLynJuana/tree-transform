package handlers

import (
	"log"
	"net/http"
	"../serialization"
)

func TreeHandler(rw http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
        rw.Write([]byte("Not a post request"))
        return
    } 
	body := req.Body
	tree, err := serialization.DeserializeTree(body)
	if err != nil {
    	log.Fatalf("Error decoding request body: %s", err)
    }
	serialization.SaveTree(tree, "server-tree.json")
}