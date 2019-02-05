package authentication

import "time"
                    
// EntityKey combined entity type and ID structure which uniquely identifies a single entity.
type EntityKeyModel struct {
    // Id unique ID of the entity.
    Id string `json:"Id,omitempty"`
    // Type entity type. See https://api.playfab.com/docs/tutorials/entities/entitytypes
    Type string `json:"Type,omitempty"`
}

// GetEntityTokenRequest this API must be called with X-SecretKey, X-Authentication or X-EntityToken headers. An optional EntityKey may be
// included to attempt to set the resulting EntityToken to a specific entity, however the entity must be a relation of the
// caller, such as the master_player_account of a character. If sending X-EntityToken the account will be marked as freshly
// logged in and will issue a new token. If using X-Authentication or X-EntityToken the header must still be valid and
// cannot be expired or revoked.
type GetEntityTokenRequestModel struct {
    // Entity the entity to perform this action on.
    Entity EntityKeyModel `json:"Entity,omitempty"`
}

// GetEntityTokenResponse 
type GetEntityTokenResponseModel struct {
    // Entity the entity id and type.
    Entity EntityKeyModel `json:"Entity,omitempty"`
    // EntityToken the token used to set X-EntityToken for all entity based API calls.
    EntityToken string `json:"EntityToken,omitempty"`
    // TokenExpiration the time the token will expire, if it is an expiring token, in UTC.
    TokenExpiration time.Time `json:"TokenExpiration,omitempty"`
}
