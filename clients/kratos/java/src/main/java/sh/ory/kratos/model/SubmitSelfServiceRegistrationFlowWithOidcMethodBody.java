/*
 * Ory Kratos API
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * The version of the OpenAPI document: v0.7.1-alpha.1
 * Contact: hi@ory.sh
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package sh.ory.kratos.model;

import java.util.Objects;
import java.util.Arrays;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;

/**
 * SubmitSelfServiceRegistrationFlowWithOidcMethodBody is used to decode the registration form payload when using the oidc method.
 */
@ApiModel(description = "SubmitSelfServiceRegistrationFlowWithOidcMethodBody is used to decode the registration form payload when using the oidc method.")
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2021-07-22T18:06:07.711120196Z[Etc/UTC]")
public class SubmitSelfServiceRegistrationFlowWithOidcMethodBody {
  public static final String SERIALIZED_NAME_CSRF_TOKEN = "csrf_token";
  @SerializedName(SERIALIZED_NAME_CSRF_TOKEN)
  private String csrfToken;

  public static final String SERIALIZED_NAME_METHOD = "method";
  @SerializedName(SERIALIZED_NAME_METHOD)
  private String method;

  public static final String SERIALIZED_NAME_TRAITS = "traits";
  @SerializedName(SERIALIZED_NAME_TRAITS)
  private String traits;


  public SubmitSelfServiceRegistrationFlowWithOidcMethodBody csrfToken(String csrfToken) {
    
    this.csrfToken = csrfToken;
    return this;
  }

   /**
   * The CSRF Token
   * @return csrfToken
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "The CSRF Token")

  public String getCsrfToken() {
    return csrfToken;
  }


  public void setCsrfToken(String csrfToken) {
    this.csrfToken = csrfToken;
  }


  public SubmitSelfServiceRegistrationFlowWithOidcMethodBody method(String method) {
    
    this.method = method;
    return this;
  }

   /**
   * Method to use  This field must be set to &#x60;oidc&#x60; when using the oidc method.
   * @return method
  **/
  @ApiModelProperty(required = true, value = "Method to use  This field must be set to `oidc` when using the oidc method.")

  public String getMethod() {
    return method;
  }


  public void setMethod(String method) {
    this.method = method;
  }


  public SubmitSelfServiceRegistrationFlowWithOidcMethodBody traits(String traits) {
    
    this.traits = traits;
    return this;
  }

   /**
   * The provider to register with
   * @return traits
  **/
  @ApiModelProperty(required = true, value = "The provider to register with")

  public String getTraits() {
    return traits;
  }


  public void setTraits(String traits) {
    this.traits = traits;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    SubmitSelfServiceRegistrationFlowWithOidcMethodBody submitSelfServiceRegistrationFlowWithOidcMethodBody = (SubmitSelfServiceRegistrationFlowWithOidcMethodBody) o;
    return Objects.equals(this.csrfToken, submitSelfServiceRegistrationFlowWithOidcMethodBody.csrfToken) &&
        Objects.equals(this.method, submitSelfServiceRegistrationFlowWithOidcMethodBody.method) &&
        Objects.equals(this.traits, submitSelfServiceRegistrationFlowWithOidcMethodBody.traits);
  }

  @Override
  public int hashCode() {
    return Objects.hash(csrfToken, method, traits);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class SubmitSelfServiceRegistrationFlowWithOidcMethodBody {\n");
    sb.append("    csrfToken: ").append(toIndentedString(csrfToken)).append("\n");
    sb.append("    method: ").append(toIndentedString(method)).append("\n");
    sb.append("    traits: ").append(toIndentedString(traits)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }

}

