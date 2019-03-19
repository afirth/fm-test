package gbdx

import (
	"bytes"
	"encoding/json"
	"os"
	"strings"
	"testing"

	"golang.org/x/oauth2"
)

var (
	// example response from search api
	raw = []byte(`{"stats": {"recordsReturned": 1000, "totalRecords": 1000, "typeCounts": {"DigitalGlobeAcquisition": 298, "MDAProduct": 35, "SENTINEL1": 170, "SENTINEL2": 25, "GBDXCatalogRecord": 1000, "Landsat8": 223, "QB02": 76, "WV03_SWIR": 9, "WV02": 149, "WV01": 82, "WV03_VNIR": 63, "MDAAcquisition": 44, "1BProduct": 23, "IKONOSAcquisition": 82, "ESAProduct": 195, "IDAHOImage": 100, "RADARSAT2": 79, "DigitalGlobeProduct": 123, "SGFProduct": 35, "Acquisition": 647, "LandsatAcquisition": 223, "IKONOS": 82, "GE01": 42}}, "results": [{"identifier": "626230_001", "type": ["GBDXCatalogRecord", "Acquisition", "MDAAcquisition", "RADARSAT2"], "properties": {"beamMode": "Extra Fine", "imageId": "626230", "imgCycle": 155, "catalogID": "626230", "beams": "XF0W3", "archiveId": "420799", "platformName": "RADARSAT-2", "archiveFac": "PASS", "vendor": "MacDonald, Dettwiler and Associates Ltd.", "timestamp": "2018-03-26T14:07:23.167Z", "polarizations": ["HH"], "passDirection": "Descending", "incidenceAngleNearRange": 38.33, "sceneNo": 1, "footprintWkt": "MULTIPOLYGON(((-122.49986 38.50951, -121.21348 38.33235, -121.45742 37.24294, -122.72489 37.42102, -122.49986 38.50951)))", "imgRelOrNo": 264.390650969, "sensorPlatformName": "RADARSAT-2", "lookDirection": "Right"}}, {"identifier": "616502_001", "type": ["GBDXCatalogRecord", "Acquisition", "MDAAcquisition", "RADARSAT2"], "properties": {"beamMode": "Extra Fine", "imageId": "616502", "imgCycle": 153, "catalogID": "616502", "beams": "XF0W3", "archiveId": "417229", "platformName": "RADARSAT-2", "archiveFac": "PASS", "vendor": "MacDonald, Dettwiler and Associates Ltd.", "timestamp": "2018-02-06T14:07:25.913Z", "polarizations": ["HH"], "passDirection": "Descending", "incidenceAngleNearRange": 38.33, "sceneNo": 1, "footprintWkt": "MULTIPOLYGON(((-122.50144 38.50875, -121.21506 38.33158, -121.45898 37.24219, -122.72643 37.42027, -122.50144 38.50875)))", "imgRelOrNo": 264.391105183, "sensorPlatformName": "RADARSAT-2", "lookDirection": "Right"}}]}`)

	// marshalled output of response
	expected = `{"Results":[{"identifier":"626230_001","properties":{"platformName":"RADARSAT-2","timestamp":"2018-03-26T14:07:23.167Z","catalogID":"626230"}},{"identifier":"616502_001","properties":{"platformName":"RADARSAT-2","timestamp":"2018-02-06T14:07:25.913Z","catalogID":"616502"}}]}`
)

// just some sanity checks, not going to table up everything for a demo

// use marshal and unmarshal directly
func TestUnmarshalMarshal(t *testing.T) {
	c := loadRaw(raw, t)

	out, err := json.Marshal(c.response)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if g := string(out); g != expected {
		t.Errorf("output failure. Expected: %s\nGot: %s\n", expected, g)
	}
}

// calls the gbdx api for the sample WKT (denver)
// checks 3 < results < 1001 (currently the max)
func TestNewCatalogSearch(t *testing.T) {
	skipUnlessE2E(t)
	wkt := `POLYGON ((-122.41189956665039 37.59415685597818, -122.41189956665039 37.64460175855099, -122.34529495239259 37.64460175855099, -122.34529495239259 37.59415685597818, -122.41189956665039 37.59415685597818))`

	// Get an authed client
	client, err := NewClient(oauth2.NoContext, os.Getenv("USERNAME"), os.Getenv("PASSWORD"))
	if err != nil {
		t.Errorf("unable to create oauth2 http client (are USERNAME and PASSWORD in env?): %+v", err)
	}

	// exec a search
	c, err := NewCatalogSearch(client, wkt)
	if err != nil {
		t.Errorf("unable to execute search: %v", err)
	}

	// basic sanity check of results
	if l := len(c.response.Results); l <= 2 {
		t.Errorf("Expected more than 2 results, got %d", l)
	}
	if l := len(c.response.Results); l > 1000 {
		t.Errorf("Expected no more than 1000 results, got %d", l)
	}
}

func TestEncodeJSON(t *testing.T) {
	// populate a search
	c := loadRaw(raw, t)

	// get results as json
	b := new(bytes.Buffer)
	err := c.EncodeJSON(b)
	if err != nil {
		t.Errorf("error: %v", err)
	}

	if g := strings.TrimSpace(b.String()); g != expected {
		t.Errorf("output failure. Expected: %s\nGot: %s\n", expected, g)
	}
}

// helper returns a catalog search with results
// essentially mocking NewCatalogSearch
// very whitebox
func loadRaw(raw []byte, t *testing.T) (c *CatalogSearch) {
	c = &CatalogSearch{}
	r := bytes.NewReader(raw)
	err := c.decodeJSON(r)
	if err != nil {
		t.Fatalf("error loading raw data: %v", err)
	}
	return
}

// skip long tests (which require auth, in this case)
func skipUnlessE2E(t *testing.T) {
	if os.Getenv("GOTEST_E2E") == "" {
		t.Skip("Skipping e2e tests because GOTEST_E2E is empty or unset")
	}
}
