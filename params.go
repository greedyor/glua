package glua

import (
	"encoding/json"

	lua "github.com/yuin/gopher-lua"
)

func LuaValueToString(L *lua.LState, value lua.LValue) string {
	switch value.Type().String() {
	case "nil":
		return ""
	case "boolean", "number", "string":
		return value.String()
	case "table":
		goMap := make(map[string]interface{})
		table := L.NewTable()
		table.Append(value)
		table.ForEach(func(key lua.LValue, value lua.LValue) {
			goMap[key.String()] = value
		})
		jsonBytes, _ := json.Marshal(goMap)
		return string(jsonBytes)
	default:
		panic("Unsupported type:" + value.Type().String())
	}
}

func LuaValueToType(L lua.LValue) interface{} {
	switch value := L.(type) {
	case *lua.LTable:
		var isArray bool
		if value.Len() > 0 {
			isArray = true
		}
		if isArray {
			array := make([]interface{}, value.Len())
			for i := 1; i <= value.Len(); i++ {
				array[i-1] = LuaValueToType(value.RawGetInt(i))
			}
			return array
		} else {
			table := make(map[string]interface{})
			value.ForEach(func(k lua.LValue, v lua.LValue) {
				table[lua.LVAsString(k)] = LuaValueToType(v)
			})
			return table
		}
	case lua.LBool:
		return bool(value)
	case lua.LNumber:
		return float64(value)
	case lua.LString:
		return string(value)
	default:
		return nil
	}
}

func TypeToLuaValue(L *lua.LState, data interface{}) lua.LValue {
	switch value := data.(type) {
	case map[string]interface{}:
		table := L.NewTable()
		for key, val := range value {
			table.RawSetString(key, TypeToLuaValue(L, val))
		}
		return table
	case []interface{}:
		array := L.CreateTable(len(value), 0)
		for i, val := range value {
			array.RawSetInt(i+1, TypeToLuaValue(L, val))
		}
		return array
	case bool:
		return lua.LBool(value)
	case float64:
		return lua.LNumber(value)
	case string:
		return lua.LString(value)
	default:
		return lua.LNil
	}
}
