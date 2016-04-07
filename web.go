// main.go
package main

import(
	"fmt"
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
    err := server.ListenAndServe()

	fmt.Println(err.Error())
}
