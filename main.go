package main

import (
	"encoding/xml"
	"flag"
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

		// Default to the usual signs
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

		dispSigns := toDisplaySigns(signs)

		resultsTemplate().Execute(w, struct{ Signs []displaySign }{dispSigns})
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
		Timeout: 25 * time.Second,
	}
}
