package authentication

import "time"
                    
// EntityKey combined entity type and ID structure which uniquely identifies a single entity.
type EntityKeyModel struct {
    // Id unique ID of the entity.
    Id string `json:"Id"`
    // Type entity type. See https://docs.microsoft.com/gaming/playfab/features/data/entities/available-built-in-entity-types
    Type string `json:"Type"`
}

// EntityLineage 
type EntityLineageModel struct {
    // CharacterId the Character Id of the associated entity.
    CharacterId string `json:"CharacterId"`
    // GroupId the Group Id of the associated entity.
    GroupId string `json:"GroupId"`
    // MasterPlayerAccountId the Master Player Account Id of the associated entity.
    MasterPlayerAccountId string `json:"MasterPlayerAccountId"`
    // NamespaceId the Namespace Id of the associated entity.
    NamespaceId string `json:"NamespaceId"`
    // TitleId the Title Id of the associated entity.
    TitleId string `json:"TitleId"`
    // TitlePlayerAccountId the Title Player Account Id of the associated entity.
    TitlePlayerAccountId string `json:"TitlePlayerAccountId"`
}

// GetEntityTokenRequest this API must be called with X-SecretKey, X-Authentication or X-EntityToken headers. An optional EntityKey may be
// included to attempt to set the resulting EntityToken to a specific entity, however the entity must be a relation of the
// caller, such as the master_player_account of a character. If sending X-EntityToken the account will be marked as freshly
// logged in and will issue a new token. If using X-Authentication or X-EntityToken the header must still be valid and
// cannot be expired or revoked.
type GetEntityTokenRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
}

// GetEntityTokenResponse 
type GetEntityTokenResponseModel struct {
    // Entity the entity id and type.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // EntityToken the token used to set X-EntityToken for all entity based API calls.
    EntityToken string `json:"EntityToken"`
    // TokenExpiration the time the token will expire, if it is an expiring token, in UTC.
    TokenExpiration time.Time `json:"TokenExpiration"`
}

// IdentifiedDeviceType 
type IdentifiedDeviceType string
  
const (
     IdentifiedDeviceTypeUnknown IdentifiedDeviceType = "Unknown"
     IdentifiedDeviceTypeXboxOne IdentifiedDeviceType = "XboxOne"
     IdentifiedDeviceTypeScarlett IdentifiedDeviceType = "Scarlett"
      )
// LoginIdentityProvider 
type LoginIdentityProvider string
  
const (
     LoginIdentityProviderUnknown LoginIdentityProvider = "Unknown"
     LoginIdentityProviderPlayFab LoginIdentityProvider = "PlayFab"
     LoginIdentityProviderCustom LoginIdentityProvider = "Custom"
     LoginIdentityProviderGameCenter LoginIdentityProvider = "GameCenter"
     LoginIdentityProviderGooglePlay LoginIdentityProvider = "GooglePlay"
     LoginIdentityProviderSteam LoginIdentityProvider = "Steam"
     LoginIdentityProviderXBoxLive LoginIdentityProvider = "XBoxLive"
     LoginIdentityProviderPSN LoginIdentityProvider = "PSN"
     LoginIdentityProviderKongregate LoginIdentityProvider = "Kongregate"
     LoginIdentityProviderFacebook LoginIdentityProvider = "Facebook"
     LoginIdentityProviderIOSDevice LoginIdentityProvider = "IOSDevice"
     LoginIdentityProviderAndroidDevice LoginIdentityProvider = "AndroidDevice"
     LoginIdentityProviderTwitch LoginIdentityProvider = "Twitch"
     LoginIdentityProviderWindowsHello LoginIdentityProvider = "WindowsHello"
     LoginIdentityProviderGameServer LoginIdentityProvider = "GameServer"
     LoginIdentityProviderCustomServer LoginIdentityProvider = "CustomServer"
     LoginIdentityProviderNintendoSwitch LoginIdentityProvider = "NintendoSwitch"
     LoginIdentityProviderFacebookInstantGames LoginIdentityProvider = "FacebookInstantGames"
     LoginIdentityProviderOpenIdConnect LoginIdentityProvider = "OpenIdConnect"
     LoginIdentityProviderApple LoginIdentityProvider = "Apple"
     LoginIdentityProviderNintendoSwitchAccount LoginIdentityProvider = "NintendoSwitchAccount"
      )
// ValidateEntityTokenRequest given an entity token, validates that it hasn't expired or been revoked and will return details of the owner.
type ValidateEntityTokenRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // EntityToken client EntityToken
    EntityToken string `json:"EntityToken"`
}

// ValidateEntityTokenResponse 
type ValidateEntityTokenResponseModel struct {
    // Entity the entity id and type.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // IdentifiedDeviceType the authenticated device for this entity, for the given login
    IdentifiedDeviceType IdentifiedDeviceType `json:"IdentifiedDeviceType"`
    // IdentityProvider the identity provider for this entity, for the given login
    IdentityProvider LoginIdentityProvider `json:"IdentityProvider"`
    // Lineage the lineage of this profile.
    Lineage *EntityLineageModel `json:"Lineage,omitempty"`
}
