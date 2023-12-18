--
-- 使用 golang 实现，该实现不调用为 ide 使用
--

-- 公共请求参数 params
params = nil

-- 公共GET请求参数 GETParams
local getParams = nil

-- 公共POST请求参数 POSTParams
POSTParams = nil

-- 公共返回字段
Result = ""

-- 公共错误字段
Error = nil

-- params 处理方法
params = {
    -- json字符串转 table 格式
    ---@param jsonStr string
    decodeToTable = function(jsonStr)
    end
    ,
    -- table 格式转 json字符串
    ---@param table table
    encodeToTable = function(table)
    end
}

-- json 处理方法
json = {
    -- json字符串转 table 格式
    ---@param jsonStr string
    ---@return string, error
    decodeToTable = function(jsonStr)
    end
    ,
    -- table 格式转 json字符串
    ---@param table table
    ---@return string, error
    encodeToTable = function(table)
    end
}

-- http 处理方法
http = {
    ---@param url string
    ---@param method string
    ---@param data any
    ---@return string, error
    request = function(url, method, data)
    end
    ,
    ---@param url string
    ---@return string, error
    Get = function(url)
    end
    ,
    ---@param url string
    ---@param data table
    ---@return string, error
    PostForm = function(url, data)
    end
    ,
    ---@param url string
    ---@param data string
    ---@return string, error
    PostJSON = function(url, data)
    end
}
