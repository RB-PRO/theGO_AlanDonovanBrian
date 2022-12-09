package main

/*
go run 1/10/main.go http://www.labnet.ru/ https://pkg.go.dev/std https://pkg.go.dev/github.com/Azure/go-autorest/autorest https://pkg.go.dev/net/http
*/

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main2() {
	start := time.Now()
	url := "http://webcode.me"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err) // Отправка в канал ch
	}
	defer resp.Body.Close()

	bodyString, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bodyString))
	fileOut, errFile := os.Create("1/10/" + strconv.Itoa(1))
	if errFile != nil {
		panic(errFile)
	}
	if _, errOut := fileOut.WriteString(string(bodyString)); errOut != nil {
		panic(errOut)
	}
	fileOut.Close()

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		fmt.Println(fmt.Sprintf("while reading %s: %v", url, err))
	}
	secs := time.Since(start).Seconds()
	fmt.Println(fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url))

}

func main() {
	start := time.Now()
	ch := make(chan string)

	var urls [3]string
	for i := range urls {
		urls[i] = "https://fix-price.com/catalog/produkty-i-napitki?page=" + strconv.Itoa(i+1)
	}
	for ind, url := range urls {
		go fetch(url, ch, ind) // Запуск go-подпрограммы
	}
	fmt.Println(<-ch)
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
func fetch(url string, ch chan<- string, i int) {
	fmt.Println("start", i)
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // Отправка в канал ch
		return
	}
	if resp.StatusCode == http.StatusOK {
		bodyString, err := ioutil.ReadAll(resp.Body)

		//fmt.Println(bodyString)

		fileOut, errFile := os.Create("1/10/" + strconv.Itoa(i))
		if errFile != nil {
			panic(errFile)
		}

		if _, errOut := fileOut.WriteString(string(bodyString)); errOut != nil {
			panic(errOut)
		}
		fileOut.Close()
		// ***
		nbytes, err := io.Copy(ioutil.Discard, resp.Body)
		if err != nil {
			ch <- fmt.Sprintf("while reading %s: %v", url, err)
			return
		}
		secs := time.Since(start).Seconds()
		ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
		//output, errs := io.Copy(os.Stdout, resp.Body)

		// ***

	}
	resp.Body.Close() // Исключение утечки ресурсов
}
