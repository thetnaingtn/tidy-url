package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/thetnaingtn/tidy-url/core"
	"github.com/thetnaingtn/tidy-url/store"
)

type Handlers struct {
	core core.Core
}

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
