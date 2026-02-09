package browserfactory

import (
	"fmt"

	"github.com/SomtoJF/go-rod/initializers/fs"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

type BrowserFactory struct {
	browser *rod.Browser
	fs      *fs.TemporaryFileSystem
}

func NewBrowserFactory(fs *fs.TemporaryFileSystem) *BrowserFactory {
	return &BrowserFactory{
		browser: rod.New().MustConnect().NoDefaultDevice(),
		fs:      fs,
	}
}

func (b *BrowserFactory) GetBrowser() *rod.Browser {
	return b.browser
}

func (b *BrowserFactory) GetPageAccessibilityTree(page *rod.Page) ([]*proto.AccessibilityAXNode, error) {
	res, err := proto.AccessibilityGetFullAXTree{}.Call(page)
	if err != nil {
		return nil, err
	}
	return res.Nodes, nil
}

func (b *BrowserFactory) ScreenshotForLLM(page *rod.Page, fileName string) (string, error) {
	screenshotPath := b.fs.ConcatenatePath(fileName)

	err := rod.Try(func() {
		page.MustWaitStable()
		// Get the accessibility tree for the page
		accessibilityTree, _ := b.GetPageAccessibilityTree(page)

		// Draw transparent grid lines over the page
		page.MustEval(`() => {
			const canvas = document.createElement('canvas');
			canvas.id = 'agent-grid';
			canvas.style = 'position:fixed; top:0; left:0; pointer-events:none; z-index:9999;';
			canvas.width = window.innerWidth;
			canvas.height = window.innerHeight;
			const ctx = canvas.getContext('2d');
			ctx.strokeStyle = 'rgba(255, 0, 0, 0.2)'; // Faint red lines
			// Draw horizontal/vertical lines every 100px
			for(let i=0; i<canvas.width; i+=100) { ctx.strokeRect(i, 0, 0, canvas.height); }
			for(let i=0; i<canvas.height; i+=100) { ctx.strokeRect(0, i, canvas.width, 0); }
			document.body.appendChild(canvas);
		}`)

		tagAccessibilityNodes(page, accessibilityTree)

		page.MustScreenshot(screenshotPath)

	})

	if err != nil {
		return "", err
	}

	return screenshotPath, nil
}

func tagAccessibilityNodes(page *rod.Page, accessibilityTree []*proto.AccessibilityAXNode) {
	for _, node := range accessibilityTree {
		page.MustEval(fmt.Sprintf(`() => {
			const node = document.getElementById('%s');
			node.style.border = '2px solid red';
		}`, node.NodeID))
	}
}
