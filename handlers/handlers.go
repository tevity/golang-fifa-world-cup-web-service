package handlers

import (
	"golang-fifa-world-cup-web-service/data"
	"net/http"
)

// RootHandler returns an empty body status code
func RootHandler(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusNoContent)
}

// ListWinners returns winners from the list
func ListWinners(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	year := req.URL.Query().Get("year")
	if winners, err := data.ListAllJSON(); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
	} else {
		if year == "" {
			res.Write(winners)
		} else {
			if filteredWinners, err := data.ListAllByYear(year); err != nil {
				res.WriteHeader(http.StatusBadRequest)
			} else {
				res.Write(filteredWinners)
			}
		}
	}
}

// AddNewWinner adds new winner to the list
func AddNewWinner(res http.ResponseWriter, req *http.Request) {

}

// WinnersHandler is the dispatcher for all /winners URL
func WinnersHandler(res http.ResponseWriter, req *http.Request) {

}
