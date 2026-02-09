package fake

import (
	"fmt"

	"github.com/SomtoJF/go-rod/browserfactory"
	"github.com/SomtoJF/go-rod/initializers/fs"
	"github.com/go-rod/rod/lib/proto"
)

type Endpoint struct {
	browserClient browserfactory.BrowserClient
	fs            *fs.TemporaryFileSystem
}

func NewFakeEndpoint(bc browserfactory.BrowserClient, fs *fs.TemporaryFileSystem) *Endpoint {
	return &Endpoint{
		browserClient: bc,
		fs:            fs,
	}
}

func (e *Endpoint) GetPageAccessibilityTree() []*proto.AccessibilityAXNode {
	b := e.browserClient.GetBrowser()
	page := b.MustPage("https://www.wikipedia.org/").MustWindowFullscreen()

	accessibilityTree, err := e.browserClient.GetPageAccessibilityTree(page)
	if err != nil {
		return nil, err
	}

	fmt.Println(accessibilityTree)

	return accessibilityTree
}
