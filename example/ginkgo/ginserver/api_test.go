package ginserver_test

import (

	. "bddgo"
	. "bddgo/example/ginkgo/ginserver/handlers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

var _ = Describe("Api", func() {

	BeforeSuite(func(){
		Router = gin.Default()
		Router.GET("success", Success)
		Router.GET("json", Json)
		Router.POST("login", Login)
		Router.POST("register", Register)
	})

	BeforeEach(func(){

	})

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
})
