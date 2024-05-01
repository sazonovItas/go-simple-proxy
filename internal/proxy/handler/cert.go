package proxy

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"errors"
	"fmt"
	"math/big"
	"net"
	"time"

	configproxy "github.com/sazonovItas/go-simple-proxy/internal/config/proxy"
)

type CertManager struct {
	// ca certificate
	ca *x509.Certificate

	// ca private key
	caPrivateKey *rsa.PrivateKey

	roots *x509.CertPool

	privateKey *rsa.PrivateKey

	validity     time.Duration
	keyID        []byte
	organization string
}

// TODO: Add config
// TODO: Rebuild and send some parts to pkg
func NewCertManager(cfg configproxy.ProxySecrets) (*CertManager, error) {
	const op = "internal.proxy.handler.cert.NewCertManager"

	tlsCert, err := tls.LoadX509KeyPair(cfg.Cert, cfg.Key)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	privateKey, ok := tlsCert.PrivateKey.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("%s: %w", op, errors.New("failed cast private key"))
	}

	ca, err := x509.ParseCertificate(tlsCert.Certificate[0])
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	roots := x509.NewCertPool()
	roots.AddCert(ca)

	privKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	publicKey := privKey.Public()

	pkixpub, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	h := sha1.New()
	_, err = h.Write(pkixpub)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	keyID := h.Sum(nil)

	return &CertManager{
		ca:           ca,
		caPrivateKey: privateKey,

		privateKey: privKey,
		keyID:      keyID,

		validity:     time.Hour,
		organization: "itas",
		roots:        roots,
	}, nil
}

func (cm *CertManager) GenFakeCert(hostname string) (*tls.Certificate, error) {
	const op = "internal.proxy.handler.cert.GenFakeCert"

	host, _, err := net.SplitHostPort(hostname)
	if err == nil {
		hostname = host
	}

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		return nil, err
	}

	tmpl := &x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			CommonName:   hostname,
			Organization: []string{cm.organization},
		},
		SubjectKeyId:          cm.keyID,
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		NotBefore:             time.Now().Add(-cm.validity),
		NotAfter:              time.Now().Add(cm.validity),
	}

	if ip := net.ParseIP(hostname); ip != nil {
		tmpl.IPAddresses = []net.IP{ip}
	} else {
		tmpl.DNSNames = []string{hostname}
	}

	raw, err := x509.CreateCertificate(
		rand.Reader,
		tmpl,
		cm.ca,
		cm.privateKey.Public(),
		cm.caPrivateKey,
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	x509c, err := x509.ParseCertificate(raw)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	cert := &tls.Certificate{
		Certificate: [][]byte{raw, cm.ca.Raw},
		PrivateKey:  cm.privateKey,
		Leaf:        x509c,
	}
	return cert, nil
}

func (cm *CertManager) NewTLSConfig(hostname string) *tls.Config {
	tlsConfig := &tls.Config{
		GetCertificate: func(clientHello *tls.ClientHelloInfo) (*tls.Certificate, error) {
			host := clientHello.ServerName
			if host == "" {
				host = hostname
			}

			return cm.GenFakeCert(host)
		},
		NextProtos: []string{"http/1.1"},
	}
	tlsConfig.InsecureSkipVerify = true
	return tlsConfig
}
