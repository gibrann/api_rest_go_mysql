package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func setupRoutesForPedidos(router *mux.Router) {
	// First enable CORS. If you don't need cors, comment the next line
	enableCORS(router)

	router.HandleFunc("/Pedidos", func(w http.ResponseWriter, r *http.Request) {
		Pedidos, err := getPedidos()
		if err == nil {
			respondWithSuccess(Pedidos, w)
		} else {
			respondWithError(err, w)
		}
	}).Methods(http.MethodGet)
	router.HandleFunc("/Pedido/{id}", func(w http.ResponseWriter, r *http.Request) {
		idAsString := mux.Vars(r)["id"]
		id, err := stringToInt64(idAsString)
		if err != nil {
			respondWithError(err, w)
			// We return, so we stop the function flow
			return
		}
		Pedido, err := getPedidoById(id)
		if err != nil {
			respondWithError(err, w)
		} else {
			respondWithSuccess(Pedido, w)
		}
	}).Methods(http.MethodGet)

	router.HandleFunc("/Pedido", func(w http.ResponseWriter, r *http.Request) {
		// Declare a var so we can decode json into it
		var Pedido Pedido
		err := json.NewDecoder(r.Body).Decode(&Pedido)
		if err != nil {
			respondWithError(err, w)
		} else {
			err := createPedido(Pedido)
			if err != nil {
				respondWithError(err, w)
			} else {
				respondWithSuccess(true, w)
			}
		}
	}).Methods(http.MethodPost)

	router.HandleFunc("/Pedido", func(w http.ResponseWriter, r *http.Request) {
		// Declare a var so we can decode json into it
		var Pedido Pedido
		err := json.NewDecoder(r.Body).Decode(&Pedido)
		if err != nil {
			respondWithError(err, w)
		} else {
			err := updatePedido(Pedido)
			if err != nil {
				respondWithError(err, w)
			} else {
				respondWithSuccess(true, w)
			}
		}
	}).Methods(http.MethodPut)
	router.HandleFunc("/Pedido/{id}", func(w http.ResponseWriter, r *http.Request) {
		idAsString := mux.Vars(r)["id"]
		id, err := stringToInt64(idAsString)
		if err != nil {
			respondWithError(err, w)
			// We return, so we stop the function flow
			return
		}
		err = deletePedido(id)
		if err != nil {
			respondWithError(err, w)
		} else {
			respondWithSuccess(true, w)
		}
	}).Methods(http.MethodDelete)
}

func enableCORS(router *mux.Router) {
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", AllowedCORSDomain)
	}).Methods(http.MethodOptions)
	router.Use(middlewareCors)
}
func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			// Just put some headers to allow CORS...
			w.Header().Set("Access-Control-Allow-Origin", AllowedCORSDomain)
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			// and call next handler!
			next.ServeHTTP(w, req)
		})
}

// Helper functions for respond with 200 or 500 code
func respondWithError(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(err.Error())
}

func respondWithSuccess(data interface{}, w http.ResponseWriter) {

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
