package glua

import (
	lua "github.com/yuin/gopher-lua"
)

func SetError(L *lua.LState) int {
	L.SetGlobal("gluaError", ValueToLua(L, LuaValueToString(L, L.CheckAny(1))))
	return 1
}

func (gl *GluaVM) GetError() error {
	return gl.luaError
}
