{{template "base/base.html" .}}
{{define "head"}}
<title>Title</title>
{{end}}
{{define "body"}}
<div class="container">
  <div class="row">
    <div class="col-sm-12">
      <h2>Register</h2>
      <div class="row">
        <div class="col-sm-4">
          {{if .flash.error}}
            <p>{{.flash.error}}</p>
          {{end}}
          {{if .flash.notice}}
            <p>{{.flash.notice}}</p>
          {{end}}
          {{range $key, $val := .Errors}}
          <p>{{$val}}</p>
          {{end}}
          <form action="/user/register" method="POST">
            <div class="form-group">
              <label for="firstname">First name</label>
              <input type="text" class="form-control" id="firstname" name="firstname" required placeholder="">
            </div>
            <div class="form-group">
              <label for="lastname">Last name</label>
              <input type="text" class="form-control" id="lastname" name="lastname" required placeholder="">
            </div>
            <div class="form-group">
              <label for="email">Email address</label>
              <input type="text" class="form-control" id="email" name="email" required placeholder="">
            </div>
            <div class="form-group">
              <label for="username">Username</label>
              <input type="text" class="form-control" id="username" name="username" required placeholder="">
            </div>
            <div class="form-group">
              <label for="password">Password</label>
              <input type="password" class="form-control" id="password" name="password" required placeholder="">
            </div>
            <div class="form-group">
              <label for="passwordConfirm">Password confirmation</label>
              <input type="password" class="form-control" id="passwordConfirm" name="password_confirm" required placeholder="">
            </div>
            <button type="submit" class="btn btn-default">Submit</button>
          </form>
        </div>
      </div>
    </div>    
  </div>
</div>
{{end}}