# Beego starter kit

You can find the framework documentation on the official site [Beego framework](http://beego.me/)  
Frontend build process uses [Gulp](http://gulpjs.com/), it compiles sass to css and allows you to write es6 js with [babeljs](https://babeljs.io/), find the gulpfile [here](https://github.com/alesma/beego-starter/blob/master/static/gulpfile.js)

# Get it running
You will need a mysql installation (I am using [mariadb](https://mariadb.org/)).  
The mysql connector is configured in main.go, which probably needs some changes according to your db setup.  

Install go dependencies by running:
```
go get github.com/astaxie/beego github.com/beego/bee github.com/go-sql-driver/mysql github.com/twinj/uuid golang.org/x/crypto/bcrypt gopkg.in/gomail.v1 github.com/twinj/uuid
```
Run it with:
```
bee run
```

Gulp runs on nodejs and has its dependencies, install them with npm by running in /static
```
npm install
```
Run gulp with
```
gulp
```

