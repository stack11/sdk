/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v0.0.1-alpha.17
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 */



#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
#[serde(tag = "method")]
pub enum SubmitSelfServiceRegistrationFlowBody {
    #[serde(rename="oidc")]
    SubmitSelfServiceRegistrationFlowWithOidcMethodBody {
        /// The CSRF Token
        #[serde(rename = "csrf_token", skip_serializing_if = "Option::is_none")]
        csrf_token: Option<String>,
        /// The provider to register with
        #[serde(rename = "traits")]
        traits: String,
    },
    #[serde(rename="password")]
    SubmitSelfServiceRegistrationFlowWithPasswordMethodBody {
        /// The CSRF Token
        #[serde(rename = "csrf_token", skip_serializing_if = "Option::is_none")]
        csrf_token: Option<String>,
        /// Password to sign the user up with
        #[serde(rename = "password")]
        password: String,
        /// The identity's traits
        #[serde(rename = "traits")]
        traits: serde_json::Value,
    },
}



/// Method to use  This field must be set to `password` when using the password method.
#[derive(Clone, Copy, Debug, Eq, PartialEq, Ord, PartialOrd, Hash, Serialize, Deserialize)]
pub enum Method {
    #[serde(rename = "password")]
    Password,
    #[serde(rename = "oidc")]
    Oidc,
}

