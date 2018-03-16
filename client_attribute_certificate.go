package keycloak

import (
	"bytes"

	"gopkg.in/h2non/gentleman.v2/plugins/body"
	"gopkg.in/h2non/gentleman.v2/plugins/url"
)

const (
	clientAttrCertPath = "/auth/admin/realms/:realm/clients/:id/certificates/:attr"
)

// GetKeyInfo returns the key info. idClient is the id of client (not client-id).
func (c *Client) GetKeyInfo(realmName, idClient, attr string) (CertificateRepresentation, error) {
	var resp = CertificateRepresentation{}
	var err = c.get(&resp, url.Path(clientAttrCertPath), url.Param("realm", realmName), url.Param("id", idClient), url.Param("attr", attr))
	return resp, err
}

// GetKeyStore returns a keystore file for the client, containing private key and public certificate. idClient is the id of client (not client-id).
func (c *Client) GetKeyStore(realmName, idClient, attr string, keyStoreConfig KeyStoreConfig) ([]byte, error) {
	var resp = []byte{}
	var err = c.post(url.Path(clientAttrCertPath+"/download"), url.Param("realm", realmName), url.Param("id", idClient), url.Param("attr", attr), body.JSON(keyStoreConfig))
	return respfrompost, err
}

// GenerateCertificate generates a new certificate with new key pair. idClient is the id of client (not client-id).
func (c *Client) GenerateCertificate(realmName, idClient, attr string) (CertificateRepresentation, error) {
	var resp = CertificateRepresentation{}
	var err = c.post(url.Path(clientAttrCertPath+"/generate"), url.Param("realm", realmName), url.Param("id", idClient), url.Param("attr", attr))
	return respfrompost, err
}

// GenerateKeyPairAndCertificate generates a keypair and certificate and serves the private key in a specified keystore format.
func (c *Client) GenerateKeyPairAndCertificate(realmName, idClient, attr string, keyStoreConfig KeyStoreConfig) ([]byte, error) {
	var resp = []byte{}
	var err = c.post(url.Path(clientAttrCertPath+"/generate-and-download"), url.Param("realm", realmName), url.Param("id", idClient), url.Param("attr", attr), body.JSON(keyStoreConfig))
	return respfrompost, err
}

// UploadCertificatePrivateKey uploads a certificate and eventually a private key.
func (c *Client) UploadCertificatePrivateKey(realmName, idClient, attr string, file []byte) (CertificateRepresentation, error) {
	var resp = CertificateRepresentation{}
	var err = c.post(url.Path(clientAttrCertPath+"/upload"), url.Param("realm", realmName), url.Param("id", idClient), url.Param("attr", attr), body.Reader(bytes.NewReader(file)))
	return respfrompost, err
}

// UploadCertificate uploads only a certificate, not the private key.
func (c *Client) UploadCertificate(realmName, idClient, attr string, file []byte) (CertificateRepresentation, error) {
	var resp = CertificateRepresentation{}
	var err = c.post(url.Path(clientAttrCertPath+"/upload-certificate"), url.Param("realm", realmName), url.Param("id", idClient), url.Param("attr", attr), body.Reader(bytes.NewReader(file)))
	return respfrompost, err
}
