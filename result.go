package glua

import (
	lua "github.com/yuin/gopher-lua"
)

func SetResult(L *lua.LState) int {
	L.SetGlobal("gluaResult", ValueToLua(L, LuaValueToString(L, L.CheckAny(1))))
	return 1
}

func (G *GluaVM) GetResultString() string {
	return G.GetGlobal("gluaResult").String()
}
