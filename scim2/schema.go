package scim2

type BaseRecord struct {
	Id         string   `json:"id"`
	ExternalId string   `json:"externalId"`
	Meta       string   `json:"meta"`
	Schemas    []string `json:"schemas"`
}

type Attribute struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	AttrType    AttributeType `json:"type"`

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

func NewAttribute() *Attribute {
	attr := &Attribute{
		AttrType:   AttributeType۰STRING,
		Mutability: Mutability۰READ_WRITE,
		Returned:   ReturnedType۰DEFAULT,
		Uniqueness: UniquenessType۰NONE,
	}

	return attr
}
