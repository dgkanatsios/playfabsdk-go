package data

import "time"
                    
// AbortFileUploadsRequest aborts the pending upload of the requested files.
type AbortFileUploadsRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Entity the entity to perform this action on.
    Entity* EntityKeyModel `json:"Entity,omitempty"`
    // FileNames names of the files to have their pending uploads aborted.
    FileNames []string `json:"FileNames,omitempty"`
    // ProfileVersion the expected version of the profile, if set and doesn't match the current version of the profile the operation will not
// be performed.
    ProfileVersion int32 `json:"ProfileVersion,omitempty"`
}

// AbortFileUploadsResponse 
type AbortFileUploadsResponseModel struct {
    // Entity the entity id and type.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // ProfileVersion the current version of the profile, can be used for concurrency control during updates.
    ProfileVersion int32 `json:"ProfileVersion,omitempty"`
}

// DeleteFilesRequest deletes the requested files from the entity's profile.
type DeleteFilesRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Entity the entity to perform this action on.
    Entity* EntityKeyModel `json:"Entity,omitempty"`
    // FileNames names of the files to be deleted.
    FileNames []string `json:"FileNames,omitempty"`
    // ProfileVersion the expected version of the profile, if set and doesn't match the current version of the profile the operation will not
// be performed.
    ProfileVersion int32 `json:"ProfileVersion,omitempty"`
}

// DeleteFilesResponse 
type DeleteFilesResponseModel struct {
    // Entity the entity id and type.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // ProfileVersion the current version of the profile, can be used for concurrency control during updates.
    ProfileVersion int32 `json:"ProfileVersion,omitempty"`
}

// EntityKey combined entity type and ID structure which uniquely identifies a single entity.
type EntityKeyModel struct {
    // Id unique ID of the entity.
    Id string `json:"Id,omitempty"`
    // Type entity type. See https://docs.microsoft.com/gaming/playfab/features/data/entities/available-built-in-entity-types
    Type string `json:"Type,omitempty"`
}

// FinalizeFileUploadsRequest finalizes the upload of the requested files. Verifies that the files have been successfully uploaded and moves the file
// pointers from pending to live.
type FinalizeFileUploadsRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Entity the entity to perform this action on.
    Entity* EntityKeyModel `json:"Entity,omitempty"`
    // FileNames names of the files to be finalized. Restricted to a-Z, 0-9, '(', ')', '_', '-' and '.'
    FileNames []string `json:"FileNames,omitempty"`
    // ProfileVersion the current version of the profile, can be used for concurrency control during updates.
    ProfileVersion int32 `json:"ProfileVersion,omitempty"`
}

// FinalizeFileUploadsResponse 
type FinalizeFileUploadsResponseModel struct {
    // Entity the entity id and type.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // Metadata collection of metadata for the entity's files
    Metadata map[string]GetFileMetadataModel `json:"Metadata,omitempty"`
    // ProfileVersion the current version of the profile, can be used for concurrency control during updates.
    ProfileVersion int32 `json:"ProfileVersion,omitempty"`
}

// GetFileMetadata 
type GetFileMetadataModel struct {
    // Checksum checksum value for the file
    Checksum string `json:"Checksum,omitempty"`
    // DownloadUrl download URL where the file can be retrieved
    DownloadUrl string `json:"DownloadUrl,omitempty"`
    // FileName name of the file
    FileName string `json:"FileName,omitempty"`
    // LastModified last UTC time the file was modified
    LastModified time.Time `json:"LastModified,omitempty"`
    // Size storage service's reported byte count
    Size int32 `json:"Size,omitempty"`
}

// GetFilesRequest returns URLs that may be used to download the files for a profile for a limited length of time. Only returns files that
// have been successfully uploaded, files that are still pending will either return the old value, if it exists, or
// nothing.
type GetFilesRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Entity the entity to perform this action on.
    Entity* EntityKeyModel `json:"Entity,omitempty"`
}

// GetFilesResponse 
type GetFilesResponseModel struct {
    // Entity the entity id and type.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // Metadata collection of metadata for the entity's files
    Metadata map[string]GetFileMetadataModel `json:"Metadata,omitempty"`
    // ProfileVersion the current version of the profile, can be used for concurrency control during updates.
    ProfileVersion int32 `json:"ProfileVersion,omitempty"`
}

// GetObjectsRequest gets JSON objects from an entity profile and returns it.
type GetObjectsRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Entity the entity to perform this action on.
    Entity* EntityKeyModel `json:"Entity,omitempty"`
    // EscapeObject determines whether the object will be returned as an escaped JSON string or as a un-escaped JSON object. Default is JSON
// object.
    EscapeObject bool `json:"EscapeObject"`
}

// GetObjectsResponse 
type GetObjectsResponseModel struct {
    // Entity the entity id and type.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // Objects requested objects that the calling entity has access to
    Objects map[string]ObjectResultModel `json:"Objects,omitempty"`
    // ProfileVersion the current version of the profile, can be used for concurrency control during updates.
    ProfileVersion int32 `json:"ProfileVersion,omitempty"`
}

// InitiateFileUploadMetadata 
type InitiateFileUploadMetadataModel struct {
    // FileName name of the file.
    FileName string `json:"FileName,omitempty"`
    // UploadUrl location the data should be sent to via an HTTP PUT operation.
    UploadUrl string `json:"UploadUrl,omitempty"`
}

// InitiateFileUploadsRequest returns URLs that may be used to upload the files for a profile 5 minutes. After using the upload calls
// FinalizeFileUploads must be called to move the file status from pending to live.
type InitiateFileUploadsRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Entity the entity to perform this action on.
    Entity* EntityKeyModel `json:"Entity,omitempty"`
    // FileNames names of the files to be set. Restricted to a-Z, 0-9, '(', ')', '_', '-' and '.'
    FileNames []string `json:"FileNames,omitempty"`
    // ProfileVersion the expected version of the profile, if set and doesn't match the current version of the profile the operation will not
// be performed.
    ProfileVersion int32 `json:"ProfileVersion,omitempty"`
}

// InitiateFileUploadsResponse 
type InitiateFileUploadsResponseModel struct {
    // Entity the entity id and type.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // ProfileVersion the current version of the profile, can be used for concurrency control during updates.
    ProfileVersion int32 `json:"ProfileVersion,omitempty"`
    // UploadDetails collection of file names and upload urls
    UploadDetails []InitiateFileUploadMetadataModel `json:"UploadDetails,omitempty"`
}

// ObjectResult 
type ObjectResultModel struct {
    // DataObject un-escaped JSON object, if EscapeObject false or default.
    DataObject interface{} `json:"DataObject,omitempty"`
    // EscapedDataObject escaped string JSON body of the object, if EscapeObject is true.
    EscapedDataObject string `json:"EscapedDataObject,omitempty"`
    // ObjectName name of the object. Restricted to a-Z, 0-9, '(', ')', '_', '-' and '.'
    ObjectName string `json:"ObjectName,omitempty"`
}

// OperationTypes 
type OperationTypes string
  
const (
     OperationTypesCreated OperationTypes = "Created"
     OperationTypesUpdated OperationTypes = "Updated"
     OperationTypesDeleted OperationTypes = "Deleted"
     OperationTypesNone OperationTypes = "None"
      )
// SetObject 
type SetObjectModel struct {
    // DataObject body of the object to be saved. If empty and DeleteObject is true object will be deleted if it exists, or no operation
// will occur if it does not exist. Only one of Object or EscapedDataObject fields may be used.
    DataObject interface{} `json:"DataObject,omitempty"`
    // DeleteObject flag to indicate that this object should be deleted. Both DataObject and EscapedDataObject must not be set as well.
    DeleteObject bool `json:"DeleteObject"`
    // EscapedDataObject body of the object to be saved as an escaped JSON string. If empty and DeleteObject is true object will be deleted if it
// exists, or no operation will occur if it does not exist. Only one of DataObject or EscapedDataObject fields may be used.
    EscapedDataObject string `json:"EscapedDataObject,omitempty"`
    // ObjectName name of object. Restricted to a-Z, 0-9, '(', ')', '_', '-' and '.'.
    ObjectName string `json:"ObjectName,omitempty"`
}

// SetObjectInfo 
type SetObjectInfoModel struct {
    // ObjectName name of the object
    ObjectName string `json:"ObjectName,omitempty"`
    // OperationReason optional reason to explain why the operation was the result that it was.
    OperationReason string `json:"OperationReason,omitempty"`
    // SetResult indicates which operation was completed, either Created, Updated, Deleted or None.
    SetResult OperationTypes `json:"SetResult,omitempty"`
}

// SetObjectsRequest sets JSON objects on the requested entity profile. May include a version number to be used to perform optimistic
// concurrency operations during update. If the current version differs from the version in the request the request will be
// ignored. If no version is set on the request then the value will always be updated if the values differ. Using the
// version value does not guarantee a write though, ConcurrentEditError may still occur if multiple clients are attempting
// to update the same profile.
type SetObjectsRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Entity the entity to perform this action on.
    Entity* EntityKeyModel `json:"Entity,omitempty"`
    // ExpectedProfileVersion optional field used for concurrency control. By specifying the previously returned value of ProfileVersion from
// GetProfile API, you can ensure that the object set will only be performed if the profile has not been updated by any
// other clients since the version you last loaded.
    ExpectedProfileVersion int32 `json:"ExpectedProfileVersion,omitempty"`
    // Objects collection of objects to set on the profile.
    Objects []SetObjectModel `json:"Objects,omitempty"`
}

// SetObjectsResponse 
type SetObjectsResponseModel struct {
    // ProfileVersion new version of the entity profile.
    ProfileVersion int32 `json:"ProfileVersion,omitempty"`
    // SetResults new version of the entity profile.
    SetResults []SetObjectInfoModel `json:"SetResults,omitempty"`
}
