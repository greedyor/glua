# glua

Relying on the gopher-lua plugin to quickly process lua.



## usage

Import a package.
```
import (
    "github.com/greedyor/glua"
)
```

Run scripts 

```
data, err := glua.ExecToPath("./xxapi.lua")
if err != nil {
	panic("ExecToPath error:", err)
}
```
```
data, err := glua.ExecToCode("print("hello")")
if err != nil {
	panic("ExecToCode error:", err)
}
```


```xxapi.lua``` example

```
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
go get github.com/
```

GopherLua supports >= Go1.9.

