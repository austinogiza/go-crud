package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (app *application) statusHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Status")

	currentStatus := AppStatus{
		Status:      "Available",
		Environment: app.config.env,
		Version:     version,
	}

	js, err := json.MarshalIndent(currentStatus, "", "\t")

	if err != nil {
		app.logger.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}
