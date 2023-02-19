package matchmaker

import "time"
                    
// AuthUserRequest this API allows the external match-making service to confirm that the user has a valid Session Ticket for the title, in
// order to securely enable match-making. The client passes the user's Session Ticket to the external match-making service,
// which then passes the Session Ticket in as the AuthorizationTicket in this call.
type AuthUserRequestModel struct {
    // AuthorizationTicket session Ticket provided by the client.
    AuthorizationTicket string `json:"AuthorizationTicket,omitempty"`
}

// AuthUserResponse 
type AuthUserResponseModel struct {
    // Authorized boolean indicating if the user has been authorized to use the external match-making service.
    Authorized bool `json:"Authorized"`
    // PlayFabId playFab unique identifier of the account that has been authorized.
    PlayFabId string `json:"PlayFabId,omitempty"`
}

// ItemInstance a unique instance of an item in a user's inventory. Note, to retrieve additional information for an item such as Tags,
// Description that are the same across all instances of the item, a call to GetCatalogItems is required. The ItemID of can
// be matched to a catalog entry, which contains the additional information. Also note that Custom Data is only set when
// the User's specific instance has updated the CustomData via a call to UpdateUserInventoryItemCustomData. Other fields
// such as UnitPrice and UnitCurrency are only set when the item was granted via a purchase.
type ItemInstanceModel struct {
    // Annotation game specific comment associated with this instance when it was added to the user inventory.
    Annotation string `json:"Annotation,omitempty"`
    // BundleContents array of unique items that were awarded when this catalog item was purchased.
    BundleContents []string `json:"BundleContents,omitempty"`
    // BundleParent unique identifier for the parent inventory item, as defined in the catalog, for object which were added from a bundle or
// container.
    BundleParent string `json:"BundleParent,omitempty"`
    // CatalogVersion catalog version for the inventory item, when this instance was created.
    CatalogVersion string `json:"CatalogVersion,omitempty"`
    // CustomData a set of custom key-value pairs on the instance of the inventory item, which is not to be confused with the catalog
// item's custom data.
    CustomData map[string]string `json:"CustomData,omitempty"`
    // DisplayName catalogItem.DisplayName at the time this item was purchased.
    DisplayName string `json:"DisplayName,omitempty"`
    // Expiration timestamp for when this instance will expire.
    Expiration time.Time `json:"Expiration,omitempty"`
    // ItemClass class name for the inventory item, as defined in the catalog.
    ItemClass string `json:"ItemClass,omitempty"`
    // ItemId unique identifier for the inventory item, as defined in the catalog.
    ItemId string `json:"ItemId,omitempty"`
    // ItemInstanceId unique item identifier for this specific instance of the item.
    ItemInstanceId string `json:"ItemInstanceId,omitempty"`
    // PurchaseDate timestamp for when this instance was purchased.
    PurchaseDate time.Time `json:"PurchaseDate,omitempty"`
    // RemainingUses total number of remaining uses, if this is a consumable item.
    RemainingUses int32 `json:"RemainingUses,omitempty"`
    // UnitCurrency currency type for the cost of the catalog item. Not available when granting items.
    UnitCurrency string `json:"UnitCurrency,omitempty"`
    // UnitPrice cost of the catalog item in the given currency. Not available when granting items.
    UnitPrice uint32 `json:"UnitPrice,omitempty"`
    // UsesIncrementedBy the number of uses that were added or removed to this item in this call.
    UsesIncrementedBy int32 `json:"UsesIncrementedBy,omitempty"`
}

// PlayerJoinedRequest 
type PlayerJoinedRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // LobbyId unique identifier of the Game Server Instance the user is joining. This must be a Game Server Instance started with the
// Matchmaker/StartGame API.
    LobbyId string `json:"LobbyId,omitempty"`
    // PlayFabId playFab unique identifier for the player joining.
    PlayFabId string `json:"PlayFabId,omitempty"`
}

// PlayerJoinedResponse 
type PlayerJoinedResponseModel struct {
}

// PlayerLeftRequest 
type PlayerLeftRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // LobbyId unique identifier of the Game Server Instance the user is leaving. This must be a Game Server Instance started with the
// Matchmaker/StartGame API.
    LobbyId string `json:"LobbyId,omitempty"`
    // PlayFabId playFab unique identifier for the player leaving.
    PlayFabId string `json:"PlayFabId,omitempty"`
}

// PlayerLeftResponse 
type PlayerLeftResponseModel struct {
}

// UserInfoRequest 
type UserInfoRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // MinCatalogVersion minimum catalog version for which data is requested (filters the results to only contain inventory items which have a
// catalog version of this or higher).
    MinCatalogVersion int32 `json:"MinCatalogVersion,omitempty"`
    // PlayFabId playFab unique identifier of the user whose information is being requested.
    PlayFabId string `json:"PlayFabId,omitempty"`
}

// UserInfoResponse 
type UserInfoResponseModel struct {
    // Inventory array of inventory items in the user's current inventory.
    Inventory []ItemInstanceModel `json:"Inventory,omitempty"`
    // IsDeveloper boolean indicating whether the user is a developer.
    IsDeveloper bool `json:"IsDeveloper"`
    // PlayFabId playFab unique identifier of the user whose information was requested.
    PlayFabId string `json:"PlayFabId,omitempty"`
    // SteamId steam unique identifier, if the user has an associated Steam account.
    SteamId string `json:"SteamId,omitempty"`
    // TitleDisplayName title specific display name, if set.
    TitleDisplayName string `json:"TitleDisplayName,omitempty"`
    // Username playFab unique user name.
    Username string `json:"Username,omitempty"`
    // VirtualCurrency array of virtual currency balance(s) belonging to the user.
    VirtualCurrency map[string]int32 `json:"VirtualCurrency,omitempty"`
    // VirtualCurrencyRechargeTimes array of remaining times and timestamps for virtual currencies.
    VirtualCurrencyRechargeTimes map[string]VirtualCurrencyRechargeTimeModel `json:"VirtualCurrencyRechargeTimes,omitempty"`
}

// VirtualCurrencyRechargeTime 
type VirtualCurrencyRechargeTimeModel struct {
    // RechargeMax maximum value to which the regenerating currency will automatically increment. Note that it can exceed this value
// through use of the AddUserVirtualCurrency API call. However, it will not regenerate automatically until it has fallen
// below this value.
    RechargeMax int32 `json:"RechargeMax,omitempty"`
    // RechargeTime server timestamp in UTC indicating the next time the virtual currency will be incremented.
    RechargeTime time.Time `json:"RechargeTime,omitempty"`
    // SecondsToRecharge time remaining (in seconds) before the next recharge increment of the virtual currency.
    SecondsToRecharge int32 `json:"SecondsToRecharge,omitempty"`
}
