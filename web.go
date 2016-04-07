// main.go
package main

import(
	"os"
    "net/http"
    "net/http/httputil"
    "net/url"
)

func main() {
    sourceAddress := ":" + os.Getenv("PORT")
    destinationUrlString := "http://azure.wt6.pw/"
    destinationUrl, _ := url.Parse(destinationUrlString)
    proxyHandler := httputil.NewSingleHostReverseProxy(destinationUrl)
    server := http.Server{
        Addr: sourceAddress,
        Handler: proxyHandler,
    }
    server.ListenAndServe()
}
