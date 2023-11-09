## Notes

(A very rough version of gibberish notes)

- To implement a basic server, I used the net/http package.
The skeleton for a basic server must have a function which 
1. handles the request a given pattern in the server and redirects to a function.

```
// This was done with the help of http.HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
	
//Example:

// HandleFunc registers the handler function for the given pattern in the DefaultServeMux.
http.HandleFunc("/", office)

//and here office is a handler function of type handler func(http.ResponseWriter, *http.Request)

//Example:
func office(responseWriter http.ResponseWriter, request *http.Request) {
// Fprintf(w io.Writer, format string, a ...any) (n int, err error)
// Fprintf formats according to a format specifier and writes to w.
// It returns the number of bytes written and any write error encountered.
fmt.Fprintf(responseWriter, "Homepage live: Where APIs flow smoother than Michael's awkward office moments!")
}
```

2. listen and serves to an address.

```
// ListenAndServe(addr string, handler http.Handler) error
// ListenAndServe listens on the TCP network address addr and then calls Serve with handler to handle requests on incoming connections.
// Accepted connections are configured to enable TCP keep-alives.
// The handler is typically nil, in which case the DefaultServeMux is used.
// ListenAndServe always returns a non-nil error.

http.ListenAndServe("127.0.0.1:8080", nil) // Using DefaultServeMux, hence sending nil for second parameters
```

- To write respond back to the client, we create a new encode which takes a writer as argument and the encode the data (characters in this case).
```
// NewEncoder(w io.Writer) *json.Encoder
// NewEncoder returns a new encoder that writes to w.
jsonEncoder := json.NewEncoder(responseWriter)

// Encode(v any) error
// Encode writes the JSON encoding of v to the stream, followed by a newline character.
jsonEncoder.Encode(characters)
```