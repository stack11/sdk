/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * The version of the OpenAPI document: v0.6.3-alpha.1
 * Contact: hi@ory.sh
 * Generated by: https://openapi-generator.tech
 */




#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct SubmitSelfServiceLoginFlow {
    /// Sending the anti-csrf token is only required for browser login flows.
    #[serde(rename = "csrf_token", skip_serializing_if = "Option::is_none")]
    pub csrf_token: Option<String>,
    /// Method should be set to \"password\" when logging in using the identifier and password strategy.
    #[serde(rename = "method", skip_serializing_if = "Option::is_none")]
    pub method: Option<String>,
    /// The user's password.
    #[serde(rename = "password", skip_serializing_if = "Option::is_none")]
    pub password: Option<String>,
    /// Identifier is the email or username of the user trying to log in.
    #[serde(rename = "password_identifier", skip_serializing_if = "Option::is_none")]
    pub password_identifier: Option<String>,
}

impl SubmitSelfServiceLoginFlow {
    pub fn new() -> SubmitSelfServiceLoginFlow {
        SubmitSelfServiceLoginFlow {
            csrf_token: None,
            method: None,
            password: None,
            password_identifier: None,
        }
    }
}


