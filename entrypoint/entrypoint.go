package entrypoint

import (
	"strings"
	"fmt"

	"github.com/AKovalevich/scrabbler/entrypoint/textclassifier"
)

// Base Entrypoint interface
type Entrypoint interface {
	// Start entrypoint
	Start()
	// Stop enptrypoint
	Stop()
	// Initialize entrypoint
	Init()
}

type EntrypointList []*Entrypoint

// String is the method to format the flag's value, part of the flag.Value interface.
// The String method's output will be used in diagnostics.
//func (e *EntrypointList) String() string {
//	return strings.Join(*e, ",")
//}

// Set is the method to set the flag value, part of the flag.Value interface.
// Set's argument is a string to be parsed to set the flag.
// It's a comma-separated list, so we split it.
func (e *EntrypointList) Set(value string) error {
	entrypoints := strings.Split(value, ",")
	if len(entrypoints) == 0 {
		return fmt.Errorf("bad EntryPointList format: %s", value)
	}
	for _, entrypointName := range entrypoints {
		// Try to create entrypoint
		switch entrypointName {
		case "textclassifier":
			textClassifierEntrypoint := textclassifier.New()
			*e = append(*e, textClassifierEntrypoint)
			break
		}
	}
	return nil
}

// Get return the EntryPoints map
func (e *EntrypointList) Get() interface{} {
	return EntrypointList(*e)
}

// SetValue sets the EntryPoints map with val
func (e *EntrypointList) SetValue(val interface{}) {
	*e = EntrypointList(val.(EntrypointList))
}

// Type is type of the struct
func (dep *EntrypointList) Type() string {
	return "defaultentrypoints"
}
