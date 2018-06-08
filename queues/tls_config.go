package queues

import (
  "crypto/tls"
  "io/ioutil"
  "crypto/x509"
  "github.com/elastic/beats/libbeat/logp"
)

func SetupTLSConfig(caPath string, certPath string, keyPath string) tls.Config {
// func setupTLSCOnfig(caPath string, certPath string, keyPath string) {
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
    RootCAs: rootCertPool,
    InsecureSkipVerify: true,
    Certificates: clientCert,
  }
}
