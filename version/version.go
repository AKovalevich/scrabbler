package version

import (
	"log"

	"github.com/hashicorp/go-version"
)

const (
	VERSION = "0.0.1"
	CODENAME = "scrab"
)

func Current() (string) {
	v1, err := version.NewVersion(VERSION)
	if err != nil {
		log.Print("")
	}
	return v1.String()
}

func Codename() string {
	return CODENAME
}
