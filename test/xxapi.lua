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

    -- local a = json.encodeToTable(tableData.icons[0])
    -- print(a)

    -- 返回信息
    result.success(json.encodeToTable(fields))
end
