package scim2

type AttributeType string

const (
	AttributeType۰STRING               AttributeType = "string"
	AttributeType۰BOOLEAN              AttributeType = "boolean"
	AttributeType۰DECIMALAttributeType               = "decimal"
	AttributeType۰INTEGER              AttributeType = "integer"
	AttributeType۰DATETIME             AttributeType = "datetime"
	AttributeType۰BINARY               AttributeType = "binary"
	AttributeType۰REFERENCE            AttributeType = "reference"
	AttributeType۰COMPLEX              AttributeType = "complex"
)

// Http method's are in http.MethodGet, http.MethodPost

type Mutability string

const (
	Mutability۰READ_ONLY  Mutability = "readOnly"
	Mutability۰READ_WRITE Mutability = "readWrite"
	Mutability۰IMMUTABLE  Mutability = "immutable"
	Mutability۰WRITE_ONLY Mutability = "writeOnly"
)

type PatchOp string

const (
	PatchOp۰ADD     PatchOp = "add"
	PatchOp۰REMOVE  PatchOp = "remove"
	PatchOp۰REPLACE PatchOp = "replace"
)

type ReturnedType string

const (
	ReturnedType۰ALWAYS  ReturnedType = "always"
	ReturnedType۰NEVER   ReturnedType = "never"
	ReturnedType۰DEFAULT ReturnedType = "default"
	ReturnedType۰REQUEST ReturnedType = "request"
)

type ScimOperation uint8

const (
	ScimOperation۰CREATE ScimOperation = iota
	ScimOperation۰READ
	ScimOperation۰REPLACE
	ScimOperation۰DELETE
	ScimOperation۰PATCH
	ScimOperation۰SEARCH
	ScimOperation۰BULK
)

type SortOrder uint8

const (
	SortOrder۰ASCENDING SortOrder = iota
	SortOrder۰DESCENDING
)

type UniquenessType string

const (
	UniquenessType۰NONE   UniquenessType = "none"
	UniquenessType۰SERVER UniquenessType = "server"
	UniquenessType۰GLOBAL UniquenessType = "global"
)
