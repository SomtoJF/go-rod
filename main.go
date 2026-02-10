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
		boundsStr := "nil"
		if node.Bounds != nil {
			boundsStr = fmt.Sprintf("{X: %.2f, Y: %.2f, Width: %.2f, Height: %.2f}", node.Bounds.X, node.Bounds.Y, node.Bounds.Width, node.Bounds.Height)
		}
		fmt.Printf("Name: %q, Role: %q, Value: %q, Index: %d, Bounds: %v\n", name, role, value, node.Index, boundsStr)

	}

	time.Sleep(1 * time.Hour)
}
