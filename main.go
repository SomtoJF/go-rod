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
		fmt.Printf("Description: %q\n", node.Description)
	}

	time.Sleep(1 * time.Hour)
}
