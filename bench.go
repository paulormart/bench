// Bench simple benchmark an HTTP endpoint
package bench

import (
	"time"
	"net/http"
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"io"
)

func Url(url string) (report string, err error) {
	start := time.Now()
	if !strings.HasPrefix(url, "http://") {
		url = fmt.Sprintf("http://%s", url)
	}
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		return report, err
	}
	// read content body
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		return report, err
	}
	report = fmt.Sprintf("%.4fs  %d  %s %v", time.Since(start).Seconds(), nbytes, url, resp.Status)
	return report, err
}

func urlAsync(url string, ch chan<- string){
	report, err := Url(url)
	if err != nil {
		ch<- fmt.Sprintf("error while reading %s: %v", url, err)
	}
	ch<-report
}

func Urls(urls []string) (report string, err error) {
	start := time.Now()
	ch := make(chan string)
	for _, url := range urls {
		// start go routine
		go urlAsync(url, ch)
	}
	for range urls {
		// receive from channel
		output := <-ch
		fmt.Println(output)
	}
	report = fmt.Sprintf("%.4fs  %d", time.Since(start).Seconds(), len(urls))
	return report, err
}