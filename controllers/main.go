package controllers

type MainController struct {
	baseController
}

func (c *MainController) Get() {
  c.setupView("index")
}

