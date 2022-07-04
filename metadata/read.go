package metadata

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const MetadataURL = "http://169.254.169.254/metadata/v1.json"

func Read() (*InstanceMetadata, error) {
	resp, err := http.Get(MetadataURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var res InstanceMetadata
	return &res, json.Unmarshal(body, &res)
}
