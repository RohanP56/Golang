package main

import (
	"fmt"
	"net/url"
)

func main() {
	myURL := "https://www.youtube.com/watch?v=vu6ZQ-t1sUk&list=PLzjZaW71kMwSEVpdbHPr0nPo5zdzbDulm&index=26&t=305s"
	fmt.Printf("Type of URL: %T\n", myURL)

	//Convert this string into URL
	parseURL, err := url.Parse(myURL)
	if err != nil {
		fmt.Println("Error occur while parseing the value:", err)
	}
	fmt.Printf("Type of URL: %T\n", parseURL) // *url.URL
	fmt.Println("Scheme:", parseURL.Scheme)
	fmt.Println("Path:", parseURL.Path)
	fmt.Println("Host:", parseURL.Host)
	fmt.Println("RawQuey:", parseURL.RawQuery)

	//Modify URL
	parseURL.Path = "/newPath"
	parseURL.Host = "rohan.com"
	fmt.Println(parseURL)

	// Constructing new URL
	newURL := parseURL.String()
	fmt.Println(newURL)
}
