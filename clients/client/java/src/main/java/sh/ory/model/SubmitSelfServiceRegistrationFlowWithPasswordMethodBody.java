/*
 * Ory APIs
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v0.0.1-alpha.17
 * Contact: support@ory.sh
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package sh.ory.model;

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
 * SubmitSelfServiceRegistrationFlowWithPasswordMethodBody is used to decode the registration form payload when using the password method.
 */
@ApiModel(description = "SubmitSelfServiceRegistrationFlowWithPasswordMethodBody is used to decode the registration form payload when using the password method.")
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2021-08-01T10:37:52.194203415Z[Etc/UTC]")
public class SubmitSelfServiceRegistrationFlowWithPasswordMethodBody {
  public static final String SERIALIZED_NAME_CSRF_TOKEN = "csrf_token";
  @SerializedName(SERIALIZED_NAME_CSRF_TOKEN)
  private String csrfToken;

  /**
   * Method to use  This field must be set to &#x60;password&#x60; when using the password method.
   */
  @JsonAdapter(MethodEnum.Adapter.class)
  public enum MethodEnum {
    PASSWORD("password"),
    
    OIDC("oidc");

    private String value;

    MethodEnum(String value) {
      this.value = value;
    }

    public String getValue() {
      return value;
    }

    @Override
    public String toString() {
      return String.valueOf(value);
    }

    public static MethodEnum fromValue(String value) {
      for (MethodEnum b : MethodEnum.values()) {
        if (b.value.equals(value)) {
          return b;
        }
      }
      throw new IllegalArgumentException("Unexpected value '" + value + "'");
    }

    public static class Adapter extends TypeAdapter<MethodEnum> {
      @Override
      public void write(final JsonWriter jsonWriter, final MethodEnum enumeration) throws IOException {
        jsonWriter.value(enumeration.getValue());
      }

      @Override
      public MethodEnum read(final JsonReader jsonReader) throws IOException {
        String value =  jsonReader.nextString();
        return MethodEnum.fromValue(value);
      }
    }
  }

  public static final String SERIALIZED_NAME_METHOD = "method";
  @SerializedName(SERIALIZED_NAME_METHOD)
  private MethodEnum method;

  public static final String SERIALIZED_NAME_PASSWORD = "password";
  @SerializedName(SERIALIZED_NAME_PASSWORD)
  private String password;

  public static final String SERIALIZED_NAME_TRAITS = "traits";
  @SerializedName(SERIALIZED_NAME_TRAITS)
  private Object traits;


  public SubmitSelfServiceRegistrationFlowWithPasswordMethodBody csrfToken(String csrfToken) {
    
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


  public SubmitSelfServiceRegistrationFlowWithPasswordMethodBody method(MethodEnum method) {
    
    this.method = method;
    return this;
  }

   /**
   * Method to use  This field must be set to &#x60;password&#x60; when using the password method.
   * @return method
  **/
  @ApiModelProperty(required = true, value = "Method to use  This field must be set to `password` when using the password method.")

  public MethodEnum getMethod() {
    return method;
  }


  public void setMethod(MethodEnum method) {
    this.method = method;
  }


  public SubmitSelfServiceRegistrationFlowWithPasswordMethodBody password(String password) {
    
    this.password = password;
    return this;
  }

   /**
   * Password to sign the user up with
   * @return password
  **/
  @ApiModelProperty(required = true, value = "Password to sign the user up with")

  public String getPassword() {
    return password;
  }


  public void setPassword(String password) {
    this.password = password;
  }


  public SubmitSelfServiceRegistrationFlowWithPasswordMethodBody traits(Object traits) {
    
    this.traits = traits;
    return this;
  }

   /**
   * The identity&#39;s traits
   * @return traits
  **/
  @ApiModelProperty(required = true, value = "The identity's traits")

  public Object getTraits() {
    return traits;
  }


  public void setTraits(Object traits) {
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
    SubmitSelfServiceRegistrationFlowWithPasswordMethodBody submitSelfServiceRegistrationFlowWithPasswordMethodBody = (SubmitSelfServiceRegistrationFlowWithPasswordMethodBody) o;
    return Objects.equals(this.csrfToken, submitSelfServiceRegistrationFlowWithPasswordMethodBody.csrfToken) &&
        Objects.equals(this.method, submitSelfServiceRegistrationFlowWithPasswordMethodBody.method) &&
        Objects.equals(this.password, submitSelfServiceRegistrationFlowWithPasswordMethodBody.password) &&
        Objects.equals(this.traits, submitSelfServiceRegistrationFlowWithPasswordMethodBody.traits);
  }

  @Override
  public int hashCode() {
    return Objects.hash(csrfToken, method, password, traits);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class SubmitSelfServiceRegistrationFlowWithPasswordMethodBody {\n");
    sb.append("    csrfToken: ").append(toIndentedString(csrfToken)).append("\n");
    sb.append("    method: ").append(toIndentedString(method)).append("\n");
    sb.append("    password: ").append(toIndentedString(password)).append("\n");
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

