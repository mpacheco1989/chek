package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

// A Response struct to map the Entire Response
type Response struct {
	Name    string    `json:"name"`
	Pokemon []Pokemon `json:"pokemon_entries"`
}

// A Pokemon Struct to map every pokemon to.
type Pokemon struct {
	EntryNo int            `json:"entry_number"`
	Species PokemonSpecies `json:"pokemon_species"`
}

type PokemonSpecies struct {
	Name string `json:"name"`
}

// pokemon struct (Model)
type Pokemonout struct {
	ID     string `json:"id"`
	Nombre string `json:"nombre"`
}

// Get all pokemons
func getPokemones(w http.ResponseWriter, r *http.Request) {
	// Init books var as a slice Book struct
	var pokemones []Pokemonout
	fmt.Println("all pokemones")

	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response

	json.Unmarshal(responseData, &responseObject)
	for i := 0; i < len(responseObject.Pokemon); i++ {

		pokemones = append(pokemones, Pokemonout{ID: strconv.Itoa(responseObject.Pokemon[i].EntryNo), Nombre: responseObject.Pokemon[i].Species.Name})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pokemones)

}

// Get single pokemon
func getPokemon(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r) // Gets params
	index, err := strconv.Atoi(params["id"])
	var pokemones []Pokemonout

	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response

	json.Unmarshal(responseData, &responseObject)
	pokemones = append(pokemones, Pokemonout{ID: strconv.Itoa(responseObject.Pokemon[index-1].EntryNo), Nombre: responseObject.Pokemon[index-1].Species.Name})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pokemones)
}

func main() {
	// iniciando router
	r := mux.NewRouter()
	r.Use(corsMiddleware)

	// Route handles & endpoints
	r.HandleFunc("/pokemones", getPokemones).Methods("GET")
	r.HandleFunc("/pokemones/{id}", getPokemon).Methods("GET")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")                                                            // 允许访问所有域，可以换成具体url，注意仅具体url才能带cookie信息
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token") //header的类型
		w.Header().Add("Access-Control-Allow-Credentials", "true")                                                    //设置为true，允许ajax异步请求带cookie信息
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")                             //允许请求方法
		w.Header().Set("content-type", "application/json;charset=UTF-8")                                              //返回数据格式是json
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}
