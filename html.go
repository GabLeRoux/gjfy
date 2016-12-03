package main

const (
	htmlMaster = `
	{{define "master"}}
	<!DOCTYPE html>
	<html>
	<head>
		<title>GJFY{{block "title" .}}{{end}}</title>
		<link rel="shortcut icon" type="image/x-icon" href="favicon.ico">
		<link rel="stylesheet" type="text/css" href="custom.css">
	</head>
	<body>
	<div id="contentcontainer">
	<div id="content">
	{{block "content" .}}{{end}}
	</div>
	{{block "footer" .}}
		<div id="footer">
			<a href="https://github.com/sstark/gjfy">gjfy</a>
		</div>
	{{end}}
	</div>
	</body>
	</html>
	{{end}}
	`
	htmlView = `
	{{define "title"}} - View Secret{{end}}
	{{define "content"}}
	<h2 id="mainheading">The following text is just for you:</h2>
	<div id="main">
	<div>
		The link you invoked contains a secret (a password for example)
		somebody wants to share with you. It will be valid only for a short
		time and you may not be able to invoke it again. Please make sure
		you memorise the secret or write it down in an appropriate way.
	</div>
	<div>The secret contained in this link is as follows:</div>
	<pre id="secret">{{.Secret}}</pre>
	</div>
	{{end}}
	`
	htmlViewInfo = `
	{{define "title"}} - View Info{{end}}
	{{define "content"}}
	<h2 id="mainheading">Metadata for {{.Id}}</h2>
	<div id="main">
	<table id="info">
	<tr>
		<th>Id</th>
		<th>MaxClicks</th>
		<th>Clicks</th>
		<th>DateAdded</th>
		<th>AuthToken</th>
	</tr>
	<tr>
		<td><a href="{{.Url}}">{{.Id}}</a></td>
		<td>{{.MaxClicks}}</td>
		<td>{{.Clicks}}</td>
		<td>{{.DateAdded}}</td>
		<td>{{.AuthToken}}</td>
	</tr>
	</table>
	</div>
	{{end}}
	`
	htmlViewErr = `
	{{define "title"}} - Error{{end}}
	{{define "content"}}
	<h2 id="errorheading">Not available</h2>
	<div id="main">
	<p id="errormessage">This ID is not valid anymore. Please request another one from the person who sent you this link.</p>
	</div>
	{{end}}
	`
	cssFileName = "custom.css"
)

var (
	favicon = []byte{0x0, 0x0, 0x1, 0x0, 0x1, 0x0, 0x10, 0x10, 0x10, 0x0, 0x1,
		0x0, 0x4, 0x0, 0x28, 0x1, 0x0, 0x0, 0x16, 0x0, 0x0, 0x0, 0x28, 0x0, 0x0,
		0x0, 0x10, 0x0, 0x0, 0x0, 0x20, 0x0, 0x0, 0x0, 0x1, 0x0, 0x4, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x10, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0xff, 0xff, 0x0, 0x0, 0xff, 0xff, 0x0, 0x0, 0xff,
		0xff, 0x0, 0x0, 0xfe, 0x7f, 0x0, 0x0, 0xfe, 0x7f, 0x0, 0x0, 0xff, 0xff,
		0x0, 0x0, 0xfe, 0x7f, 0x0, 0x0, 0xfe, 0x7f, 0x0, 0x0, 0xff, 0x3f, 0x0, 0x0,
		0xff, 0x9f, 0x0, 0x0, 0xfd, 0xdf, 0x0, 0x0, 0xfc, 0x9f, 0x0, 0x0, 0xfe,
		0x3f, 0x0, 0x0, 0xff, 0xff, 0x0, 0x0, 0xff, 0xff, 0x0, 0x0, 0xff, 0xff,
		0x0, 0x0}
)
