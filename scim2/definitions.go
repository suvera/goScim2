package scim2

/**
 * Group Definitions
 */

func DefaultSchemas() *ListResponse {

	r := &ListResponse{
		TotalResults: 6,
		ItemsPerPage: 10,
		StartIndex:   1,
		Schemas:      []string{URN_LIST_RESPONSE},
	}

	r.Resources = append(r.Resources, schemaSchema())

	return r
}
func initSchema(id string) *Schema {
	s := &Schema{}
	s.Id = id
	s.Schemas = append(s.Schemas, URN_SCHEMA)
	s.Meta = MetaRecord{
		ResourceType: NAME_SCHEMA,
		Location:     DEFAULT_SERVER + PATH_SCHEMAS + id,
		Created:      toJsonDate(1603095350),
		LastModified: toJsonDate(1603095350),
		Version:      `W/"123"`,
	}

	return s
}
func schemaSchema() *Schema {
	s := initSchema(URN_SCHEMA)
	s.Name = NAME_SCHEMA
	s.Description = NAME_SCHEMAS
	s.Attributes = buildGroupSchemaAttributes()

	return s
}

func buildGroupSchemaAttributes() []Attribute {
	var attrs []Attribute

	attrs = append(attrs, *NewAttribute("displayName").
		SetDescription("A human-readable name for the Group. REQUIRED.").
		SetType("string").
		SetCaseExact(false).
		SetMultiValued(false).
		SetMutability("readWrite").
		SetRequired(true).
		SetReturned("default").
		SetUniqueness("none"))

	attrs = append(attrs, *NewAttribute("members").
		SetDescription("A list of members of the Group.").
		SetType("complex").
		SetCaseExact(false).
		SetMultiValued(true).
		SetMutability("readWrite").
		SetRequired(false).
		SetReturned("default").
		SetUniqueness("none").
		AddSubAttribute(*NewAttribute("type").
			SetDescription("A label indicating the type of resource, e.g., 'User' or 'Group'.").
			SetType("string").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("immutable").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none").
			SetCanonicalValues(strArrToIface("User", "Group"))).
		AddSubAttribute(*NewAttribute("display").
			SetDescription("Display name for the member").
			SetType("string").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("immutable").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none")).
		AddSubAttribute(*NewAttribute("value").
			SetDescription("Identifier of the member of this Group.").
			SetType("string").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("immutable").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none")).
		AddSubAttribute(*NewAttribute("$ref").
			SetDescription("The URI corresponding to a SCIM resource that is a member of this Group.").
			SetType("reference").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("immutable").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none")))

	return attrs
}
