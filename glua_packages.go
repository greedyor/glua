package glua

import (
	lua "github.com/yuin/gopher-lua"
)

// Temporarily set up some Lua methods...

var Packages = map[string]map[string]lua.LGFunction{
	"json":   JsonFunctions,
	"http":   HttpFunctions,
	"result": SystemFunctions,
}

var JsonFunctions = map[string]lua.LGFunction{
	"decodeToTable": JsonToTable,
	"encodeToTable": TableToJson,
}

var HttpFunctions = map[string]lua.LGFunction{
	"request":  Request,
	"get":      Get,
	"post":     PostForm,
	"postJSON": PostJSON,
}

var SystemFunctions = map[string]lua.LGFunction{
	"errors":  SetError,
	"success": SetSuccess,
}

// import base packages
func ImportGluaPackges(glua *GluaVM, packageArgs []string) *GluaVM {
	for _, v := range packageArgs {
		if _, flag := Packages[v]; flag {
			glua.LState.SetGlobal(v, glua.LState.SetFuncs(glua.LState.NewTable(), Packages[v]))
		}
	}
	return glua
}
