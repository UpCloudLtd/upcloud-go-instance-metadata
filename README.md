# UpCloud instance metadata Go library

[![Go Reference](https://pkg.go.dev/badge/github.com/UpCloudLtd/upcloud-go-instance-metadata.svg)](https://pkg.go.dev/github.com/UpCloudLtd/upcloud-go-instance-metadata)

Go library for reading UpCloud server instance metadata from `http://169.254.169.254/metadata/v1.json`.

## Installation

You can use the following command to retrieve the this library:

```sh
go get github.com/UpCloudLtd/upcloud-go-instance-metadata
```

## Usage

To fetch an unmarshal instance metadata, see [example](metadata/example_read_test.go).
