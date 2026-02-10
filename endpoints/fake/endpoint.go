package fake

import (
	"fmt"

	"github.com/SomtoJF/go-rod/browserfactory"
	"github.com/SomtoJF/go-rod/initializers/fs"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

type Endpoint struct {
	browserClient browserfactory.BrowserClient
	fs            *fs.TemporaryFileSystem
	browser       *rod.Browser
}

func NewFakeEndpoint(bc browserfactory.BrowserClient, fs *fs.TemporaryFileSystem) *Endpoint {
	b := bc.GetBrowser()
	return &Endpoint{
		browserClient: bc,
		fs:            fs,
		browser:       b,
	}
}

func (e *Endpoint) GetPageAccessibilityTree() ([]*proto.AccessibilityAXNode, error) {
	b := e.browser
	page := b.MustPage("https://www.wikipedia.org/").MustWindowFullscreen()

	accessibilityTree, err := e.browserClient.GetPageAccessibilityTree(page)
	if err != nil {
		return nil, err
	}

	fmt.Println(accessibilityTree)

	return accessibilityTree, nil
}

func (e *Endpoint) ScreenshotForLLM() (string, error) {
	b := e.browser
	page := b.MustPage("https://www.wikipedia.org/").MustWindowFullscreen()

	return e.browserClient.ScreenshotForLLM(page, "wikipedia.png")
}
