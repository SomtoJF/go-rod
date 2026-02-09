package browserfactory

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

type BrowserClient interface {
	GetBrowser() *rod.Browser
	GetPageAccessibilityTree(*rod.Page) ([]*proto.AccessibilityAXNode, error)
	ScreenshotForLLM(*rod.Page, string) (string, error)
}
