package tls

import (
	"crypto/tls"
	"errors"
	"net/http"
	"strings"
	"time"
)

type Subject struct {
	CommonName       string   `json:"common_name"`
	Country          []string `json:"country,omitempty"`
	Organization     []string `json:"organization,omitempty"`
	OrganizationUnit []string `json:"organization_unit,omitempty"`
}

type Issuer struct {
	CommonName       string   `json:"common_name"`
	Country          []string `json:"country,omitempty"`
	Organization     []string `json:"organization,omitempty"`
	OrganizationUnit []string `json:"organization_unit,omitempty"`
}

type CertInfo struct {
	IsCA                  bool      `json:"is_ca,omitempty"`
	Version               int       `json:"version,omitempty"`
	DNSNames              []string  `json:"dns_names,omitempty"`
	Subject               *Subject  `json:"subject,omitempty"`
	Issuer                *Issuer   `json:"issuer,omitempty"`
	NotBefore             time.Time `json:"not_before,omitempty"`
	NotAfter              time.Time `json:"not_after,omitempty"`
	IssuingCertificateURL []string  `json:"issuing_certificate_url,omitempty"`
}

func GetCertInfo(url string) (*CertInfo, error) {
	if !strings.HasPrefix(url, "https://") {
		return nil, errors.New("url should start with https://")
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   15 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	rci := resp.TLS.PeerCertificates[0]

	ci := &CertInfo{
		IsCA:     rci.IsCA,
		Version:  rci.Version,
		DNSNames: rci.DNSNames,
		Subject: &Subject{
			CommonName:       rci.Subject.CommonName,
			Country:          rci.Subject.Country,
			Organization:     rci.Subject.Organization,
			OrganizationUnit: rci.Subject.OrganizationalUnit,
		},
		Issuer: &Issuer{
			CommonName:       rci.Issuer.CommonName,
			Country:          rci.Issuer.Country,
			Organization:     rci.Issuer.Organization,
			OrganizationUnit: rci.Issuer.OrganizationalUnit,
		},
		NotBefore:             rci.NotBefore,
		NotAfter:              rci.NotAfter,
		IssuingCertificateURL: rci.IssuingCertificateURL,
	}
	return ci, nil
}
