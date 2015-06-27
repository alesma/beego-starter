{{template "base/base.html" .}}
{{define "head"}}
<title>Title</title>
{{end}}
{{define "body"}}
<div class="container">
  <div class="row">
    <div class="col-sm-12">
      <h2>Profile</h2>
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
          <form action="/user/profile" method="POST">
            <div class="form-group">
              <label for="firstname">First name</label>
              <input type="text" class="form-control" id="firstname" name="firstname" required placeholder="" value="{{.Firstname}}">
            </div>
            <div class="form-group">
              <label for="lastname">Last name</label>
              <input type="text" class="form-control" id="lastname" name="lastname" required placeholder="" value="{{.Lastname}}">
            </div>
            <div class="form-group">
              <label for="email">Email address</label>
              <input type="text" class="form-control" id="email" name="email" required placeholder="" value="{{.Email}}">
            </div>
            <div class="form-group">
              <label for="username">Username</label>
              <input type="text" class="form-control" id="username" name="username" placeholder="" disabled value="{{.Username}}">
            </div>
            <button type="submit" class="btn btn-default">Submit</button>
          </form>
        </div>
      </div>
    </div>    
  </div>
</div>
{{end}}