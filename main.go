package main

import (
	"encoding/xml"
	"flag"
	"html/template"
	"log"
	"net/http"
	"net/http/cgi"
	"os"
	"time"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

type messageSigns struct {
	Signs []sign `xml:"messageSign"`
}

func (ms messageSigns) FindByName(names ...string) []sign {
	var signs []sign
	for i := range ms.Signs {
		mSign := ms.Signs[i]
		for _, name := range names {
			if mSign.Name == name {
				signs = append(signs, mSign)
			}
		}
	}

	return signs
}

type sign struct {
	Location  string `xml:"location"`
	DmsID     string `xml:"dmsid"`
	Name      string `xml:"name"`
	Message   string `xml:"message"`
	Updated   string `xml:"updated"`
	Beacon    string `xml:"beacon"`
	Latitude  string `xml:"latitude"`
	Longitude string `xml:"longitude"`
}

func main() {
	signsLoc := flag.String("url", "https://chart.maryland.gov/rss/ProduceRss.aspx?Type=DMSXML", "The url for the Signs XML")
	certDir := flag.String("certDir", "/usr/pkg/etc/openssl/certs", "The CA cert dir")

	flag.Parse()

	err := os.Setenv("SSL_CERT_DIR", *certDir)
	checkErr(err)

	log.Fatalln(cgi.Serve(displaySignData(*signsLoc)))
}

func displaySignData(signsLoc string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		r.ParseForm()

		// Default to the usual two signs
		names := []string{"8003", "8001"}

		// if we were passed explicit names,
		// use those instead
		if len(r.Form["name"]) > 0 {
			names = r.Form["name"]
		}

		signs, err := getSigns(signsLoc, names)
		if err != nil {
			errTemplate().Execute(w, struct{ ErrMessage string }{err.Error()})
			return
		}

		resultsTemplate().Execute(w, struct{ Signs []sign }{signs})
	}
}

func getSigns(signsLoc string, names []string) ([]sign, error) {
	var ms messageSigns
	c := setupClient()

	resp, err := c.Get(signsLoc)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	err = xml.NewDecoder(resp.Body).Decode(&ms)
	if err != nil {
		return nil, err
	}

	return ms.FindByName(names...), nil
}

func setupClient() *http.Client {
	return &http.Client{
		Timeout: 3 * time.Second,
	}
}

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
