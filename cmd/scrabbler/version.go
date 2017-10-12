package scrabbler

import (
	"runtime"
	"fmt"

	"github.com/AKovalevich/scrabbler/version"
	"github.com/containous/flaeg"
)

const (
	VersionTemplate = `Version:      %s
Codename:     %s
Go version:   %s
Built:        %s
OS/Arch:      %s`
)

// newVersionCmd builds a new Version command
func newVersionCmd() *flaeg.Command {
	//version Command init
	return &flaeg.Command{
		Name:                  "version",
		Description:           `Print version`,
		Config:                struct{}{},
		DefaultPointersConfig: struct{}{},
		Run: func() error {
			fmt.Printf(VersionTemplate + "\n",
					version.Current(),
					version.Codename(),
					runtime.Version(),
					runtime.GOOS,
					runtime.GOARCH,
			)
			return nil
		},
	}
}
