package metadata

import (
	"encoding/json"
	"io"
	"net/http"
)

const MetadataURL = "http://169.254.169.254/metadata/v1.json"

// Read fetches and unmarshals metadata from metadata service
func Read() (*InstanceMetadata, error) {
	resp, err := http.Get(MetadataURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var res InstanceMetadata
	return &res, json.Unmarshal(body, &res)
}
