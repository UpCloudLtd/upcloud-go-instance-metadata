package metadata_test

import (
	"log"
	"os"

	"github.com/UpCloudLtd/upcloud-go-instance-metadata/metadata"
)

func Example_read() {
	m, err := metadata.Read()
	if err != nil {
		log.Printf("read failed: %s", err.Error())
		os.Exit(1)
	}
	if m != nil {
		log.Printf("I'm at '%s' and my instance ID is '%s'", m.Region, m.InstanceID)
	}
}
