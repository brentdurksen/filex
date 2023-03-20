package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"
)

// App struct
type App struct {
	ctx context.Context
}

type DirListing struct {
	Name    string
	Size    int64
	ModTime time.Time
	IsDir   bool
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) ListDir(path string, hidden bool) []DirListing {
	dir, err := os.ReadDir(path)
	if err != nil {
		return nil
	}
	var r []DirListing
	for _, listing := range dir {
		info, err := listing.Info()
		if err != nil {
			continue
		}
		name := info.Name()
		if name[0] == '.' && !hidden {
			continue
		}
		r = append(r, DirListing{info.Name(), info.Size(), info.ModTime(), info.IsDir()})

	}
	return r
}

func (a *App) OpenFile(path string) {
	fmt.Println(path)

	exec.Command("open", path).Start()
}
