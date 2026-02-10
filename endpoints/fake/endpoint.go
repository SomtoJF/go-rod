package fake

import (
	"github.com/SomtoJF/go-rod/browserfactory"
	"github.com/SomtoJF/go-rod/initializers/fs"
	"github.com/go-rod/rod"
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

func (e *Endpoint) ScreenshotForLLM() (string, []*browserfactory.TaggedAccessibilityNode, error) {
	b := e.browser
	page := b.MustPage("https://ng.indeed.com/jobs?q=software%20engineer&l=&from=searchOnHP%2Cwhatautocomplete%2CwhatautocompleteSourceStandard").MustWindowFullscreen()

	return e.browserClient.ScreenshotForLLM(page, "wikipedia.png")
}
