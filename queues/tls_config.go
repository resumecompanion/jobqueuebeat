package queues

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/elastic/beats/libbeat/logp"
	"io/ioutil"
)

// SetupTLSConfig setup tls config for mysql
func SetupTLSConfig(caPath string, certPath string, keyPath string) tls.Config {
	rootCertPool := x509.NewCertPool()
	pem, err := ioutil.ReadFile(caPath)

	if err != nil {
		logp.Warn(err.Error())
	}

	rootCertPool.AppendCertsFromPEM(pem)

	clientCert := make([]tls.Certificate, 0, 1)
	certs, errKey := tls.LoadX509KeyPair(certPath, keyPath)

	if errKey != nil {
		logp.Warn(errKey.Error())
	}

	clientCert = append(clientCert, certs)

	return tls.Config{
		RootCAs:            rootCertPool,
		InsecureSkipVerify: true,
		Certificates:       clientCert,
	}
}
