package controllers

import (
  "github.com/astaxie/beego"
)

type baseController struct {
  beego.Controller
  isLogged bool
}

func (c *baseController) setupView(view string) {
  c.TplNames = view + ".tpl"
  session := c.GetSession("user")
  c.isLogged = false
  if session != nil {
    c.Data["InSession"] = 1 // for login bar
    m := session.(map[string]interface{})
    c.Data["username"] = m["username"] // first name
    c.isLogged = true
  }
}