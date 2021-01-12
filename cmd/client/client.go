package main

import (
	"log"
	"net/http"
	"bytes"
)

func SendRequest(path string, body bytes.Buffer) error {
	client := http.Client{}
	req, err := http.NewRequest( 
        "POST", path, &body, 
    )
    if err != nil {
    	log.Fatalf("Error creating request: %s", err)
    }
    _, err = client.Do(req)
    return err
}