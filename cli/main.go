package main

import (
	"flag"
	"fmt"
	"net/url"
)

func main() {
	decode_mode := flag.Bool("d", false, "Enable debug mode")
	flag.Parse()
	if *decode_mode {
		fmt.Println("Enter a valid URL: ")
		var user_input string

		_, err := fmt.Scan(&user_input)
		if err != nil {
			fmt.Println("User input error:", err)
			return
		}
		user_url, err := url.Parse(user_input)
		if err != nil {
			fmt.Println("Parsing error", err)
		}
		if isValidURL(user_url) {
			if !urlPathIsEmpty(user_url) {
				decodedPath, err := url.QueryUnescape(user_url.Path)
				if err != nil {
					return
				}
				fmt.Printf("URL with path decoded:\n%s://%s%s\n", user_url.Scheme, user_url.Host, decodedPath)
			} else {
				fmt.Println("Path is empty, nothing to decode.")
			}
		} else {
			fmt.Println("Invalid URL!")
			return
		}
	}

}

func isValidURL(u *url.URL) bool {
	return u.Scheme != "" && u.Host != ""
}

func urlPathIsEmpty(u *url.URL) bool {
	return u.Path == "" || u.Path == "/"
}
