package glua

import (
	"encoding/json"

	lua "github.com/yuin/gopher-lua"
)

func TableToJson(L *lua.LState) int {
	table := L.CheckTable(1)

	goMap := make(map[string]interface{})
	table.ForEach(func(key lua.LValue, value lua.LValue) {
		goMap[key.String()] = value
	})

	jsonBytes, err := json.Marshal(goMap)
	if err != nil {
		L.Push(lua.LString(""))
		return 2
	}

	L.Push(lua.LString(jsonBytes))
	return 1
}

func JsonToTable(L *lua.LState) int {
	jsonData := L.CheckString(1)

	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		panic(err)
	}

	table := L.NewTable()
	for key, value := range data {
		L.SetTable(table, lua.LString(key), ValueToLua(L, value))
	}

	L.Push(table)
	return 1
}
