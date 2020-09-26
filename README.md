# servicehub-providers
useful providers for servicehub

# Providers
* httpserver
* pprof
* health
* i18n

## httpserver
http-server ServiceProvider 提供 http 服务器，供其他 Provider 使用。

支持拦截器、多种请求处理器，*请求参数* 和 *返回参数* 不分顺序

### 例子
[例子详情](./examples/httpserver/main.go)

### 请求参数
支持请求类型:
* http.ResponseWriter
* *http.Request
* httpserver.Context
* struct or struct Pointer
* map[string]interface{} or map[string]interface{} Pointer
* []byte
* string
* slice

### 数据校验
支持结构体数据校验，参考 **github.com/go-playground/validator**

### 返回参数
支持返回类型：
* int 表示 Response Status
* io.ReadCloser
* io.Reader
* interface{}
* error

### 拦截器
```go
func(handler func(ctx httpserver.Context) error) func(ctx httpserver.Context) error {
    return handler // 返回新的处理器
}
```
