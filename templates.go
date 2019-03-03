package main

import "html/template"

func resultsTemplate() *template.Template {
	tmplStr := `
		<!DOCTYPE html>
		<html>
			<head>
				<title>Signs</title>
				<style>
					.sign-container {
						display: flex;
						flex-direction: column;
						width: 100%;
						height: 100vh;
					}
		
					.sign {
						background-color: #201c29;
						padding: 2em;
						text-align: center;
						color: #a1a1af;
						margin-bottom: 1em;
						flex-grow: 1;
						display: flex;
						flex-direction: column;
						align-items: center;
						justify-content: center;
						border-radius: 1em;
					}
		
					.sign .text {
						margin-top: 0.5em;
						margin-bottom: 0.5em;
					}
		
					.sign .message {
						font-size: 4vh;
					}

					.sign .location {
						font-size: 3vh;
					}

					.sign .time {
						font-size: 2.5vh;
					}
				</style>
			</head>
			<body>
				<div class="sign-container">
					{{range .Signs}}
						<div class="sign">
							<p class="location text">{{.Location}}</p>
							{{range .MessageLines}}
								<p class="message text">{{.}}</p>
							{{end}}
							<p class="time text">Last Updated: {{.Updated}}</p>
						</div>
					{{else}}
						<div class="sign">
							<p class="message text">No results!</p>
						</div>
					{{end}}
				</div>
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
