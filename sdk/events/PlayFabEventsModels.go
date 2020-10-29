package events

import "time"
                    
// EntityKey combined entity type and ID structure which uniquely identifies a single entity.
type EntityKeyModel struct {
    // Id unique ID of the entity.
    Id string `json:"Id,omitempty"`
    // Type entity type. See https://docs.microsoft.com/gaming/playfab/features/data/entities/available-built-in-entity-types
    Type string `json:"Type,omitempty"`
}

// EventContents 
type EventContentsModel struct {
    // CustomTags the optional custom tags associated with the event (e.g. build number, external trace identifiers, etc.). Before an
// event is written, this collection and the base request custom tags will be merged, but not overriden. This enables the
// caller to specify static tags and per event tags.
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Entity entity associated with the event. If null, the event will apply to the calling entity.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // EventNamespace the namespace in which the event is defined. Allowed namespaces can vary by API.
    EventNamespace string `json:"EventNamespace,omitempty"`
    // Name the name of this event.
    Name string `json:"Name,omitempty"`
    // OriginalId the original unique identifier associated with this event before it was posted to PlayFab. The value might differ from
// the EventId value, which is assigned when the event is received by the server.
    OriginalId string `json:"OriginalId,omitempty"`
    // OriginalTimestamp the time (in UTC) associated with this event when it occurred. If specified, this value is stored in the
// OriginalTimestamp property of the PlayStream event.
    OriginalTimestamp time.Time `json:"OriginalTimestamp,omitempty"`
    // Payload arbitrary data associated with the event. Only one of Payload or PayloadJSON is allowed.
    Payload interface{} `json:"Payload,omitempty"`
    // PayloadJSON arbitrary data associated with the event, represented as a JSON serialized string. Only one of Payload or PayloadJSON is
// allowed.
    PayloadJSON string `json:"PayloadJSON,omitempty"`
}

// WriteEventsRequest 
type WriteEventsRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Events collection of events to write to PlayStream.
    Events []EventContentsModel `json:"Events,omitempty"`
}

// WriteEventsResponse 
type WriteEventsResponseModel struct {
    // AssignedEventIds the unique identifiers assigned by the server to the events, in the same order as the events in the request. Only
// returned if FlushToPlayStream option is true.
    AssignedEventIds []string `json:"AssignedEventIds,omitempty"`
}
