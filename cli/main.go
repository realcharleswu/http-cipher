package main

import (
	"fmt"
	"net/url"
)

func main() {
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
			fmt.Println("path has something", user_url.Path)
		} else {
			fmt.Println("Path has nothing")
		}
	} else {
		fmt.Println("Invalid URL!")
		return
	}

}

func isValidURL(u *url.URL) bool {
	return u.Scheme != "" && u.Host != ""
}

func urlPathIsEmpty(u *url.URL) bool {
	return u.Path == "" || u.Path == "/"
}
