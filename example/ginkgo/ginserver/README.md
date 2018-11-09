GinServer
===========

在这里我们使用[ginkgo](http://onsi.github.io/ginkgo/)这个库来学习在gin server工程里面实践bdd开发。

## Install

```shell
    go get github.com/onsi/ginkgo/ginkgo
    go get github.com/onsi/gomega
```

## Start

### bootstrap

create a project in $GOPATH/src/ginserver

```
ginkgo bootstrap
```

will create a file ginserver_suite_test.go

```go
package book_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBook(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Book Suite")
}

```

### Prepare some gin handler for api test

在handlers里面我们建立了几个简单handlers: 

1. success.go 

```go 
package handlers 

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context) {
	c.String(http.StatusOK, "success")
}
```

1. json.go

```go
package handlers

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int `json:"code"`
	Message string `json:"message"`
}

func Json(c *gin.Context) {
	c.JSON(200, Response{
		Code: 0,
		Message: "success",
	})
}
```

1. login.go

```go
package handlers

import (
	"github.com/gin-gonic/gin"
)


func Login(c *gin.Context) {
	c.JSON(200, Response{
		Code: 1,
		Message: "wrong account",
	})
}
```

1. register.go

```go
package handlers

import (
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	c.JSON(200, Response{
		Code: 0,
		Message: "success registered",
	})
}
```

下面我们将编写相关的test去测试上面的接口


### 增加测试用例

run command 

```shell
ginkgo generate api
```

将会创建一个api_test.go

```go
package book_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "bddgo/ginkgo/book"
)

var _ = Describe("Book", func() {
})
```

此时没有任何的spec去执行，需要我们自己去编写Describe, Context, 和It

首先我们需要将路由在测试之前注册起来，这里我们在BeforeSuite的钩子里面执行注册。增加如下的代码:

```go

var Router *gin.Engine

BeforeSuite(func(){
    Router = gin.Default()
    Router.GET("success", Show)
    Router.GET("json", Json)
    Router.POST("login", Login)
    Router.POST("register", Register)
})
```

注意，我们首先定义了个全局的Router变量，方便在后续的测试脚本里面引用到路由信息。

我们根据请求的方式，分组了两个Describer, 第一个是get,是关于success.go 和 json.go 这两个handler的测试：

```go
    Describe("request get api", func(){
		Context("request /success", func(){
			It("should be success", func(){
				res := Get("/success", Router)
				Expect(string(res)).To(Equal("success"))
			})
		})
		Context("request /json", func(){
			It("should be json", func(){
				res := Get("/json", Router)
				Expect(string(res)).To(Equal(`{"code":0,"message":"success"}`))
			})
		})
	})
```

/success是直接返回了字符串"success", 所以断言是 `Expect(string(res)).To(Equal("success"))`
而/json是返回的json，所以断言是`Expect(string(res)).To(Equal(`{"code":0,"message":"success"}`))`

接下来是关于两种post请求的Describe：

```go
Describe("request post api", func(){
		Context("request /login", func(){
			It("with wrong account", func(){
				param := map[string]string{
					"UserName": "ney",
					"Password": "12345",
				}
				res := PostForm("/login", param, Router)
				Expect(string(res)).To(Equal(`{"code":1,"message":"wrong account"}`))
			})
		})

		Context("request /register", func(){
			It("with json post", func(){
				param := map[string]interface{}{
					"UserName": "ney",
					"Password": "12345",
				}

				res := PostJson("/register", param, Router)
				Expect(string(res)).To(Equal(`{"code":0,"message":"success registered"}`))

			})
		})
	})
```

/login的请求格式是x-www-form-urlencoded, 而/register的请求数据的格式是application/json.

最后我们执行test:

    ginkgo

将会输出如下的结果:

    Running Suite: Ginserver Suite
    ==============================
    Random Seed: 1541757855
    Will run 4 of 4 specs

    [GIN] 2018/11/09 - 18:04:21 | 200 |      72.084µs |       192.0.2.1 | GET      /success
    •[GIN] 2018/11/09 - 18:04:21 | 200 |      94.664µs |       192.0.2.1 | GET      /json
    •[GIN] 2018/11/09 - 18:04:21 | 200 |       5.814µs |       192.0.2.1 | POST     /login?UserName=ney&Password=12345
    •[GIN] 2018/11/09 - 18:04:21 | 200 |       4.639µs |       192.0.2.1 | POST     /register
    •
    Ran 4 of 4 Specs in 0.003 seconds
    SUCCESS! -- 4 Passed | 0 Failed | 0 Pending | 0 Skipped
    PASS
    [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

    [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
    - using env:   export GIN_MODE=release
    - using code:  gin.SetMode(gin.ReleaseMode)

    [GIN-debug] GET    /success                  --> bddgo/example/ginkgo/ginserver/handlers.Success (3 handlers)
    [GIN-debug] GET    /json                     --> bddgo/example/ginkgo/ginserver/handlers.Json (3 handlers)
    [GIN-debug] POST   /login                    --> bddgo/example/ginkgo/ginserver/handlers.Login (3 handlers)
    [GIN-debug] POST   /register                 --> bddgo/example/ginkgo/ginserver/handlers.Register (3 handlers)

    Ginkgo ran 1 suite in 5.474448034s
    Test Suite Passed

全部通过了！