package handlers

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	response "github.com/thetnaingtn/tidy-url/foundation"
	"github.com/thetnaingtn/tidy-url/store"
)

type Response struct {
	ShortURL string `json:"short_url"`
}

func (h Handlers) Tidy(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var payload store.Payload

	body, err := io.ReadAll(r.Body)
	if err != nil {

	}

	if err := json.Unmarshal(body, &payload); err != nil {

	}

	shortURL, err := h.core.GenerateTidyUrl(payload)
	if err != nil {

	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	resp := Response{ShortURL: shortURL}

	response, err := json.Marshal(resp)

	if err != nil {

	}

	w.Write(response)
}

func (h Handlers) Expand(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	encodedString := params.ByName("id")
	tidyurl, err := h.core.GetLongURL(encodedString)

	if err == sql.ErrNoRows {
		log.Println(err)
		message := struct {
			Message string `json:"message"`
		}{
			Message: "No Proper URL found for the given URL",
		}
		resp, _ := json.Marshal(message)
		response.Response(w, resp, http.StatusNotFound)
		return
	}

	if err != nil {
		log.Println(err)
	}

	http.Redirect(w, r, tidyurl.LongURL, http.StatusMovedPermanently)
}
