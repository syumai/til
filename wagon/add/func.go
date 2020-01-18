package add

import "fmt"

func Add(a, b int) (int, error) {
	addFunc, ok := module.Export.Entries["add"]
	if !ok {
		return 0, fmt.Errorf("failed to get add func from wasm module")
	}
	o, err := vm.ExecCode(int64(addFunc.Index), uint64(a), uint64(b))
	if err != nil {
		return 0, fmt.Errorf("unexpected error occured on execute add func: %w", err)
	}
	result, ok := o.(uint32)
	if !ok {
		return 0, fmt.Errorf("return type must be uint32")
	}
	return int(result), nil
}
