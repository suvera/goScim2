package scim2

type MetaRecord struct {
	ResourceType string   `json:"resourceType"`
	Created      JsonDate `json:"created"`
	LastModified JsonDate `json:"lastModified"`
	Location     string   `json:"location"`
	Version      string   `json:"version"`
}

type BaseRecord struct {
	Id         string     `json:"id"`
	ExternalId string     `json:"externalId"`
	Meta       MetaRecord `json:"meta"`
	Schemas    []string   `json:"schemas"`
}

type Attribute struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Type        AttributeType `json:"type"`

	CaseExact   bool       `json:"caseExact"`
	MultiValued bool       `json:"multiValued"`
	Mutability  Mutability `json:"mutability"`

	Required   bool           `json:"required"`
	Returned   ReturnedType   `json:"returned"`
	Uniqueness UniquenessType `json:"uniqueness"`

	CanonicalValues []interface{} `json:"canonicalValues"`
	ReferenceTypes  []string      `json:"referenceTypes"`
	SubAttributes   []Attribute   `json:"subAttributes"`
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
func (a *Attribute) AddSubAttribute(value Attribute) *Attribute {
	a.SubAttributes = append(a.SubAttributes, value)
	return a
}

type Schema struct {
	BaseRecord
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Attributes  []Attribute `json:"attributes"`
}

type ListResponse struct {
	Schemas      []string      `json:"schemas"`
	TotalResults int           `json:"totalResults"`
	ItemsPerPage int           `json:"itemsPerPage"`
	StartIndex   int           `json:"startIndex"`
	Resources    []interface{} `json:"resources"`
}

// USER Record
type UserName struct {
	Formatted       string `json:"formatted"`
	FamilyName      string `json:"familyName"`
	GivenName       string `json:"givenName"`
	MiddleName      string `json:"middleName"`
	HonorificPrefix string `json:"honorificPrefix"`
	HonorificSuffix string `json:"honorificSuffix"`
}

type UserEmail struct {
	Value   string `json:"value"`
	Display string `json:"display"`
	Type    string `json:"type"`
	Primary bool   `json:"primary"`
}

type UserPhoneNumber struct {
	Value   string `json:"value"`
	Display string `json:"display"`
	Type    string `json:"type"`
	Primary bool   `json:"primary"`
}

type UserIm struct {
	Value   string `json:"value"`
	Display string `json:"display"`
	Type    string `json:"type"`
	Primary bool   `json:"primary"`
}

type UserPhoto struct {
	Value   string `json:"value"`
	Type    string `json:"type"`
	Primary bool   `json:"primary"`
}

type UserAddress struct {
	Formatted     string `json:"formatted"`
	StreetAddress string `json:"streetAddress"`
	Locality      string `json:"locality"`
	Region        string `json:"region"`
	PostalCode    string `json:"postalCode"`
	Country       string `json:"country"`
	Type          string `json:"type"`
	Primary       bool   `json:"primary"`
}

type UserGroup struct {
	Value   string `json:"value"`
	Ref     string `json:"$ref"`
	Display string `json:"display"`
	Type    string `json:"type"`
}

type UserEntitlement struct {
	Value   string `json:"value"`
	Display string `json:"display"`
	Type    string `json:"type"`
	Primary bool   `json:"primary"`
}

type UserRole struct {
	Value   string `json:"value"`
	Display string `json:"display"`
	Type    string `json:"type"`
	Primary bool   `json:"primary"`
}

type UserX509Certificate struct {
	Value   string `json:"value"`
	Display string `json:"display"`
	Type    string `json:"type"`
	Primary bool   `json:"primary"`
}

type UserRecord struct {
	BaseRecord
	UserName          string                 `json:"userName"`
	Name              UserName               `json:"name"`
	DisplayName       string                 `json:"displayName"`
	NickName          string                 `json:"nickName"`
	ProfileUrl        string                 `json:"profileUrl"`
	Title             string                 `json:"title"`
	UserType          string                 `json:"userType"`
	PreferredLanguage string                 `json:"preferredLanguage"`
	Locale            string                 `json:"locale"`
	Timezone          string                 `json:"timezone"`
	Active            bool                   `json:"active"`
	Password          string                 `json:"password"`
	Emails            []UserEmail            `json:"emails"`
	PhoneNumbers      []UserPhoneNumber      `json:"phoneNumbers"`
	Ims               []UserIm               `json:"ims"`
	Photos            []UserPhoto            `json:"photos"`
	Addresses         []UserAddress          `json:"addresses"`
	Groups            []UserGroup            `json:"groups"`
	Entitlements      []UserEntitlement      `json:"entitlements"`
	Roles             []UserRole             `json:"roles"`
	X509Certificates  []UserX509Certificate  `json:"x509Certificates"`
	Extensions        map[string]interface{} `json:"extensions"`
}

type UserPatchOp struct {
	Operation PatchOp    `json:"op"`
	Path      string     `json:"path"`
	Value     UserRecord `json:"value"`
}

type UserPatchRequest struct {
	Schemas    []string `json:"schemas"`
	Operations []UserPatchOp
}

type GroupMember struct {
	Type    string `json:"type"`
	Display string `json:"display"`
	Value   string `json:"value"`
	Ref     string `json:"$ref"`
}

type GroupRecord struct {
	DisplayName string                 `json:"displayName"`
	Members     []GroupMember          `json:"members"`
	Extensions  map[string]interface{} `json:"extensions"`
}

type GroupPatchOp struct {
	Operation PatchOp     `json:"op"`
	Path      string      `json:"path"`
	Value     GroupRecord `json:"value"`
}

type GroupPatchRequest struct {
	Schemas    []string `json:"schemas"`
	Operations []GroupPatchOp
}

type SearchRequest struct {
	Schemas            []string `json:"schemas"`
	Attributes         []string `json:"attributes"`
	ExcludedAttributes []string `json:"excludedAttributes"`
	Filter             string   `json:"filter"`
	SortBy             string   `json:"sortBy"`
	SortOrder          string   `json:"sortOrder"`
	StartIndex         int      `json:"startIndex"`
	Count              int      `json:"count"`
}
