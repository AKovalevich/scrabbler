package main // import "github.com/hashicorp/vault"

import (
	"github.com/AKovalevich/scrabbler/cli"
	"os"
)

func main() {
	os.Exit(cli.Run())
}
