package glua

import (
	"fmt"
	"sync/atomic"

	lua "github.com/yuin/gopher-lua"
)

type lvaVM struct {
	*lua.LState
	Path   string
	Loaded atomic.Bool
}

// create new a struct and load lua script
func Exec(path string, importPackages []string) (*lvaVM, error) {
	lv := New(path)

	ImportGluaPackges(lv.LState, importPackages)

	if err := lv.Load(); err != nil {
		return lv, err
	}
	return lv, nil
}

func New(path string) *lvaVM {
	lv := &lvaVM{
		Path:   path,
		LState: lua.NewState(),
	}
	return lv
}

func (gl *lvaVM) Load() error {
	if len(gl.Path) == 0 {
		return fmt.Errorf("plugin file empty")
	}
	if err := gl.DoFile(gl.Path); err != nil {
		return err
	}
	gl.Loaded.Store(true)
	return nil
}
