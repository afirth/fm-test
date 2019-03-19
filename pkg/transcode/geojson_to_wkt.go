package transcode

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/paulmach/orb/encoding/wkt"
	"github.com/paulmach/orb/geojson"
)

// GeoJSON2WKT takes a geojson feature collection and transcodes it into a WKT string
// Feature collection must contain exactly one feature (as the sample aoi)
func GeoJSON2WKT(r io.Reader) (out string, err error) {

	// get bytes
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return "", fmt.Errorf("error transcoding geojson to wkt: %v", err)
	}

	//Unmarshal the json into an orb fc
	fc, err := geojson.UnmarshalFeatureCollection(b)
	if err != nil {
		return "", fmt.Errorf("error unmarshalling geojson: %v", err)
	}

	// Test the fc contains one polygon
	// code adapted from https://github.com/paulmach/orb/blob/master/geojson/feature_collection_test.go
	if fc.Type != "FeatureCollection" {
		return "", fmt.Errorf("should have type of FeatureCollection, got %v", fc.Type)
	}
	if gl := len(fc.Features); gl != 1 {
		return "", fmt.Errorf("should have one feature, got %d", gl)
	}

	f := fc.Features[0]
	if gt := f.Geometry.GeoJSONType(); gt != "Polygon" {
		return "", fmt.Errorf("incorrect feature type, want a Polygon, got %v", gt)
	}

	//Marshal the fc into WKT
	return wkt.MarshalString(f.Geometry), nil
}
