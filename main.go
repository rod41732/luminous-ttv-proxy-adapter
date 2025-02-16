package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

const defaultQueryString = "?platform=web&allow_source=true&allow_audio_only=true&fast_bread=true&warp=true&supported_codecs=av1%2ch265%2ch264"
const defaultLumiHost = "https://luminous.alienpls.org"

// proxy the request to actual luminous TTV server
// path -> name : channel name, possibly with .m3u8 suffix (e.g. forsen or forsen.m3u8)
// query -> host: proto://host of luminous service, defaults to `defaultLumiHost` which is a reasonable lumi server
// query -> params: query string to be sent with URL, default to `defaultQueryString` which sets bunch of low latency/encoding related flags
//
// Response: this will proxy response, status, headers from upstream
func proxyHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)

	channelNameRaw := r.PathValue("name")
	fmt.Printf("NameRaw is %#v\n", channelNameRaw)
	channelName := strings.TrimSuffix(channelNameRaw, ".m3u8")
	fmt.Printf("Name is %#v\n", channelName)

	query := r.URL.Query()
	host := query.Get("host")
	if host == "" {
		host = defaultLumiHost
	}
	host = strings.TrimRight(host, "/")

	params := query.Get("params")
	if params != "" {
		if params[0] != '?' {
			params = "?" + params
		}
	} else {
		params = defaultQueryString
	}

	buildURL := host + "/live/" + channelName + params
	fmt.Printf("URL will be %#v\n", buildURL)

	// proxy request
	resp, err := http.Get(buildURL)
	if err != nil {
		http.Error(w, "Failed to fetch remote content", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Copy headers
	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	// Set the status code and copy the response body
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func main() {
	http.HandleFunc("/proxy/{name}", proxyHandler)
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
