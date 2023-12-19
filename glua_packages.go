package glua

import (
	lua "github.com/yuin/gopher-lua"
)

// Temporarily set up some Lua methods...

var Packages = map[string]map[string]lua.LGFunction{
	"json": JsonFunctions,
	"http": HttpFunctions,
	"sys":  SystemFunctions,
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
	"errors": SetError,
	"result": SetResult,
}

// import base packages
func ImportGluaPackges(luaVM *lua.LState, packageArgs []string) {
	for _, v := range packageArgs {
		if _, flag := Packages[v]; flag {
			luaVM.SetGlobal(v, luaVM.SetFuncs(luaVM.NewTable(), Packages[v]))
		}
	}
}
