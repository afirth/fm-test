package api

import (
	"io"
	"net/http"

	"github.com/afirth/fm-test/gbdx"
	"github.com/afirth/fm-test/transcode"
)

// Please excuse the logging - I would add nice json log handlers for prod, e.g zap or whatever you're using

// SearchHandler executes a search for catalog items inside a geojson polygon.
// The client should already be authenticated
func (client *Client) SearchHandler(w http.ResponseWriter, r *http.Request) {
	// geojson -> wkt
	wkt, err := transcode.GeoJSON2WKT(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// execute search
	c, err := gbdx.NewCatalogSearch(client.gbdx, wkt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return search results as json
	err = c.EncodeJSON(w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// HealthCheckHandler returns ok when the server is alive
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	// can expand or adapt to /metrics
	io.WriteString(w, `{"alive": true}`)
}
