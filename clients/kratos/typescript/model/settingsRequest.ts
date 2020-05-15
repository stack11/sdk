/**
 * Ory Kratos
 * Welcome to the ORY Kratos HTTP API documentation!
 *
 * The version of the OpenAPI document: latest
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from '../api';
import { Identity } from './identity';
import { SettingsRequestMethod } from './settingsRequestMethod';

/**
* This request is used when an identity wants to update settings (e.g. profile data, passwords, ...) in a selfservice manner.  We recommend reading the [User Settings Documentation](../self-service/flows/user-settings)
*/
export class SettingsRequest {
    /**
    * Active, if set, contains the registration method that is being used. It is initially not set.
    */
    'active'?: string;
    /**
    * ExpiresAt is the time (UTC) when the request expires. If the user still wishes to update the setting, a new request has to be initiated.
    */
    'expiresAt': Date;
    'id': string;
    'identity': Identity;
    /**
    * IssuedAt is the time (UTC) when the request occurred.
    */
    'issuedAt': Date;
    /**
    * Methods contains context for all enabled registration methods. If a registration request has been processed, but for example the password is incorrect, this will contain error messages.
    */
    'methods': { [key: string]: SettingsRequestMethod; };
    /**
    * RequestURL is the initial URL that was requested from ORY Kratos. It can be used to forward information contained in the URL\'s path or query for example.
    */
    'requestUrl': string;
    /**
    * UpdateSuccessful, if true, indicates that the settings request has been updated successfully with the provided data. Done will stay true when repeatedly checking. If set to true, done will revert back to false only when a request with invalid (e.g. \"please use a valid phone number\") data was sent.
    */
    'updateSuccessful': boolean;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "active",
            "baseName": "active",
            "type": "string"
        },
        {
            "name": "expiresAt",
            "baseName": "expires_at",
            "type": "Date"
        },
        {
            "name": "id",
            "baseName": "id",
            "type": "string"
        },
        {
            "name": "identity",
            "baseName": "identity",
            "type": "Identity"
        },
        {
            "name": "issuedAt",
            "baseName": "issued_at",
            "type": "Date"
        },
        {
            "name": "methods",
            "baseName": "methods",
            "type": "{ [key: string]: SettingsRequestMethod; }"
        },
        {
            "name": "requestUrl",
            "baseName": "request_url",
            "type": "string"
        },
        {
            "name": "updateSuccessful",
            "baseName": "update_successful",
            "type": "boolean"
        }    ];

    static getAttributeTypeMap() {
        return SettingsRequest.attributeTypeMap;
    }
}

