package lib

import (
	"fmt"

	lua "github.com/yuin/gopher-lua"
)

func ValueToLua(L *lua.LState, value interface{}) lua.LValue {
	switch v := value.(type) {
	case int:
		return lua.LNumber(v)
	case float64:
		return lua.LNumber(v)
	case bool:
		return lua.LBool(v)
	case string:
		return lua.LString(v)
	case []interface{}: // in case is array
		table := L.NewTable()
		for i, item := range v {
			L.SetTable(table, lua.LNumber(i+1), ValueToLua(L, item))
		}
		return table
	case map[string]interface{}: // in case is map
		table := L.NewTable()
		for key, item := range v {
			L.SetTable(table, lua.LString(key), ValueToLua(L, item))
		}
		return table
	default:
		panic(fmt.Sprintf("Unsupported type: %T", v))
	}
}
