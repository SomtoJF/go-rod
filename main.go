package main

import (
	"fmt"

	"github.com/SomtoJF/go-rod/initializers/fs"
	"github.com/go-rod/rod"
)

func main() {
	fs := fs.NewTemporaryFilesystem()
	defer fs.Cleanup()

	browser := rod.New().MustConnect().NoDefaultDevice()
	page := browser.MustPage("https://www.wikipedia.org/").MustWindowFullscreen()

	page.MustElement("#searchInput").MustInput("earth")
	page.MustElement("#search-form > fieldset > button").MustClick()

	screenshotPath := fs.ConcatenatePath("b.png")

	page.MustWaitStable().MustScreenshot(screenshotPath)
	fmt.Println(screenshotPath)
}
