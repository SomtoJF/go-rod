package main

import (
	"github.com/SomtoJF/go-rod/common"
	"github.com/SomtoJF/go-rod/endpoints/fake"
	"github.com/SomtoJF/go-rod/initializers/fs"
)

func main() {
	fs := fs.NewTemporaryFilesystem()
	defer fs.Cleanup()

	dependencies := common.MakeDependencies(fs)

	endpoint := fake.NewFakeEndpoint(dependencies.BrowserFactory, fs)
	endpoint.GetPageAccessibilityTree()

}
