package main

import (
	"fmt"
	"time"

	"github.com/SomtoJF/go-rod/initializers/fs"
	"github.com/go-rod/rod"
)

func main() {
	fs := fs.NewTemporaryFilesystem()
	defer fs.Cleanup()

	page := rod.New().MustConnect().MustPage("https://www.wikipedia.org/")
	page.MustWaitStable().MustScreenshot(fs.GetBasePath() + "/b.png")

	fmt.Println(fs.GetBasePath() + "/b.png")
	time.Sleep(20 * time.Second)
}
