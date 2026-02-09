package common

import (
	"github.com/SomtoJF/go-rod/browserfactory"
	"github.com/SomtoJF/go-rod/initializers/fs"
)

type Dependencies struct {
	BrowserFactory *browserfactory.BrowserFactory
}

func NewDependencies(fs *fs.TemporaryFileSystem) *Dependencies {
	return &Dependencies{
		BrowserFactory: browserfactory.NewBrowserFactory(fs),
	}
}

func MakeDependencies(fs *fs.TemporaryFileSystem) *Dependencies {
	return NewDependencies(fs)
}
