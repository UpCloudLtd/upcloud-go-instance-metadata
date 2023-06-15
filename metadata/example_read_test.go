package metadata_test

import (
	"log"

	"github.com/UpCloudLtd/upcloud-go-instance-metadata/metadata"
)

func Example_read() {
	m, _ := metadata.Read()
	log.Printf("I'm at '%s' and my instance ID is '%s'", m.Region, m.InstanceID)
}
