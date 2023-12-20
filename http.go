package glua

import (
	"bytes"
	"io"
	"net/http"
	"net/url"

	lua "github.com/yuin/gopher-lua"
)

func Request(L *lua.LState) int {
	method := L.CheckString(2)

	if method == "" {
		method = "GET"
	}

	switch method {
	case "GET":
		return Get(L)
	case "POST":
		return PostForm(L)
	case "POST_JSON":
		return PostJSON(L)
	}

	L.Push(lua.LNil)
	L.Push(lua.LString("not found method"))
	return 2
}

func Get(L *lua.LState) int {
	url := L.CheckString(1)

	resp, err := http.Get(url)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}
	if len(body) == 0 {
		L.Push(lua.LNil)
		L.Push(lua.LString("no response"))
		return 2
	}

	L.Push(lua.LString(body))
	return 1
}

func PostForm(L *lua.LState) int {
	apiURL := L.CheckString(1)
	data := L.CheckTable(2)

	urlValues := url.Values{}
	data.ForEach(func(key, value lua.LValue) {
		urlValues.Add(key.String(), value.String())
	})

	resp, err := http.PostForm(apiURL, urlValues)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}
	if len(body) == 0 {
		L.Push(lua.LNil)
		L.Push(lua.LString("no response"))
		return 2
	}

	L.Push(lua.LString(body))
	return 1
}

func PostJSON(L *lua.LState) int {
	apiURL := L.CheckString(1)
	jsonStr := L.CheckString(2)

	resp, err := http.Post(apiURL, "application/json", bytes.NewReader([]byte(jsonStr)))
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}
	if len(body) == 0 {
		L.Push(lua.LNil)
		L.Push(lua.LString("no response"))
		return 2
	}

	L.Push(lua.LString(body))
	return 1
}
