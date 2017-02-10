package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"

	"github.com/WWGLondon/graduation/trailfinder"
)

const mapPath = "map/release_party_map.json"

func main() {
	trailReader := fileReader(mapPath)
	trailFinder := &trailfinder.Trailfinder{}
	err := trailFinder.Interpret(trailReader)
	if err != nil {
		log.Fatal(err)
	}

}

func fileReader(filepath string) (io.Reader, error) {
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	buf.Write(b)
	return &buf, nil
}
