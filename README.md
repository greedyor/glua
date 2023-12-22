# Glua

Relying on the [gopher-lua](https://github.com/yuin/gopher-lua) plugin to quickly process lua.



## Usage

Import a package.
```golang
import (
    "github.com/greedyor/glua"
)
```

Run scripts 

```golang
data, err := glua.ExecToPath("./xxapi.lua")
if err != nil {
    panic("ExecToPath error:", err)
}
```
or
```golang
data, err := glua.ExecToCode("print("hello")")
if err != nil {
	panic("ExecToCode error:", err)
}
```


```xxapi.lua``` example

```lua
-- require func
local http = require("http")
local json = require("json")
local result = require("result")
-- request url
local response, err = http.request("https://github.com/manifest.json", "GET")
if err ~= nil then
    result.errors(err)
else
    -- json string to table type
    local tableData = json.decodeToTable(response)

    -- fieids set
    local fields = {
        name = tableData.name,
        shortName = tableData.short_name,
        icons = tableData.icons,
    }

    result.success(json.encodeToTable(fields))
end

```


## Installation

```
go get github.com/greedyor/glua
```

GopherLua supports >= Go1.9.

