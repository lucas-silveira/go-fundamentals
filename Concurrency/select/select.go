package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

func theFastest(url1, url2, url3 string) string {
	c1 := title(url1)
	c2 := title(url2)
	c3 := title(url3)

	// estrutura de controle específica para concorrência
	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(400 * time.Millisecond):
		return "Todos perderam"
		// default:
		// 	return "Ainda sem resposta"
	}
}

func title(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			c <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}

	return c
}

func main() {
	champion := theFastest(
		"https://google.com",
		"https://amazon.com.br",
		"https://youtube.com",
	)

	fmt.Println(champion)
}
