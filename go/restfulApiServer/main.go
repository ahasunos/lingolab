package main

import (
	"fmt"
	"net/http"
)

func office(responseWriter http.ResponseWriter, request *http.Request) {
	// Fprintf(w io.Writer, format string, a ...any) (n int, err error)
	// Fprintf formats according to a format specifier and writes to w.
	// It returns the number of bytes written and any write error encountered.
	fmt.Fprintf(responseWriter, "Homepage live: Where APIs flow smoother than Michael's awkward office moments!")
}

func requestHandler() {
	// HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
	// HandleFunc registers the handler function for the given pattern in the DefaultServeMux.
	http.HandleFunc("/", office)

	// ListenAndServe(addr string, handler http.Handler) error
	// ListenAndServe listens on the TCP network address addr and then calls Serve with handler to handle requests on incoming connections.
	// Accepted connections are configured to enable TCP keep-alives.
	// The handler is typically nil, in which case the DefaultServeMux is used.
	// ListenAndServe always returns a non-nil error.

	http.ListenAndServe("127.0.0.1:8080", nil) // Using DefaultServeMux, hence sending nil for second parameters
}

func main() {
	requestHandler()
}
