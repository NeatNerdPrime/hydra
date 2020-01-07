// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Client Client represents an OAuth 2.0 Client.
// swagger:model Client
type Client struct {

	// AllowedCORSOrigins are one or more URLs (scheme://host[:port]) which are allowed to make CORS requests
	// to the /oauth/token endpoint. If this array is empty, the sever's CORS origin configuration (`CORS_ALLOWED_ORIGINS`)
	// will be used instead. If this array is set, the allowed origins are appended to the server's CORS origin configuration.
	// Be aware that environment variable `CORS_ENABLED` MUST be set to `true` for this to work.
	AllowedCORSOrigins []string `json:"allowed_cors_origins"`

	// Audience is a whitelist defining the audiences this client is allowed to request tokens for. An audience limits
	// the applicability of an OAuth 2.0 Access Token to, for example, certain API endpoints. The value is a list
	// of URLs. URLs MUST NOT contain whitespaces.
	Audience []string `json:"audience"`

	// Boolean value specifying whether the RP requires that a sid (session ID) Claim be included in the Logout
	// Token to identify the RP session with the OP when the backchannel_logout_uri is used.
	// If omitted, the default value is false.
	BackChannelLogoutSessionRequired bool `json:"backchannel_logout_session_required,omitempty"`

	// RP URL that will cause the RP to log itself out when sent a Logout Token by the OP.
	BackChannelLogoutURI string `json:"backchannel_logout_uri,omitempty"`

	// ClientID  is the id for this client.
	ClientID string `json:"client_id,omitempty"`

	// ClientURI is an URL string of a web page providing information about the client.
	// If present, the server SHOULD display this URL to the end-user in
	// a clickable fashion.
	ClientURI string `json:"client_uri,omitempty"`

	// Contacts is a array of strings representing ways to contact people responsible
	// for this client, typically email addresses.
	Contacts []string `json:"contacts"`

	// CreatedAt returns the timestamp of the client's creation.
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// Boolean value specifying whether the RP requires that iss (issuer) and sid (session ID) query parameters be
	// included to identify the RP session with the OP when the frontchannel_logout_uri is used.
	// If omitted, the default value is false.
	FrontChannelLogoutSessionRequired bool `json:"frontchannel_logout_session_required,omitempty"`

	// RP URL that will cause the RP to log itself out when rendered in an iframe by the OP. An iss (issuer) query
	// parameter and a sid (session ID) query parameter MAY be included by the OP to enable the RP to validate the
	// request and to determine which of the potentially multiple sessions is to be logged out; if either is
	// included, both MUST be.
	FrontChannelLogoutURI string `json:"frontchannel_logout_uri,omitempty"`

	// GrantTypes is an array of grant types the client is allowed to use.
	// Pattern: client_credentials|authorization_code|implicit|refresh_token
	GrantTypes []string `json:"grant_types"`

	// URL for the Client's JSON Web Key Set [JWK] document. If the Client signs requests to the Server, it contains
	// the signing key(s) the Server uses to validate signatures from the Client. The JWK Set MAY also contain the
	// Client's encryption keys(s), which are used by the Server to encrypt responses to the Client. When both signing
	// and encryption keys are made available, a use (Key Use) parameter value is REQUIRED for all keys in the referenced
	// JWK Set to indicate each key's intended usage. Although some algorithms allow the same key to be used for both
	// signatures and encryption, doing so is NOT RECOMMENDED, as it is less secure. The JWK x5c parameter MAY be used
	// to provide X.509 representations of keys provided. When used, the bare key values MUST still be present and MUST
	// match those in the certificate.
	JSONWebKeysURI string `json:"jwks_uri,omitempty"`

	// LogoURI is an URL string that references a logo for the client.
	LogoURI string `json:"logo_uri,omitempty"`

	// Name is the human-readable string name of the client to be presented to the
	// end-user during authorization.
	Name string `json:"client_name,omitempty"`

	// Owner is a string identifying the owner of the OAuth 2.0 Client.
	Owner string `json:"owner,omitempty"`

	// PolicyURI is a URL string that points to a human-readable privacy policy document
	// that describes how the deployment organization collects, uses,
	// retains, and discloses personal data.
	PolicyURI string `json:"policy_uri,omitempty"`

	// Array of URLs supplied by the RP to which it MAY request that the End-User's User Agent be redirected using the
	// post_logout_redirect_uri parameter after a logout has been performed.
	PostLogoutRedirectURIs []string `json:"post_logout_redirect_uris"`

	// RedirectURIs is an array of allowed redirect urls for the client, for example http://mydomain/oauth/callback .
	RedirectURIs []string `json:"redirect_uris"`

	// JWS [JWS] alg algorithm [JWA] that MUST be used for signing Request Objects sent to the OP. All Request Objects
	// from this Client MUST be rejected, if not signed with this algorithm.
	RequestObjectSigningAlgorithm string `json:"request_object_signing_alg,omitempty"`

	// Array of request_uri values that are pre-registered by the RP for use at the OP. Servers MAY cache the
	// contents of the files referenced by these URIs and not retrieve them at the time they are used in a request.
	// OPs can require that request_uri values used be pre-registered with the require_request_uri_registration
	// discovery parameter.
	RequestURIs []string `json:"request_uris"`

	// ResponseTypes is an array of the OAuth 2.0 response type strings that the client can
	// use at the authorization endpoint.
	// Pattern: id_token|code|token
	ResponseTypes []string `json:"response_types"`

	// Scope is a string containing a space-separated list of scope values (as
	// described in Section 3.3 of OAuth 2.0 [RFC6749]) that the client
	// can use when requesting access tokens.
	// Pattern: ([a-zA-Z0-9\.\*]+\s?)+
	Scope string `json:"scope,omitempty"`

	// Secret is the client's secret. The secret will be included in the create request as cleartext, and then
	// never again. The secret is stored using BCrypt so it is impossible to recover it. Tell your users
	// that they need to write the secret down as it will not be made available again.
	Secret string `json:"client_secret,omitempty"`

	// SecretExpiresAt is an integer holding the time at which the client
	// secret will expire or 0 if it will not expire. The time is
	// represented as the number of seconds from 1970-01-01T00:00:00Z as
	// measured in UTC until the date/time of expiration.
	//
	// This feature is currently not supported and it's value will always
	// be set to 0.
	SecretExpiresAt int64 `json:"client_secret_expires_at,omitempty"`

	// URL using the https scheme to be used in calculating Pseudonymous Identifiers by the OP. The URL references a
	// file with a single JSON array of redirect_uri values.
	SectorIdentifierURI string `json:"sector_identifier_uri,omitempty"`

	// SubjectType requested for responses to this Client. The subject_types_supported Discovery parameter contains a
	// list of the supported subject_type values for this server. Valid types include `pairwise` and `public`.
	SubjectType string `json:"subject_type,omitempty"`

	// TermsOfServiceURI is a URL string that points to a human-readable terms of service
	// document for the client that describes a contractual relationship
	// between the end-user and the client that the end-user accepts when
	// authorizing the client.
	TermsOfServiceURI string `json:"tos_uri,omitempty"`

	// Requested Client Authentication method for the Token Endpoint. The options are client_secret_post,
	// client_secret_basic, private_key_jwt, and none.
	TokenEndpointAuthMethod string `json:"token_endpoint_auth_method,omitempty"`

	// UpdatedAt returns the timestamp of the last update.
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`

	// JWS alg algorithm [JWA] REQUIRED for signing UserInfo Responses. If this is specified, the response will be JWT
	// [JWT] serialized, and signed using JWS. The default, if omitted, is for the UserInfo Response to return the Claims
	// as a UTF-8 encoded JSON object using the application/json content-type.
	UserinfoSignedResponseAlg string `json:"userinfo_signed_response_alg,omitempty"`

	// jwks
	Jwks *SwaggerJSONWebKeySet `json:"jwks,omitempty"`

	// metadata
	Metadata RawMessage `json:"metadata,omitempty"`
}

// Validate validates this client
func (m *Client) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGrantTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResponseTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScope(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJwks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Client) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Client) validateGrantTypes(formats strfmt.Registry) error {

	if swag.IsZero(m.GrantTypes) { // not required
		return nil
	}

	return nil
}

func (m *Client) validateResponseTypes(formats strfmt.Registry) error {

	if swag.IsZero(m.ResponseTypes) { // not required
		return nil
	}

	return nil
}

func (m *Client) validateScope(formats strfmt.Registry) error {

	if swag.IsZero(m.Scope) { // not required
		return nil
	}

	if err := validate.Pattern("scope", "body", string(m.Scope), `([a-zA-Z0-9\.\*]+\s?)+`); err != nil {
		return err
	}

	return nil
}

func (m *Client) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Client) validateJwks(formats strfmt.Registry) error {

	if swag.IsZero(m.Jwks) { // not required
		return nil
	}

	if m.Jwks != nil {
		if err := m.Jwks.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("jwks")
			}
			return err
		}
	}

	return nil
}

func (m *Client) validateMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if err := m.Metadata.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("metadata")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Client) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Client) UnmarshalBinary(b []byte) error {
	var res Client
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
