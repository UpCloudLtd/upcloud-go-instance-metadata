package main

import (
	"log"

	"github.com/UpCloudLtd/upcloud-go-instance-metadata/metadata"
)

func main() {
	m, _ := metadata.Read()
	log.Printf("I'm at '%s' and my instance ID is '%s'", m.Region, m.InstanceID)
}
