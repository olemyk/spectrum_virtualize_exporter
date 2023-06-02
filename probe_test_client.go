package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/google/go-jsonnet"
)

type fakeClient struct {
	data map[string][]byte
}

func (c *fakeClient) prepare(path string, jfile string) {
	vm := jsonnet.MakeVM()
	b, err := os.ReadFile(jfile)
	if err != nil {
		log.Fatalf("Failed to read jsonnet %q: %v", jfile, err)
	}
	output, err := vm.EvaluateAnonymousSnippet(jfile, string(b))
	if err != nil {
		log.Fatalf("Failed to evaluate jsonnet %q: %v", jfile, err)
	}
	c.data[path] = []byte(output)
}

func (c *fakeClient) Get(path string, query string, obj interface{}) error {
	d, ok := c.data[path]
	if !ok {
		log.Fatalf("Tried to get unprepared URL %q", path)
	}
	return json.Unmarshal(d, obj)
}

func newFakeClient() *fakeClient {
	return &fakeClient{data: map[string][]byte{}}
}
