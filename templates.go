package main

import "html/template"

func resultsTemplate() *template.Template {
	tmplStr := `
		<!DOCTYPE html>
		<html>
			<head>
				<title>Signs</title>
			</head>
			<body>
				{{range .Signs}}
					<p>{{.Message}}</p>
				{{else}}
					<p>No results!</p>
				{{end}}
			</body>
		</html>
	`
	return template.Must(template.New("results").Parse(tmplStr))
}

func errTemplate() *template.Template {
	tmplStr := `
		<!DOCTYPE html>
		<html>
			<head>
				<title>Signs</title>
			</head>
			<body>
				<p>{{.ErrMessage}}</p>
			</body>
		</html>
	`

	return template.Must(template.New("error").Parse(tmplStr))
}
