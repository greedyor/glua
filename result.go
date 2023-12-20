package glua

import (
	"errors"

	lua "github.com/yuin/gopher-lua"
)

func SetSuccess(L *lua.LState) int {
	L.SetGlobal("gluaResult", TypeToLuaValue(L, LuaValueToString(L, L.CheckAny(1))))
	return 1
}

func (G *GluaVM) GetSuccess() string {
	return G.GetGlobal("gluaResult").String()
}

func SetError(L *lua.LState) int {
	L.SetGlobal("gluaError", TypeToLuaValue(L, LuaValueToString(L, L.CheckAny(1))))
	return 1
}

func (G *GluaVM) GetError() error {
	errString := LuaValueToString(G.LState, G.GetGlobal("gluaError"))
	if errString != "" {
		return errors.New(errString)
	}
	return nil
}
