package scim2

type MetaRecord struct {
	ResourceType string    `json:"resourceType,omitempty"`
	Created      *JsonDate `json:"created,omitempty"`
	LastModified *JsonDate `json:"lastModified,omitempty"`
	Location     string    `json:"location,omitempty"`
	Version      string    `json:"version,omitempty"`
}

type BaseRecord struct {
	Id         string     `json:"id,omitempty"`
	ExternalId string     `json:"externalId,omitempty"`
	Meta       MetaRecord `json:"meta,omitempty"`
	Schemas    []string   `json:"schemas,omitempty"`
}

type Attribute struct {
	Name        string        `json:"name,omitempty"`
	Description string        `json:"description,omitempty"`
	Type        AttributeType `json:"type,omitempty"`

	CaseExact   bool       `json:"caseExact,omitempty"`
	MultiValued bool       `json:"multiValued,omitempty"`
	Mutability  Mutability `json:"mutability,omitempty"`

	Required   bool           `json:"required,omitempty"`
	Returned   ReturnedType   `json:"returned,omitempty"`
	Uniqueness UniquenessType `json:"uniqueness,omitempty"`

	CanonicalValues []interface{} `json:"canonicalValues,omitempty"`
	ReferenceTypes  []string      `json:"referenceTypes,omitempty"`
	SubAttributes   []*Attribute  `json:"subAttributes,omitempty"`
}

func NewAttribute(name string) *Attribute {
	attr := &Attribute{
		Name:       name,
		Type:       AttributeType۰STRING,
		Mutability: Mutability۰READ_WRITE,
		Returned:   ReturnedType۰DEFAULT,
		Uniqueness: UniquenessType۰NONE,
	}

	return attr
}

func (a *Attribute) SetDescription(value string) *Attribute {
	a.Description = value
	return a
}
func (a *Attribute) SetType(value AttributeType) *Attribute {
	a.Type = value
	return a
}
func (a *Attribute) SetCaseExact(value bool) *Attribute {
	a.CaseExact = value
	return a
}
func (a *Attribute) SetMultiValued(value bool) *Attribute {
	a.MultiValued = value
	return a
}
func (a *Attribute) SetMutability(value Mutability) *Attribute {
	a.Mutability = value
	return a
}
func (a *Attribute) SetRequired(value bool) *Attribute {
	a.Required = value
	return a
}
func (a *Attribute) SetReturned(value ReturnedType) *Attribute {
	a.Returned = value
	return a
}
func (a *Attribute) SetUniqueness(value UniquenessType) *Attribute {
	a.Uniqueness = value
	return a
}
func (a *Attribute) SetCanonicalValues(value []interface{}) *Attribute {
	a.CanonicalValues = value
	return a
}
func (a *Attribute) AddSubAttribute(value *Attribute) *Attribute {
	a.SubAttributes = append(a.SubAttributes, value)
	return a
}

type Schema struct {
	BaseRecord
	Name        string       `json:"name,omitempty"`
	Description string       `json:"description,omitempty"`
	Attributes  []*Attribute `json:"attributes,omitempty"`
}

type ListResponse struct {
	Schemas      []string      `json:"schemas,omitempty"`
	TotalResults int           `json:"totalResults,omitempty"`
	ItemsPerPage int           `json:"itemsPerPage,omitempty"`
	StartIndex   int           `json:"startIndex,omitempty"`
	Resources    []interface{} `json:"resources,omitempty"`
}

// USER Record
type UserName struct {
	Formatted       string `json:"formatted,omitempty"`
	FamilyName      string `json:"familyName,omitempty"`
	GivenName       string `json:"givenName,omitempty"`
	MiddleName      string `json:"middleName,omitempty"`
	HonorificPrefix string `json:"honorificPrefix,omitempty"`
	HonorificSuffix string `json:"honorificSuffix,omitempty"`
}

type UserEmail struct {
	Value   string `json:"value,omitempty"`
	Display string `json:"display,omitempty"`
	Type    string `json:"type,omitempty"`
	Primary bool   `json:"primary,omitempty"`
}

type UserPhoneNumber struct {
	Value   string `json:"value,omitempty"`
	Display string `json:"display,omitempty"`
	Type    string `json:"type,omitempty"`
	Primary bool   `json:"primary,omitempty"`
}

type UserIm struct {
	Value   string `json:"value,omitempty"`
	Display string `json:"display,omitempty"`
	Type    string `json:"type,omitempty"`
	Primary bool   `json:"primary,omitempty"`
}

type UserPhoto struct {
	Value   string `json:"value,omitempty"`
	Type    string `json:"type,omitempty"`
	Primary bool   `json:"primary,omitempty"`
}

type UserAddress struct {
	Formatted     string `json:"formatted,omitempty"`
	StreetAddress string `json:"streetAddress,omitempty"`
	Locality      string `json:"locality,omitempty"`
	Region        string `json:"region,omitempty"`
	PostalCode    string `json:"postalCode,omitempty"`
	Country       string `json:"country,omitempty"`
	Type          string `json:"type,omitempty"`
	Primary       bool   `json:"primary,omitempty"`
}

type UserGroup struct {
	Value   string `json:"value,omitempty"`
	Ref     string `json:"$ref,omitempty"`
	Display string `json:"display,omitempty"`
	Type    string `json:"type,omitempty"`
}

type UserEntitlement struct {
	Value   string `json:"value,omitempty"`
	Display string `json:"display,omitempty"`
	Type    string `json:"type,omitempty"`
	Primary bool   `json:"primary,omitempty"`
}

type UserRole struct {
	Value   string `json:"value,omitempty"`
	Display string `json:"display,omitempty"`
	Type    string `json:"type,omitempty"`
	Primary bool   `json:"primary,omitempty"`
}

type UserX509Certificate struct {
	Value   string `json:"value,omitempty"`
	Display string `json:"display,omitempty"`
	Type    string `json:"type,omitempty"`
	Primary bool   `json:"primary,omitempty"`
}

type UserRecord struct {
	BaseRecord
	UserName          string                 `json:"userName,omitempty"`
	Name              UserName               `json:"name,omitempty"`
	DisplayName       string                 `json:"displayName,omitempty"`
	NickName          string                 `json:"nickName,omitempty"`
	ProfileUrl        string                 `json:"profileUrl,omitempty"`
	Title             string                 `json:"title,omitempty"`
	UserType          string                 `json:"userType,omitempty"`
	PreferredLanguage string                 `json:"preferredLanguage,omitempty"`
	Locale            string                 `json:"locale,omitempty"`
	Timezone          string                 `json:"timezone,omitempty"`
	Active            bool                   `json:"active,omitempty"`
	Password          string                 `json:"password,omitempty"`
	Emails            []UserEmail            `json:"emails,omitempty"`
	PhoneNumbers      []UserPhoneNumber      `json:"phoneNumbers,omitempty"`
	Ims               []UserIm               `json:"ims,omitempty"`
	Photos            []UserPhoto            `json:"photos,omitempty"`
	Addresses         []UserAddress          `json:"addresses,omitempty"`
	Groups            []UserGroup            `json:"groups,omitempty"`
	Entitlements      []UserEntitlement      `json:"entitlements,omitempty"`
	Roles             []UserRole             `json:"roles,omitempty"`
	X509Certificates  []UserX509Certificate  `json:"x509Certificates,omitempty"`
	Extensions        map[string]interface{} `json:"extensions,omitempty"`
}

type UserPatchOp struct {
	Operation PatchOp    `json:"op,omitempty"`
	Path      string     `json:"path,omitempty"`
	Value     UserRecord `json:"value,omitempty"`
}

type UserPatchRequest struct {
	Schemas    []string `json:"schemas,omitempty"`
	Operations []UserPatchOp
}

type GroupMember struct {
	Type    string `json:"type,omitempty"`
	Display string `json:"display,omitempty"`
	Value   string `json:"value,omitempty"`
	Ref     string `json:"$ref,omitempty"`
}

type GroupRecord struct {
	DisplayName string                 `json:"displayName,omitempty"`
	Members     []GroupMember          `json:"members,omitempty"`
	Extensions  map[string]interface{} `json:"extensions,omitempty"`
}

type GroupPatchOp struct {
	Operation PatchOp     `json:"op,omitempty"`
	Path      string      `json:"path,omitempty"`
	Value     GroupRecord `json:"value,omitempty"`
}

type GroupPatchRequest struct {
	Schemas    []string `json:"schemas,omitempty"`
	Operations []GroupPatchOp
}

type SearchRequest struct {
	Schemas            []string `json:"schemas,omitempty"`
	Attributes         []string `json:"attributes,omitempty"`
	ExcludedAttributes []string `json:"excludedAttributes,omitempty"`
	Filter             string   `json:"filter,omitempty"`
	SortBy             string   `json:"sortBy,omitempty"`
	SortOrder          string   `json:"sortOrder,omitempty"`
	StartIndex         int      `json:"startIndex,omitempty"`
	Count              int      `json:"count,omitempty"`
}

type AuthenticationScheme struct {
	Type             string `json:"type,omitempty"`
	Name             string `json:"name,omitempty"`
	Description      string `json:"description,omitempty"`
	SpecUri          string `json:"specUri,omitempty"`
	DocumentationUri string `json:"documentationUri,omitempty"`
	Primary          string `json:"primary,omitempty"`
}

type Filter struct {
	Supported  bool `json:"supported,omitempty"`
	MaxResults int  `json:"maxResults,omitempty"`
}

type Bulk struct {
	Supported      bool `json:"supported,omitempty"`
	MaxOperations  int  `json:"maxOperations,omitempty"`
	MaxPayloadSize int  `json:"maxPayloadSize,omitempty"`
}

type Supported struct {
	Supported bool `json:"supported,omitempty"`
}

type SpConfig struct {
	BaseRecord
	DocumentationUri      string                 `json:"documentationUri,omitempty"`
	Patch                 Supported              `json:"patch,omitempty"`
	Bulk                  Bulk                   `json:"bulk,omitempty"`
	Filter                Filter                 `json:"filter,omitempty"`
	Etag                  Supported              `json:"etag,omitempty"`
	ChangePassword        Supported              `json:"changePassword,omitempty"`
	Sort                  Supported              `json:"sort,omitempty"`
	AuthenticationSchemes []AuthenticationScheme `json:"authenticationSchemes,omitempty"`
}

type SchemaExt struct {
	Schema   string `json:"schema,omitempty"`
	Required bool   `json:"required,omitempty"`
}

type ResourceType struct {
	BaseRecord
	Name             string      `json:"name,omitempty"`
	Description      string      `json:"description,omitempty"`
	Endpoint         string      `json:"endpoint,omitempty"`
	Schema           string      `json:"schema,omitempty"`
	SchemaExtensions []SchemaExt `json:"schemaExtensions,omitempty"`
}
