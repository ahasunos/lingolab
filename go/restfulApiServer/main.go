package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Character struct {
	Name     string `json:"Name"`
	RealName string `json:"RealName"`
	Role     string `json:"Role"`
}

type Characters []Character

func getCharacters(responseWriter http.ResponseWriter, request *http.Request) {
	characters := fetchCharacters()
	fmt.Println("getCharacters endpoint hit")

	// NewEncoder(w io.Writer) *json.Encoder
	// NewEncoder returns a new encoder that writes to w.
	jsonEncoder := json.NewEncoder(responseWriter)

	// Encode(v any) error
	// Encode writes the JSON encoding of v to the stream, followed by a newline character.
	jsonEncoder.Encode(characters)

	// The above could have been achieved in a single line like below:
	// I broken down so that I could know about both NewEncoder and Encode
	// json.NewEncoder(responseWrite).Encode(characters)
}

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
	http.HandleFunc("/characters", getCharacters)

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

func fetchCharacters() Characters {
	// Ideally, this should have been fetched from database or something!
	return Characters{
		Character{
			Name:     "Michael Scott",
			RealName: "Steve Carell",
			Role:     "Regional Manager",
		},
		Character{
			Name:     "Jim Halpert",
			RealName: "John Krasinski",
			Role:     "Sales Representative",
		},
		Character{
			Name:     "Dwight Schrute",
			RealName: "Rainn Wilson",
			Role:     "Assistant (to the) Regional Manager",
		},
	}
}
