package core

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"
)

// CertConfig is the configuration for the certificate.
type CertConfig struct {
	CACertPath   string `mapstructure:"CACertPath"`
	NodeKeyPath  string `mapstructure:"NodeKeyPath"`
	NodeCertPath string `mapstructure:"NodeCertPath"`
}

// LoadTLS loads the tls config.
func LoadTLS(cfg *CertConfig, isServer bool) (*tls.Config, error) {
	caPem, err := os.ReadFile(cfg.CACertPath)
	if err != nil {
		return nil, err
	}
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(caPem) {
		return nil, fmt.Errorf("failed to append CA certificate")
	}
	nodeCert, err := tls.LoadX509KeyPair(cfg.NodeCertPath, cfg.NodeKeyPath)
	if err != nil {
		return nil, err
	}

	conf := &tls.Config{
		Certificates: []tls.Certificate{nodeCert},
	}
	if isServer {
		conf.ClientCAs = certPool
		conf.ClientAuth = tls.RequireAndVerifyClientCert
	} else {
		conf.RootCAs = certPool
	}

	return conf, nil
}
