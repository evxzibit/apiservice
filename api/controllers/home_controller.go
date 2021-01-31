package controllers

import (
	"net/http"

	"apiservice/api/responses"
)

// Home the landing home page
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Hey ha")
}