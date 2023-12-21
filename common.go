package glua

import (
	"fmt"
	"sync/atomic"

	lua "github.com/yuin/gopher-lua"
)

type GluaVM struct {
	*lua.LState
	GluaContent
	Loaded atomic.Bool
}

type GluaContent struct {
	LuaCode string
	LuaPath string
}

func ExecToPath(path string, importPackages ...string) (res string, err error) {
	gl := New().SetPuth(path).Imports(importPackages)
	defer gl.Close()

	if err = gl.LoadFile(); err != nil {
		return
	}

	res = gl.GetSuccess()
	err = gl.GetError()

	return
}

func ExecToCode(code string, importPackages ...string) (res string, err error) {
	gl := New().SetCode(code).Imports(importPackages)
	defer gl.Close()

	if err = gl.DoString(code); err != nil {
		return
	}

	res = gl.GetSuccess()
	err = gl.GetError()

	return
}

// create new a struct and load lua script
func Exec(path string, importPackages ...string) (gl *GluaVM, err error) {
	gl = New().SetPuth(path).Imports(importPackages)
	defer gl.Close()

	if err = gl.LoadFile(); err != nil {
		return
	}

	return
}

func New() *GluaVM {
	gl := &GluaVM{
		LState: lua.NewState(),
	}

	InitPreloadModules(gl)

	return gl
}

func (gl *GluaVM) SetPuth(path string) *GluaVM {
	gl.LuaPath = path
	return gl
}

func (gl *GluaVM) SetCode(code string) *GluaVM {
	gl.LuaCode = code
	return gl
}

func (gl *GluaVM) Imports(importPackages []string) *GluaVM {
	return ImportGluaPackges(gl, importPackages)
}

func (gl *GluaVM) LoadFile() error {
	if len(gl.LuaPath) == 0 {
		return fmt.Errorf("plugin file empty")
	}
	if err := gl.DoFile(gl.LuaPath); err != nil {
		return err
	}
	gl.Loaded.Store(true)
	return nil
}
