package gbdx

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// Do searches the GBDX catalog for all records matching an area inside a single polygon represented as WKT. It writes the catalogID, platformName, and timestamp of each result [maximum 1000](https://gbdxdocs.digitalglobe.com/docs/catalog-v2-course#section-searching-the-catalog) as json

// This could be extended for filters, types, and dates
// Contains the payload sent to GBDX
type catalogSearchRequest struct {
	WKT string `json:"searchAreaWkt"`
}

// Individual results
// Per the spec, we want CatalogID in the output. Lots of ways to do this
// Chose to use two another json lib which supports custom tags, which is ugly but simple
type catalogItem struct {
	Identifier string `json:"identifier"`
	Properties struct {
		PlatformName string `json:"platformName"`
		Timestamp    string `json:"timestamp"`
		CatalogID    string `json:"catalogID"`
	} `json:"properties"`
}

// A whole result from catalog/v2/search
type catalogSearchResponse struct {
	// Stats   map[string]interface{} //omitted because we don't want it in the output
	Results []catalogItem
}

// CatalogSearch retrieves data from the gbdx endpoint
type CatalogSearch struct {
	response catalogSearchResponse
	request  catalogSearchRequest
	url      string
}

// EncodeJSON writes the results to an io.Writer as unescaped json. The client is responsible for setting any headers including the return code if required
func (c *CatalogSearch) EncodeJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	err := enc.Encode(&c.response)
	if err != nil {
		return err
	}
	return nil
}

// NewCatalogSearch takes a oauth2 client and a WKT string of a single polygon and gets results from the catalog search endpoint for all matching records (max 1000) inside the WKT
// I'd probably split this up in IRL and definitely clean up the error handling
func NewCatalogSearch(client *http.Client, wkt string) (c *CatalogSearch, err error) {
	c = new(CatalogSearch)
	c.request.WKT = wkt
	c.url = BasePath + SearchPath
	if err != nil {
		return nil, err
	}

	//request
	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(c.request)

	req, err := http.NewRequest("POST", c.url, payload)
	if err != nil {
		return nil, err
	}
	req.Header.Add("content-type", "application/json")

	//response
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	err = c.decodeJSON(resp.Body)
	if err != nil {
		return nil, err
	}

	return c, nil
}

// helper method to decode a stream into a response
// populates c.Response, discarding unwanted info
func (c *CatalogSearch) decodeJSON(r io.Reader) (err error) {
	err = json.NewDecoder(r).Decode(&c.response)
	return
}
