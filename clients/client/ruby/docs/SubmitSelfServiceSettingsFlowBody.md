# OryHydraClient::SubmitSelfServiceSettingsFlowBody

## Class instance methods

### `openapi_one_of`

Returns the list of classes defined in oneOf.

#### Example

```ruby
require 'ory-client'

OryHydraClient::SubmitSelfServiceSettingsFlowBody.openapi_one_of
# =>
# [
#   :'SubmitSelfServiceSettingsFlowWithPasswordMethodBody',
#   :'SubmitSelfServiceSettingsFlowWithProfileMethodBody'
# ]
```

### `openapi_discriminator_name`

Returns the discriminator's property name.

#### Example

```ruby
require 'ory-client'

OryHydraClient::SubmitSelfServiceSettingsFlowBody.openapi_discriminator_name
# => :'method'
```

### `openapi_discriminator_name`

Returns the discriminator's mapping.

#### Example

```ruby
require 'ory-client'

OryHydraClient::SubmitSelfServiceSettingsFlowBody.openapi_discriminator_mapping
# =>
# {
#   :'oidc' => :'SubmitSelfServiceSettingsFlowWithOidcMethodBody',
#   :'password' => :'SubmitSelfServiceSettingsFlowWithPasswordMethodBody',
#   :'profile' => :'SubmitSelfServiceSettingsFlowWithProfileMethodBody'
# }

### build

Find the appropriate object from the `openapi_one_of` list and casts the data into it.

#### Example

```ruby
require 'ory-client'

OryHydraClient::SubmitSelfServiceSettingsFlowBody.build(data)
# => #<SubmitSelfServiceSettingsFlowWithPasswordMethodBody:0x00007fdd4aab02a0>

OryHydraClient::SubmitSelfServiceSettingsFlowBody.build(data_that_doesnt_match)
# => nil
```

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| **data** | **Mixed** | data to be matched against the list of oneOf items |

#### Return type

- `SubmitSelfServiceSettingsFlowWithPasswordMethodBody`
- `SubmitSelfServiceSettingsFlowWithProfileMethodBody`
- `nil` (if no type matches)

