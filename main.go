package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"golang.org/x/oauth2"
	auth "golang.org/x/oauth2/google"
	"google.golang.org/api/idtoken"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/submita", callService2a)
	http.HandleFunc("/submitb", callService2b)
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
			<p> Cloud Run Service 1a </p>
            <form action="/submita" method="get">
				<input type="submit" />
			</form>
			<p> Cloud Run Service 1b </p>
			<form action="/submitb" method="get">
				<input type="submit" />
            </form>
		</body>
	</html>`
	io.WriteString(res, html)
}

func callService2a(res http.ResponseWriter, req *http.Request) {
	// resp, _ := http.Get("https://arjun-temp-service-2-5amxaxbpha-uc.a.run.app")
	// respBytes, _ := io.ReadAll(resp.Body)
	// io.WriteString(res, string(respBytes))
	var client http.Client
	ctx := context.Background()
	audience := "https://arjun-temp-service-2-5amxaxbpha-uc.a.run.app"
	ts, err := idtoken.NewTokenSource(ctx, audience)
	if err != nil {
		log.Println("err in getting a new token source", err)
	}
	token, err := ts.Token()
	if err != nil {
		log.Println("err in getting a new token from token source", err)
	}
	req, _ = http.NewRequest(http.MethodGet, audience, nil)
	token.SetAuthHeader(req)
	resp, err := client.Do(req)
	if err != nil {
		log.Println("err in getting a response.", err)
	}
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("err in reading the response body.", err)
	}
	fmt.Printf("submita %v+", token)
	log.Println(string(respBytes))
	io.WriteString(res, string(respBytes))

}

func callService2b(res http.ResponseWriter, req *http.Request) {
	ctx := context.Background()
	var client http.Client
	var token *oauth2.Token
	scopes := []string{
		"https://www.googleapis.com/auth/cloud-platform",
	}
	audience := "https://arjun-temp-service-2-5amxaxbpha-uc.a.run.app"

	credentials, err := auth.FindDefaultCredentials(ctx, scopes...)
	if err == nil {
		token, err = credentials.TokenSource.Token()
		if err != nil {
			log.Println(err)
		}
	}
	reqService2, _ := http.NewRequest(http.MethodGet, audience, nil)
	reqService2.Header.Set("Authorization", token.Type()+" "+token.Extra("id_token").(string))
	resp, err := client.Do(reqService2)
	if err != nil {
		log.Println("err in getting a response.", err)
	}
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("err in reading the response body.", err)
	}
	fmt.Printf("submitb %v+", token)
	log.Println(string(respBytes))
	io.WriteString(res, string(respBytes))

}

// func makeGetRequest(res http.ResponseWriter, req *http.Request) {
// 	// Example `audience` value (Cloud Run): https://my-cloud-run-service.run.app/
// 	// (`targetURL` and `audience` will differ for non-root URLs and GET parameters)
// 	audience := "https://arjun-temp-service-2-5amxaxbpha-uc.a.run.app"
// 	targetURL := "https://arjun-temp-service-2-5amxaxbpha-uc.a.run.app"
// 	ctx := context.Background()

// 	// client is a http.Client that automatically adds an "Authorization" header
// 	// to any requests made.
// 	client, err := idtoken.NewClient(ctx, audience)
// 	if err != nil {
// 		fmt.Errorf("idtoken.NewClient: %v", err)
// 	}

// 	resp, err := client.Get(targetURL)
// 	if err != nil {
// 		fmt.Errorf("client.Get: %v", err)
// 	}
// 	defer resp.Body.Close()
// 	if _, err := io.Copy(res, resp.Body); err != nil {
// 		fmt.Errorf("io.Copy: %v", err)
// 	}
// }
