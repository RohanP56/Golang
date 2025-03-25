# URL (Uniform Resource Locator) Handling
- In golang we deal with URL using `net/url` package.
- When we store any url that mainly stored as a String.
``` Go
func main() {
	myURL := "https://jsonplaceholder.typicode.com/todos"
	fmt.Printf("Type of URL: %T\n", myURL)
}
```
### Parse URL:
- We can convert this string type data into URL, using `url.Parse()` method, which is exists inside `net/url` package.
- This method returns 2 value
    - Parse URL
    - Error
- Once we have a URL Object, now we can access various field
    - **Scheme:** Indicats the protocol (http/https)
    - **Host**: Specifies the domain name and optionally the port number
    - **Path:** Represents the path components of the URL, which specifies the resource's location on the server.
    - **RawQuery:** Contains the raw query sting, including query params.
``` Go
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
}
```

### Modify URL:
- Like parse URL, we can also modify the URLs.
- After modifying `URL` convert it into `String`, so we can re-use it.
``` Go
func main() {
    //Modify URL
	parseURL.Path = "/newPath"
	parseURL.Host = "rohan.com"
	fmt.Println(parseURL)

	// Constructing new URL
	newURL := parseURL.String()
	fmt.Println(newURL)
}
```