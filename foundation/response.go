package response

import "net/http"

func Response(w http.ResponseWriter, response []byte, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	w.Write(response)
}
