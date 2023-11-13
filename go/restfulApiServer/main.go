package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Character struct {
	Id       int64  `json:"Id"`
	Name     string `json:"Name"`
	RealName string `json:"RealName"`
	Role     string `json:"Role"`
}

type Characters []Character

var CharactersDB Characters

func getCharacters(responseWriter http.ResponseWriter, request *http.Request) {
	characters := CharactersDB
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

// Dummy payload for post api
// {
// 	"Id": 4
// 	"Name": "Pam Beesly (Pamela)",
// 	"RealName": "Jenna Fischer",
// 	"Role": "Receptionist"
// }

// {
//  "Id": 5
// 	"Name": "Kevin Malone",
// 	"RealName": "Brian Baumgartner",
// 	"Role": "Accountant"
// }

func postCharacters(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Println("Post endpoint hit")

	// func (http.ResponseWriter).Header() http.Header
	// Header returns the header map that will be sent by WriteHeader. The Header map also is the mechanism with which Handlers can set HTTP trailers.

	// Set(key string, value string)
	// Set sets the header entries associated with key to the single element value.
	// It replaces any existing values associated with key.
	// The key is case insensitive; it is canonicalized by textproto.CanonicalMIMEHeaderKey.
	// To use non-canonical keys, assign to the map directly.

	responseWriter.Header().Set("Content-Type", "application/json")

	var character Character

	// NewDecoder(r io.Reader) *json.Decoder
	// NewDecoder returns a new decoder that reads from r.
	// The decoder introduces its own buffering and may read data from r beyond the JSON values requested.
	jsonDecoder := json.NewDecoder(request.Body)

	// func (*json.Decoder).Decode(v any) error
	// Decode reads the next JSON-encoded value from its input and stores it in the value pointed to by v.
	err := jsonDecoder.Decode(&character)
	if err != nil {
		responseWriter.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(responseWriter, "Bad request, post valid payload!")
		return
	}

	CharactersDB = append(CharactersDB, character)
	json.NewEncoder(responseWriter).Encode(character)
}

func office(responseWriter http.ResponseWriter, request *http.Request) {
	// Fprintf(w io.Writer, format string, a ...any) (n int, err error)
	// Fprintf formats according to a format specifier and writes to w.
	// It returns the number of bytes written and any write error encountered.
	fmt.Fprintf(responseWriter, "Homepage live: Where APIs flow smoother than Michael's awkward office moments!")
}

func requestHandler() {

	// Using gorilla mux, as it provides verbs
	// Earlier using any http verbs on the APIs were returning values

	// func mux.NewRouter() *mux.Router
	// NewRouter returns a new router instance.
	// StrictSlash defines the trailing slash behavior for new routes.
	// The initial value is false.

	// When true, if the route path is "/path/",
	// accessing "/path" will perform a redirect to the former and vice versa.
	// In other words, your application will always see the path as specified in the route.

	httpRouter := mux.NewRouter().StrictSlash(true)

	// As provided by default net/http package:
	// HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
	// HandleFunc registers the handler function for the given pattern in the DefaultServeMux.

	// For gorilla mux:
	// HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) *mux.Route
	// HandleFunc registers a new route with a matcher for the URL path. See Route.Path() and Route.HandlerFunc().
	httpRouter.HandleFunc("/", office).Methods("GET")
	httpRouter.HandleFunc("/characters", getCharacters).Methods("GET")
	httpRouter.HandleFunc("/addCharacters", postCharacters).Methods("POST")

	// ListenAndServe(addr string, handler http.Handler) error
	// ListenAndServe listens on the TCP network address addr and then calls Serve with handler to handle requests on incoming connections.
	// Accepted connections are configured to enable TCP keep-alives.
	// The handler is typically nil, in which case the DefaultServeMux is used.
	// ListenAndServe always returns a non-nil error.

	http.ListenAndServe("0.0.0.0:8080", httpRouter) // Using DefaultServeMux, hence sending nil for second parameters
}

func main() {
	loadCharacters()
	fmt.Println("Server is starting...")
	requestHandler()
}

func loadCharacters() {
	// Ideally, this should have been fetched from database or something!
	CharactersDB = append(CharactersDB, Character{
		Id:       1,
		Name:     "Michael Scott",
		RealName: "Steve Carell",
		Role:     "Regional Manager",
	},
		Character{
			Id:       2,
			Name:     "Jim Halpert",
			RealName: "John Krasinski",
			Role:     "Sales Representative",
		},
		Character{
			Id:       3,
			Name:     "Dwight Schrute",
			RealName: "Rainn Wilson",
			Role:     "Assistant (to the) Regional Manager",
		})
}
