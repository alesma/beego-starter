package models
import (
    "github.com/astaxie/beego/orm"
    "time"
    "strings"
    "github.com/astaxie/beego/validation"
)
type AuthUser struct {
    Id          int
    Firstname   string `valid:"Required;MaxSize(25)"`
    Lastname    string `valid:"Required;MaxSize(25)"`
    Username    string `orm:"unique";valid:"Required;MaxSize(25);MinSize(4)"`
    Email       string `orm:"unique";valid:"Email;MaxSize(100);MinSize(4)"`
    Password    string `valid:"Required"`
    Reg_key     string
    Reg_date    time.Time `orm:"auto_now_add;type(datetime)"`
}

func (u *AuthUser) Valid(v *validation.Validation) {
    if strings.Index(u.Username, "admin") != -1 {
        v.SetError("Username", "Can't contain 'admin' in Name")
    }
}

func init() {
    orm.RegisterModel(new(AuthUser))
}