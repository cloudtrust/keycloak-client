package api

import (
	"bytes"

	"github.com/cloudtrust/keycloak-client/v2"
	"gopkg.in/h2non/gentleman.v2/plugins/body"
	"gopkg.in/h2non/gentleman.v2/plugins/url"
)

const (
	// API Keycloak out-of-the-box
	kcClientAttrCertPath = "/auth/admin/realms/:realm/clients/:id/certificates/:attr"
)

// GetKeyInfo returns the key info. idClient is the id of client (not client-id).
func (c *Client) GetKeyInfo(accessToken string, realmName, idClient, attr string) (keycloak.CertificateRepresentation, error) {
	var resp = keycloak.CertificateRepresentation{}
	var err = c.forRealm(realmName).
		get(accessToken, &resp, url.Path(kcClientAttrCertPath), url.Param("realm", realmName), url.Param("id", idClient), url.Param("attr", attr))
	return resp, err
}

// GetKeyStore returns a keystore file for the client, containing private key and public certificate. idClient is the id of client (not client-id).
func (c *Client) GetKeyStore(accessToken string, realmName, idClient, attr string, keyStoreConfig keycloak.KeyStoreConfig) ([]byte, error) {
	var resp = []byte{}
	_, err := c.forRealm(realmName).
		post(accessToken, &resp, url.Path(kcClientAttrCertPath+"/download"), url.Param("realm", realmName), url.Param("id", idClient), url.Param("attr", attr), body.JSON(keyStoreConfig))
	return resp, err
}

// GenerateCertificate generates a new certificate with new key pair. idClient is the id of client (not client-id).
func (c *Client) GenerateCertificate(accessToken string, realmName, idClient, attr string) (keycloak.CertificateRepresentation, error) {
	var resp = keycloak.CertificateRepresentation{}
	_, err := c.forRealm(realmName).
		post(accessToken, &resp, url.Path(kcClientAttrCertPath+"/generate"), url.Param("realm", realmName), url.Param("id", idClient), url.Param("attr", attr))
	return resp, err
}

// GenerateKeyPairAndCertificate generates a keypair and certificate and serves the private key in a specified keystore format.
func (c *Client) GenerateKeyPairAndCertificate(accessToken string, realmName, idClient, attr string, keyStoreConfig keycloak.KeyStoreConfig) ([]byte, error) {
	var resp = []byte{}
	_, err := c.forRealm(realmName).
		post(accessToken, &resp, url.Path(kcClientAttrCertPath+"/generate-and-download"), url.Param("realm", realmName), url.Param("id", idClient), url.Param("attr", attr), body.JSON(keyStoreConfig))
	return resp, err
}

// UploadCertificatePrivateKey uploads a certificate and eventually a private key.
func (c *Client) UploadCertificatePrivateKey(accessToken string, realmName, idClient, attr string, file []byte) (keycloak.CertificateRepresentation, error) {
	var resp = keycloak.CertificateRepresentation{}
	_, err := c.forRealm(realmName).
		post(accessToken, &resp, url.Path(kcClientAttrCertPath+"/upload"), url.Param("realm", realmName), url.Param("id", idClient), url.Param("attr", attr), body.Reader(bytes.NewReader(file)))
	return resp, err
}

// UploadCertificate uploads only a certificate, not the private key.
func (c *Client) UploadCertificate(accessToken string, realmName, idClient, attr string, file []byte) (keycloak.CertificateRepresentation, error) {
	var resp = keycloak.CertificateRepresentation{}
	_, err := c.forRealm(realmName).
		post(accessToken, &resp, url.Path(kcClientAttrCertPath+"/upload-certificate"), url.Param("realm", realmName), url.Param("id", idClient), url.Param("attr", attr), body.Reader(bytes.NewReader(file)))
	return resp, err
}
