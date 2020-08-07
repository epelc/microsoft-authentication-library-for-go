// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package msalgo

import "github.com/AzureAD/microsoft-authentication-library-for-go/src/internal/msalbase"

type credentialType int

const (
	clientSecret credentialType = iota
	clientAssertion
)

//ClientCredentialInterfacer is an interface representing a client secret/assertion.
// These are generated when you create your application on your Azure AD tenant.
//This is required to create a ConfidentialClientApplication instance.
type ClientCredentialInterfacer interface {
	GetCredentialType() msalbase.ClientCredentialType
	GetSecret() string
	GetAssertion() *msalbase.ClientAssertion
}

//CreateClientCredentialFromSecret returns a ClientCredentialInterfacer when given a client secret.
func CreateClientCredentialFromSecret(secret string) (ClientCredentialInterfacer, error) {
	return msalbase.CreateClientCredentialFromSecret(secret)
}

//CreateClientCredentialFromCertificate returns a ClientCredentialInterfacer when given a thumbprint and private key represnting the certificate.
func CreateClientCredentialFromCertificate(thumbprint string, key []byte) (ClientCredentialInterfacer, error) {
	return msalbase.CreateClientCredentialFromCertificate(thumbprint, key)
}

//CreateClientCredentialFromAssertion returns a ClientCredentialInterfacer when given an assertion JWT.
func CreateClientCredentialFromAssertion(assertion string) (ClientCredentialInterfacer, error) {
	return msalbase.CreateClientCredentialFromAssertion(assertion)
}
