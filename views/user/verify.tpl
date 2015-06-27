{{template "base/base.html" .}}
{{define "head"}}
<title>Title</title>
{{end}}
{{define "body"}}
<div class="container">
  <div class="row">
    <div class="col-sm-12">
      <h2>Verify</h2>
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
        </div>
      </div>
    </div>    
  </div>
</div>
{{end}}