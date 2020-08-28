package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// Em Go, a multiplexação de canais é unir a leitura de múltiplos canais num único canal

func foward(source <-chan string, target chan string) {
	for {
		target <- <-source // encaminhando para o target o dado recebido em source
		// enquanto houver dados em source, o loop irá fazer iterações
		// caso não haja mais dados em source, a operação ficará bloqueada (aguardando por novos dados)
	}
}

// multiplexar
func join(c1, c2 <-chan string) <-chan string {
	c := make(chan string)
	go foward(c1, c)
	go foward(c2, c)

	return c
}

// generator
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
	c := join(
		title("https://www.cod3r.com.br", "https://google.com"),
		title("https://amazon.com.br", "https://youtube.com"),
	)

	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
}
