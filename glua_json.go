package glua

import (
	"encoding/json"

	lua "github.com/yuin/gopher-lua"
)

func TableToJson(L *lua.LState) int {
	value := L.Get(1)

	goData := LuaValueToType(value)

	jsonBytes, err := json.Marshal(goData)
	if err != nil {
		L.RaiseError("failed to encode JSON: %v", err)
		return 0
	}

	L.Push(lua.LString(jsonBytes))
	return 1
}

func JsonToTable(L *lua.LState) int {
	jsonString := L.ToString(1)

	var goData interface{}
	err := json.Unmarshal([]byte(jsonString), &goData)
	if err != nil {
		L.RaiseError("failed to decode JSON: %!(NOVERB)v", err)
		return 0
	}

	luaValue := TypeToLuaValue(L, goData)

	L.Push(luaValue)
	return 1
}
