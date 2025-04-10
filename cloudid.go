package common

import (
	"net/http"
	"time"
)

const (
	// AWS is the identifier for AWS cloud provider
	AWS = "AWS"
	// GCP is the identifier for Google Cloud Platform
	GCP = "GCP"
	// Unknown is the identifier for unknown cloud provider
	Unknown = "Unknown"
)

func isAWS() bool {
	client := http.Client{Timeout: 500 * time.Millisecond}
	resp, err := client.Get("http://169.254.169.254/latest/meta-data/")
	return err == nil && resp.StatusCode == 200
}

func isGCP() bool {
	client := http.Client{Timeout: 500 * time.Millisecond}
	req, _ := http.NewRequest("GET", "http://metadata.google.internal/computeMetadata/v1/", nil)
	req.Header.Set("Metadata-Flavor", "Google")
	resp, err := client.Do(req)
	return err == nil && resp.StatusCode == 200
}

// CloudID detects the cloud provider
func CloudID() string {
	switch {
	case isAWS():
		return "AWS"
	case isGCP():
		return "GCP"
	default:
		return "Unknown"
	}
}
