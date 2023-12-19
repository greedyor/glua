package glua

import (
	"fmt"
	"sync/atomic"

	lua "github.com/yuin/gopher-lua"
)

type GluaVM struct {
	*lua.LState
	LuaCode string
	Path    string
	Loaded  atomic.Bool
}

func ExecToPath(path string, importPackages []string) (res string, err error) {
	gl := &GluaVM{
		Path:   path,
		LState: lua.NewState(),
	}

	ImportGluaPackges(gl, importPackages)

	if err = gl.Load(); err != nil {
		return
	}

	res = gl.GetSuccess()
	err = gl.GetError()

	return
}

func ExecToCode(code string, importPackages []string) (res string, err error) {
	gl := &GluaVM{
		LuaCode: code,
		LState:  lua.NewState(),
	}

	ImportGluaPackges(gl, importPackages)

	if err = gl.DoString(code); err != nil {
		return
	}

	res = gl.GetSuccess()
	err = gl.GetError()

	return
}

// create new a struct and load lua script
func Exec(path string, importPackages []string) (gl *GluaVM, err error) {
	gl = New(path).Imports(importPackages)

	if err = gl.Load(); err != nil {
		return
	}
	return
}

func New(path string) *GluaVM {
	return &GluaVM{
		Path:   path,
		LState: lua.NewState(),
	}
}

func (gl *GluaVM) Imports(importPackages []string) *GluaVM {
	return ImportGluaPackges(gl, importPackages)
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
