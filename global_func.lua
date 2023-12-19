--
-- The following is the implementation method of Golang, Func for IDE use
--

-- set return
result = {
    ---@param error error
    errors = function(error)
    end
    ,
    ---@param data any
    success = function(data)
    end
}

-- json func，encode and decode
json = {
    ---@param jsonStr string
    ---@return string
    decodeToTable = function(jsonStr)
    end
    ,
    ---@param tableData table
    ---@return string
    encodeToTable = function(tableData)
    end
}

-- request method func
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
