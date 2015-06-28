{{template "base/base.html" .}}
{{define "head"}}
<title>Title</title>
{{end}}
{{define "body"}}
<div class="container">
  <div class="row">
    <div class="col-sm-12">
      <h2>Login</h2>
      <div class="row">
        <div class="col-sm-4">
          {{if .flash.error}}
            <p>{{.flash.error}}</p>
          {{end}}
          {{if .flash.notice}}
            <p>{{.flash.notice}}</p>
          {{end}}
          <form action="/user/login/home" method="POST">
            {{ .xsrfdata }}
            <div class="form-group">
              <label for="username">Username</label>
              <input type="text" class="form-control" id="username" name="username" placeholder="">
            </div>
            <div class="form-group">
              <label for="password">Password</label>
              <input type="password" class="form-control" id="password" name="password" placeholder="">
            </div>
            <button type="submit" class="btn btn-default">Submit</button>
          </form>
          {{range $key, $val := .Errors}}
          <p>{{$val}}</p>
          {{end}}
        </div>
      </div>
    </div>    
  </div>
</div>
{{end}}