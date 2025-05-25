package handlers

import (

    "net/http"


)
func ParamsHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../src/templates/parameters.html")
}
