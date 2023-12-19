package glua

import (
	"fmt"
	"sync/atomic"

	lua "github.com/yuin/gopher-lua"
)

type GluaVM struct {
	*lua.LState
	Path     string
	Loaded   atomic.Bool
	luaError error
}

// create new a struct and load lua script
func Exec(path string, importPackages []string) (gl *GluaVM, err error) {
	gl = New(path)

	ImportGluaPackges(gl.LState, importPackages)

	if err = gl.Load(); err != nil {
		return
	}
	return
}

func New(path string) *GluaVM {
	lv := &GluaVM{
		Path:   path,
		LState: lua.NewState(),
	}
	return lv
}

func (gl *GluaVM) Load() error {
	if len(gl.Path) == 0 {
		return fmt.Errorf("plugin file empty")
	}
	if err := gl.DoFile(gl.Path); err != nil {
		return err
	}
	gl.Loaded.Store(true)
	return nil
}
