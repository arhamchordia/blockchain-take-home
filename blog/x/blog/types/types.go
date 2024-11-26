package types

// Event types and attribute keys for blog module
const (
	EventTypeCreatePost   = "create_post"
	EventTypeDeletePost   = "delete_post"
	EventTypeAddEditor    = "add_editor"
	EventTypeDeleteEditor = "delete_editor"
	EventTypeUpdatePost   = "update_post"
	EventTypeUpdateParams = "update_params"

	AttributeKeyAuthority  = "authority"
	AttributeKeyParams     = "params"
	AttributeKeyUpdateTime = "update_time"
	AttributeKeyEditor     = "editor"
	AttributeKeyDeleter    = "deleter"
	AttributeKeyPostID     = "post_id"
	AttributeKeyCreator    = "creator"
	AttributeKeyTitle      = "title"
)
