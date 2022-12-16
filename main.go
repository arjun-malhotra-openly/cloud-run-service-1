package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/submit", callService2)
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside Index")
	html := `<!DOCTYPE html>
	<html lang="en">
		<head>
			<meta charset = "UTF-8">
			<title>JWT&Cookie Example</title>
		</head>
		<body>
			<p> Cloud Run Service 1 </p>
            <form action="/submit" method="post">
                <input type="submit" />
            </form>
		</body>
	</html>`
	io.WriteString(res, html)
}

func callService2(res http.ResponseWriter, req *http.Request) {
	resp, _ := http.Get("https://arjun-temp-service-2-5amxaxbpha-uc.a.run.app")
	respBytes, _ := io.ReadAll(resp.Body)
	io.WriteString(res, string(respBytes))
}
