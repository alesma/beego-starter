package routers

import (
	"sg/controllers"
	"github.com/astaxie/beego"
)

func init() {
  beego.Router("/", &controllers.MainController{})
  beego.Router("/home", &controllers.MainController{})
  beego.Router("/user/login/:back", &controllers.UserController{}, "get,post:Login")
  beego.Router("/user/logout", &controllers.UserController{}, "get:Logout")
  beego.Router("/user/register", &controllers.UserController{}, "get,post:Register")
  beego.Router("/user/profile", &controllers.UserController{}, "get,post:Profile")
  beego.Router("/user/verify/:uuid([0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12})", &controllers.UserController{}, "get:Verify")
}
