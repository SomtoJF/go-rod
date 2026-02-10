package main

import (
	"fmt"
	"time"

	"github.com/SomtoJF/go-rod/common"
	"github.com/SomtoJF/go-rod/endpoints/fake"
	"github.com/SomtoJF/go-rod/initializers/fs"
)

func main() {
	fs := fs.NewTemporaryFilesystem()
	defer fs.Cleanup()

	dependencies := common.MakeDependencies(fs)

	endpoint := fake.NewFakeEndpoint(dependencies.BrowserFactory, fs)
	screenshotPath, taggedNodes, err := endpoint.ScreenshotForLLM()
	if err != nil {
		panic(err)
	}
	fmt.Println(screenshotPath)
	for _, node := range taggedNodes {
		name := ""
		if node.Node.Name != nil && !node.Node.Name.Value.Nil() {
			name = node.Node.Name.Value.String()
		}
		role := ""
		if node.Node.Role != nil && !node.Node.Role.Value.Nil() {
			role = node.Node.Role.Value.String()
		}
		value := ""
		if node.Node.Value != nil && !node.Node.Value.Value.Nil() {
			value = node.Node.Value.Value.String()
		}
		fmt.Printf("Name: %q, Role: %q, Value: %q, Index: %d\n", name, role, value, node.Index)
	}

	time.Sleep(1 * time.Hour)
}
