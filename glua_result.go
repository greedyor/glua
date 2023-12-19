package glua

import (
	"errors"

	lua "github.com/yuin/gopher-lua"
)

func SetResult(L *lua.LState) int {
	L.SetGlobal("gluaResult", ValueToLua(L, LuaValueToString(L, L.CheckAny(1))))
	return 1
}

func (G *GluaVM) GetResultString() string {
	return G.GetGlobal("gluaResult").String()
}

func SetError(L *lua.LState) int {
	L.SetGlobal("gluaError", ValueToLua(L, LuaValueToString(L, L.CheckAny(1))))
	return 1
}

func (G *GluaVM) GetError() error {
	errString := G.GetGlobal("gluaError").String()
	if errString != "" {
		return errors.New(errString)
	}
	return nil
}
