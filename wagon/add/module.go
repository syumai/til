//go:generate statik -src=./module
package add

import (
	"fmt"
	"io"
	"log"

	"github.com/go-interpreter/wagon/exec"
	"github.com/go-interpreter/wagon/wasm"
	"github.com/rakyll/statik/fs"
	_ "github.com/syumai/til/wagon/add/statik"
)

var (
	vm     *exec.VM
	module *wasm.Module
)

func init() {
	f, err := getModuleFile()
	if err != nil {
		log.Fatalf("unexpected error occured on getting module file: %v", err)
	}
	defer f.Close()

	blankImporter := func(string) (*wasm.Module, error) { return nil, nil }
	module, err = wasm.ReadModule(f, blankImporter)
	if err != nil {
		log.Fatalf("unexpected error occured on reading module: %v", err)
	}

	vm, err = exec.NewVM(module)
	if err != nil {
		log.Fatalf("unexpected error occured on creating VM: %v", err)
	}
}

func getModuleFile() (io.ReadCloser, error) {
	staticFS, err := fs.New()
	if err != nil {
		return nil, fmt.Errorf("failed to create static FS of add module: %w", err)
	}
	f, err := staticFS.Open("/add.wasm")
	if err != nil {
		return nil, fmt.Errorf("failed to open static file of add module: %w", err)
	}
	return f, nil
}
