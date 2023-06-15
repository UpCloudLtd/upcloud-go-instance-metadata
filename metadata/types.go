package metadata

type InstanceMetadata struct {
	CloudName   string   `json:"cloud_name"`
	InstanceID  string   `json:"instance_id"`
	Hostname    string   `json:"hostname"`
	Platform    string   `json:"platform"`
	Subplatform string   `json:"subplatform"`
	PublicKeys  []string `json:"public_keys"`
	Region      string   `json:"region"`
	Network     Network  `json:"network"`
	Storage     Storage  `json:"storage"`
	Tags        []string `json:"tags"`
	UserData    string   `json:"user_data"`
	VendorData  string   `json:"vendor_data"`
}

type Network struct {
	Interfaces []NetworkInterface `json:"interfaces"`
	DNS        []string           `json:"dns"`
}

type NetworkInterface struct {
	Index       int         `json:"index"`
	IPAddresses []IPAddress `json:"ip_addresses"`
	MAC         string      `json:"mac"`
	NetworkID   string      `json:"network_id"`
	Type        string      `json:"type"`
}

type IPAddress struct {
	Address  string   `json:"address"`
	DHCP     bool     `json:"dhcp"`
	DNS      []string `json:"dns"`
	Family   string   `json:"family"`
	Floating bool     `json:"floating"`
	Gateway  string   `json:"gateway"`
	Network  string   `json:"network"`
}

type Storage struct {
	Disks []StorageDisk `json:"disks"`
}

type StorageDisk struct {
	ID     string `json:"id"`
	Serial string `json:"serial"`
	Size   int    `json:"size"`
	Type   string `json:"type"`
	Tier   string `json:"tier"`
}
