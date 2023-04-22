package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/adityarudrawar/go-gorilla/database"
	"github.com/adityarudrawar/go-gorilla/structures"
	"github.com/adityarudrawar/go-gorilla/validations"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func YourHandlerGET(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonRespone, err := json.Marshal(database.Gorillas)
	if err != nil {
		return
	}
	w.Write(jsonRespone)
}

func YourHandlerPOST(w http.ResponseWriter, r *http.Request) {
	var inputGorilla structures.InputGorilla

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	log.Println("Body is: ", string(body))

	err = json.Unmarshal(body, &inputGorilla)
	if err != nil {
		log.Println("Error while unmarshaling")
	}

	log.Println("Unmarshalled Body is: ", inputGorilla)

	resp := validations.ValidInputGorillaBody(&inputGorilla)

	if resp.IsError {
		w.Write([]byte("BAD Gorilla!\n"))
	} else {
		// Write to database
		if (len(database.Gorillas) < 10){
			database.Gorillas = append(database.Gorillas, resp)
		}
		
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		jsonRespone, err := json.Marshal(resp)
		if err != nil {
			return
		}
		w.Write(jsonRespone)
	}
	// https://pkg.go.dev/net/http#pkg-examples
	// https://medium.com/geekculture/develop-rest-apis-in-go-using-gorilla-mux-5869b2179a18
}

func RootHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("[GET] /gorilla!\r\n[POST]/gorilla => See the readme"))
}

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", RootHandler).Methods("GET")
	r.HandleFunc("/gorilla", YourHandlerGET).Methods("GET")
	r.HandleFunc("/gorilla", YourHandlerPOST).Methods("POST")
	r.Use(mux.CORSMethodMiddleware(r))

	log.Println("Started Server")

	godotenv.Load()

	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":" + port, r))

}
