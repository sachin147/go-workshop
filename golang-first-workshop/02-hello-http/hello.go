package main

// TODO: import http
import( 
    "fmt"
	"net/http"
)
    
// TODO: create hello handler

func HelloServer(w http.ResponseWriter, req *http.Request) {
	//io.WriteString(w, "hello, world!\n")
    fmt.Fprintln(w, "Hello World")
}

func main() {    	
	// setup handler for /hello
    http.HandleFunc("/hello", HelloServer)
    http.ListenAndServe("127.0.0.1:8000", nil)
	// start http server
}
