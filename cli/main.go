package main

import "fmt"
import "net/url"

func main() {
	fmt.Println("Enter a valid URL: ")
	var user_input string

	_, err := fmt.Scan(&user_input)
	if err != nil {
		fmt.Println("User input error:", err)
		return
	}
	if IsUrl(user_input) {
		fmt.Println("User inputted: ", user_input)
	} else {
		fmt.Println("Invalid URL!")
		return
	}
}

func IsUrl(str string) bool {
    u, err := url.Parse(str)
    return err == nil && u.Scheme != "" && u.Host != ""
}
