package serialization

import (
	"io"
	"io/ioutil"
	"encoding/json"
	"log"
	"os"
	"bytes"
	"../trees"
)

func SerializeTree(tree trees.Tree, writer io.Writer) error {
	if err := json.NewEncoder(writer).Encode(tree); err != nil {
		return err
	}
	return nil
}

func DeserializeTree(reader io.Reader) (trees.Tree, error) {
	var tree trees.Tree
	if err := json.NewDecoder(reader).Decode(&tree); err != nil {
		return tree, err
	}
	return tree, nil
}

func SaveTree(tree trees.Tree, filename string) {
	file, err := os.Create(filename)
	if err != nil {
        log.Fatalf("Error creating file: %s", err)
    }

    defer file.Close()

	if err = SerializeTree(tree, file); err != nil {
		log.Fatalf("Error writing file: %s", err)
	}
}

func ReadTree(filename string) trees.Tree {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
        log.Fatalf("Error reading file: %s", err)
    }
    tree, err := DeserializeTree(bytes.NewReader(file))
    if err != nil {
    	log.Fatalf("Error decoding json: %s", err)
    }
    return tree
}