package controllers

import (
  "fmt"
  "strings"
  "regexp"
  "time"
  "github.com/astaxie/beego"
  "github.com/astaxie/beego/validation"
  "github.com/astaxie/beego/orm"
  "golang.org/x/crypto/bcrypt"
  "gopkg.in/gomail.v1"
  "github.com/twinj/uuid"
  "sg/models"
)

type UserController struct {
  baseController
}

func (c *UserController) Login() {
  c.setupView("user/login")
  if c.Ctx.Input.Method() == "POST" {

    flash := beego.NewFlash()
    username := c.GetString("username")
    password := c.GetString("password")

    valid := validation.Validation{}
    valid.Required(username, "username")
    valid.Required(password, "password")
    if valid.HasErrors() {
        errormap := []string{}
        for _, err := range valid.Errors {
            errormap = append(errormap, "Validation failed on "+err.Key+": "+err.Message+"\n")
        }
        flash.Error("Invalid data!")
        flash.Store(&c.Controller)
        c.Data["Errors"] = errormap
        return
    }

    o := orm.NewOrm()
    o.Using("default")
    user := &models.AuthUser{
      Username: username,
    }
    err := o.Read(user, "Username")
    if err == orm.ErrNoRows {
      flash.Error("User not found!")
      flash.Store(&c.Controller)
      return
    }
    bcryptErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
    if (bcryptErr != nil) {
      flash.Error("Wrong password!")
      flash.Store(&c.Controller)
      return
    }
    m := make(map[string]interface{})
    m["email"] = user.Email
    m["username"] = user.Username
    m["timestamp"] = time.Now()
    c.SetSession("user", m)
    c.Redirect("/"+c.Ctx.Input.Param(":id"), 302) // go to previous page after login
  }
}


func (c *UserController) Logout() {
  c.setupView("index")
  c.DestroySession()
  c.Redirect("/", 302)
}

func (c *UserController) Register() {
  // get
  c.setupView("user/register")
  if c.Ctx.Input.Method() == "POST" {

    flash := beego.NewFlash()
    valid := validation.Validation{}

    firstname := c.GetString("firstname")
    lastname := c.GetString("lastname")
    username := c.GetString("username")
    email := c.GetString("email")
    password := c.GetString("password")
    passwordConfirm := c.GetString("password_confirm")

    // password validation
    valid.Required(passwordConfirm, "password_confirm")
    r, err := regexp.Compile(strings.Join([]string{"^",password,"$"}, ""));
    if err != nil {
      fmt.Printf("There is a problem with your regexp.")
      return
    }
    valid.Match(passwordConfirm, r, "password_confirm")

    config := uuid.StateSaverConfig{SaveReport: true, SaveSchedule: 30 * time.Minute}
    uuid.SetupFileSystemStateSaver(config)
    u1 := uuid.NewV4()

    user := &models.AuthUser{
      Firstname: firstname,
      Lastname: lastname,
      Username: username,
      Email: email,
      Password: password,
      Reg_key: u1.String(),
    }

    b, err := valid.Valid(user)
    if err != nil {
        fmt.Println(err)
    }
    if !b {
      errormap := []string{}
      for _, err := range valid.Errors {
        errormap = append(errormap, "Validation failed on "+err.Key+": "+err.Message+"\n")
      }
      flash.Error("Invalid data!")
      flash.Store(&c.Controller)
      c.Data["Errors"] = errormap
      fmt.Println(errormap)
      return
    }

    hash, hashErr := bcrypt.GenerateFromPassword([]byte(password), 1);    
    if hashErr != nil {
        c.Abort("401")
    }
    user.Password = string(hash)

    o := orm.NewOrm()
    o.Using("default")

    if created, id, err := o.ReadOrCreate(user, "Username"); err == nil {
      if created {
        fmt.Println("New user registered. Id:", id)
        flash.Notice("Welcome, "+username+"!")
        flash.Store(&c.Controller)

        link := "http://localhost:8080/user/verify/" + u1.String() // u1 is UUID
        host := "smtp.gmail.com"
        port := 587
        msg := gomail.NewMessage()
        msg.SetAddressHeader("From", "youremail@mail.com", "Start Go")
        msg.SetHeader("To", email)
        msg.SetHeader("Subject", "Account Verification for Start Go")
        msg.SetBody("text/html", "To verify your account, please click on the link: <a href=\""+link+"\">"+link+"</a>")
        m := gomail.NewMailer(host, "youremail@mail", "yourpassword", port)

        if errMail := m.Send(msg); errMail != nil {
            fmt.Println("Email was not sent!")
            fmt.Println(errMail)
        }

        return;
      } else {
        flash.Error("Invalid data!")
        flash.Store(&c.Controller)
        errormap := []string{}
        errormap = append(errormap, "User already exist")
        c.Data["Errors"] = errormap
        return
      }
    }
  }
}

func (c *UserController) Profile() {
  c.setupView("user/profile")
  // This page requires login
  if c.isLogged == false {
    c.Redirect("/user/login/home", 302)
    return
  }

  // get
  session := c.GetSession("user")
  username := session.(map[string]interface {})
  user := &models.AuthUser{
    Username: username["username"].(string),
  }
  o := orm.NewOrm()
  o.Using("default")
  errRead := o.Read(user, "Username")
  if errRead == orm.ErrNoRows {
    fmt.Println("User not found")
    c.Redirect("/user/login/home", 302)
    return
  }

  // post
  if c.Ctx.Input.Method() == "POST" {
    flash := beego.NewFlash()
    valid := validation.Validation{}

    firstname := c.GetString("firstname")
    lastname := c.GetString("lastname")
    email := c.GetString("email")

    user.Firstname = firstname
    user.Lastname = lastname
    user.Email = email

    b, err := valid.Valid(user)
    if err != nil {
      fmt.Println(err)
    }
    if !b {
      errormap := []string{}
      for _, err := range valid.Errors {
        errormap = append(errormap, "Validation failed on "+err.Key+": "+err.Message+"\n")
      }
      flash.Error("Invalid data!")
      flash.Store(&c.Controller)
      c.Data["Errors"] = errormap
      fmt.Println(errormap)
      return
    }

    numRows, errUpdate := o.Update(user, "Firstname", "Lastname", "Email")
    if errUpdate != nil {
      fmt.Println(errUpdate)
      return
    }
    if numRows > 0 {
      flash.Notice("Account updated!")
      flash.Store(&c.Controller)  
    }
  }

  c.Data["Firstname"] = user.Firstname
  c.Data["Lastname"] = user.Lastname
  c.Data["Username"] = user.Username
  c.Data["Email"] = user.Email
}

func (c *UserController) Verify() {
  c.setupView("user/verify")
  flash := beego.NewFlash()

  uuid := c.Ctx.Input.Param(":uuid")
  user := &models.AuthUser{
    Reg_key: uuid,
  }
  o := orm.NewOrm()
  o.Using("default")
  errRead := o.Read(user, "Reg_key")
  if errRead != nil {
    flash.Error("Invalid data!")
    flash.Store(&c.Controller)
    return
  }
  user.Reg_key = ""
  numRows, errUpdate := o.Update(user, "Reg_key")
  if errUpdate != nil {
    fmt.Println(errUpdate)
    return
  }
  if numRows > 0 {
    flash.Notice("Account verified!")
    flash.Store(&c.Controller)
  }
}