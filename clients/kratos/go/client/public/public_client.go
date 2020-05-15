// Code generated by go-swagger; DO NOT EDIT.

package public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new public API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for public API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CompleteSelfServiceBrowserSettingsOIDCSettingsFlow(params *CompleteSelfServiceBrowserSettingsOIDCSettingsFlowParams) error

	CompleteSelfServiceBrowserSettingsPasswordStrategyFlow(params *CompleteSelfServiceBrowserSettingsPasswordStrategyFlowParams) error

	CompleteSelfServiceBrowserSettingsProfileStrategyFlow(params *CompleteSelfServiceBrowserSettingsProfileStrategyFlowParams) error

	CompleteSelfServiceBrowserVerificationFlow(params *CompleteSelfServiceBrowserVerificationFlowParams) error

	InitializeSelfServiceBrowserLoginFlow(params *InitializeSelfServiceBrowserLoginFlowParams) error

	InitializeSelfServiceBrowserLogoutFlow(params *InitializeSelfServiceBrowserLogoutFlowParams) error

	InitializeSelfServiceBrowserRegistrationFlow(params *InitializeSelfServiceBrowserRegistrationFlowParams) error

	InitializeSelfServiceBrowserVerificationFlow(params *InitializeSelfServiceBrowserVerificationFlowParams) error

	InitializeSelfServiceSettingsFlow(params *InitializeSelfServiceSettingsFlowParams) error

	SelfServiceBrowserVerify(params *SelfServiceBrowserVerifyParams) error

	Whoami(params *WhoamiParams) (*WhoamiOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CompleteSelfServiceBrowserSettingsOIDCSettingsFlow completes the browser based settings flow for the open ID connect strategy

  This endpoint completes a browser-based settings flow. This is usually achieved by POSTing data to this
endpoint.

> This endpoint is NOT INTENDED for API clients and only works with browsers (Chrome, Firefox, ...) and HTML Forms.

More information can be found at [ORY Kratos User Settings & Profile Management Documentation](../self-service/flows/user-settings).
*/
func (a *Client) CompleteSelfServiceBrowserSettingsOIDCSettingsFlow(params *CompleteSelfServiceBrowserSettingsOIDCSettingsFlowParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCompleteSelfServiceBrowserSettingsOIDCSettingsFlowParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "completeSelfServiceBrowserSettingsOIDCSettingsFlow",
		Method:             "POST",
		PathPattern:        "/self-service/browser/flows/registration/strategies/oidc/settings/connections",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CompleteSelfServiceBrowserSettingsOIDCSettingsFlowReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil
}

/*
  CompleteSelfServiceBrowserSettingsPasswordStrategyFlow completes the browser based settings flow for the password strategy

  This endpoint completes a browser-based settings flow. This is usually achieved by POSTing data to this
endpoint.

> This endpoint is NOT INTENDED for API clients and only works with browsers (Chrome, Firefox, ...) and HTML Forms.

More information can be found at [ORY Kratos User Settings & Profile Management Documentation](../self-service/flows/user-settings).
*/
func (a *Client) CompleteSelfServiceBrowserSettingsPasswordStrategyFlow(params *CompleteSelfServiceBrowserSettingsPasswordStrategyFlowParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCompleteSelfServiceBrowserSettingsPasswordStrategyFlowParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "completeSelfServiceBrowserSettingsPasswordStrategyFlow",
		Method:             "POST",
		PathPattern:        "/self-service/browser/flows/settings/strategies/password",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CompleteSelfServiceBrowserSettingsPasswordStrategyFlowReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil
}

/*
  CompleteSelfServiceBrowserSettingsProfileStrategyFlow completes the browser based settings flow for profile data

  This endpoint completes a browser-based settings flow. This is usually achieved by POSTing data to this
endpoint.

If the provided profile data is valid against the Identity's Traits JSON Schema, the data will be updated and
the browser redirected to `url.settings_ui` for further steps.

> This endpoint is NOT INTENDED for API clients and only works with browsers (Chrome, Firefox, ...) and HTML Forms.

More information can be found at [ORY Kratos User Settings & Profile Management Documentation](../self-service/flows/user-settings).
*/
func (a *Client) CompleteSelfServiceBrowserSettingsProfileStrategyFlow(params *CompleteSelfServiceBrowserSettingsProfileStrategyFlowParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCompleteSelfServiceBrowserSettingsProfileStrategyFlowParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "completeSelfServiceBrowserSettingsProfileStrategyFlow",
		Method:             "POST",
		PathPattern:        "/self-service/browser/flows/settings/strategies/profile",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CompleteSelfServiceBrowserSettingsProfileStrategyFlowReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil
}

/*
  CompleteSelfServiceBrowserVerificationFlow completes the browser based verification flows

  This endpoint completes a browser-based verification flow. This is usually achieved by POSTing data to this
endpoint.

If the provided data is valid against the Identity's Traits JSON Schema, the data will be updated and
the browser redirected to `url.settings_ui` for further steps.

> This endpoint is NOT INTENDED for API clients and only works with browsers (Chrome, Firefox, ...) and HTML Forms.

More information can be found at [ORY Kratos Email and Phone Verification Documentation](https://www.ory.sh/docs/kratos/selfservice/flows/verify-email-account-activation).
*/
func (a *Client) CompleteSelfServiceBrowserVerificationFlow(params *CompleteSelfServiceBrowserVerificationFlowParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCompleteSelfServiceBrowserVerificationFlowParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "completeSelfServiceBrowserVerificationFlow",
		Method:             "POST",
		PathPattern:        "/self-service/browser/flows/verification/{via}/complete",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CompleteSelfServiceBrowserVerificationFlowReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil
}

/*
  InitializeSelfServiceBrowserLoginFlow initializes browser based login user flow

  This endpoint initializes a browser-based user login flow. Once initialized, the browser will be redirected to
`urls.login_ui` with the request ID set as a query parameter. If a valid user session exists already, the browser will be
redirected to `urls.default_redirect_url`.

> This endpoint is NOT INTENDED for API clients and only works
with browsers (Chrome, Firefox, ...).

More information can be found at [ORY Kratos User Login and User Registration Documentation](https://www.ory.sh/docs/next/kratos/self-service/flows/user-login-user-registration).
*/
func (a *Client) InitializeSelfServiceBrowserLoginFlow(params *InitializeSelfServiceBrowserLoginFlowParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInitializeSelfServiceBrowserLoginFlowParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "initializeSelfServiceBrowserLoginFlow",
		Method:             "GET",
		PathPattern:        "/self-service/browser/flows/login",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &InitializeSelfServiceBrowserLoginFlowReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil
}

/*
  InitializeSelfServiceBrowserLogoutFlow initializes browser based logout user flow

  This endpoint initializes a logout flow.

> This endpoint is NOT INTENDED for API clients and only works
with browsers (Chrome, Firefox, ...).

On successful logout, the browser will be redirected (HTTP 302 Found) to `urls.default_return_to`.

More information can be found at [ORY Kratos User Logout Documentation](https://www.ory.sh/docs/next/kratos/self-service/flows/user-logout).
*/
func (a *Client) InitializeSelfServiceBrowserLogoutFlow(params *InitializeSelfServiceBrowserLogoutFlowParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInitializeSelfServiceBrowserLogoutFlowParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "initializeSelfServiceBrowserLogoutFlow",
		Method:             "GET",
		PathPattern:        "/self-service/browser/flows/logout",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &InitializeSelfServiceBrowserLogoutFlowReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil
}

/*
  InitializeSelfServiceBrowserRegistrationFlow initializes browser based registration user flow

  This endpoint initializes a browser-based user registration flow. Once initialized, the browser will be redirected to
`urls.registration_ui` with the request ID set as a query parameter. If a valid user session exists already, the browser will be
redirected to `urls.default_redirect_url`.

> This endpoint is NOT INTENDED for API clients and only works
with browsers (Chrome, Firefox, ...).

More information can be found at [ORY Kratos User Login and User Registration Documentation](https://www.ory.sh/docs/next/kratos/self-service/flows/user-login-user-registration).
*/
func (a *Client) InitializeSelfServiceBrowserRegistrationFlow(params *InitializeSelfServiceBrowserRegistrationFlowParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInitializeSelfServiceBrowserRegistrationFlowParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "initializeSelfServiceBrowserRegistrationFlow",
		Method:             "GET",
		PathPattern:        "/self-service/browser/flows/registration",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &InitializeSelfServiceBrowserRegistrationFlowReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil
}

/*
  InitializeSelfServiceBrowserVerificationFlow initializes browser based verification flow

  This endpoint initializes a browser-based verification flow. Once initialized, the browser will be redirected to
`urls.settings_ui` with the request ID set as a query parameter. If no valid user session exists, a login
flow will be initialized.

> This endpoint is NOT INTENDED for API clients and only works
with browsers (Chrome, Firefox, ...).

More information can be found at [ORY Kratos Email and Phone Verification Documentation](https://www.ory.sh/docs/kratos/selfservice/flows/verify-email-account-activation).
*/
func (a *Client) InitializeSelfServiceBrowserVerificationFlow(params *InitializeSelfServiceBrowserVerificationFlowParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInitializeSelfServiceBrowserVerificationFlowParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "initializeSelfServiceBrowserVerificationFlow",
		Method:             "GET",
		PathPattern:        "/self-service/browser/flows/verification/init/{via}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &InitializeSelfServiceBrowserVerificationFlowReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil
}

/*
  InitializeSelfServiceSettingsFlow initializes browser based settings flow

  This endpoint initializes a browser-based settings flow. Once initialized, the browser will be redirected to
`urls.settings_ui` with the request ID set as a query parameter. If no valid user session exists, a login
flow will be initialized.

> This endpoint is NOT INTENDED for API clients and only works
with browsers (Chrome, Firefox, ...).

More information can be found at [ORY Kratos User Settings & Profile Management Documentation](../self-service/flows/user-settings).
*/
func (a *Client) InitializeSelfServiceSettingsFlow(params *InitializeSelfServiceSettingsFlowParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInitializeSelfServiceSettingsFlowParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "initializeSelfServiceSettingsFlow",
		Method:             "GET",
		PathPattern:        "/self-service/browser/flows/settings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &InitializeSelfServiceSettingsFlowReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil
}

/*
  SelfServiceBrowserVerify completes the browser based verification flows

  This endpoint completes a browser-based verification flow.

> This endpoint is NOT INTENDED for API clients and only works with browsers (Chrome, Firefox, ...) and HTML Forms.

More information can be found at [ORY Kratos Email and Phone Verification Documentation](https://www.ory.sh/docs/kratos/selfservice/flows/verify-email-account-activation).
*/
func (a *Client) SelfServiceBrowserVerify(params *SelfServiceBrowserVerifyParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSelfServiceBrowserVerifyParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "selfServiceBrowserVerify",
		Method:             "GET",
		PathPattern:        "/self-service/browser/flows/verification/{via}/confirm/{code}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SelfServiceBrowserVerifyReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil
}

/*
  Whoami checks who the current HTTP session belongs to

  Uses the HTTP Headers in the GET request to determine (e.g. by using checking the cookies) who is authenticated.
Returns a session object or 401 if the credentials are invalid or no credentials were sent.

This endpoint is useful for reverse proxies and API Gateways.
*/
func (a *Client) Whoami(params *WhoamiParams) (*WhoamiOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWhoamiParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "whoami",
		Method:             "GET",
		PathPattern:        "/sessions/whoami",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &WhoamiReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*WhoamiOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for whoami: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
