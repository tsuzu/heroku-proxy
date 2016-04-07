// main.go
package main

import(
	"fmt"
	"os"
    "net/http"
//    "net/http/httputil"
//    "net/url"
)

func main() {
    sourceAddress := ":" + os.Getenv("PORT")
//    destinationUrlString := "http://azure.wt6.pw/"
//	destinationUrl, _ := url.Parse(destinationUrlString)
    //proxyHandler := httputil.NewSingleHostReverseProxy(destinationUrl)

	mux := http.NewServeMux()
	proxyHandler := func(rw http.ResponseWriter, req *http.Request){
		fmt.Fprintf(rw, "Hello, world")
	}

	mux.HandleFunc("/", proxyHandler)
    server := http.Server{
        Addr: sourceAddress,
        Handler: mux,
    }
    err := server.ListenAndServe()

	fmt.Println(err.Error())
}
