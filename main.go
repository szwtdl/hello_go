package main

import (
	"fmt"
	client "github.com/szwtdl/req"
)

func main() {
	fmt.Println("Hello, World!")
	httpClient := client.NewHttpClient("https://www.baidu.com")
	httpClient.SetHeader(map[string]string{
		"User-Agent":                "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3",
		"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
		"Accept-Language":           "en-US,en;q=0.5",
		"Accept-Encoding":           "gzip, deflate, br",
		"Connection":                "keep-alive",
		"Upgrade-Insecure-Requests": "1",
		"Cache-Control":             "max-age=0",
		"Pragma":                    "no-cache",
		"Sec-Fetch-Dest":            "document",
	})
	resp, err := httpClient.DoGet("index.php")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(resp))
}
