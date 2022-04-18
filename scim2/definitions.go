package scim2

/**
 * Group Definitions
 */

func DefaultSpConfig(ctx *RequestCtx) *SpConfig {
	r := &SpConfig{}
	r.Schemas = []string{URN_SP_CONFIG}
	r.Meta = MetaRecord{
		ResourceType: NAME_SP_CONFIG,
	}
	r.Patch = Supported{false}
	r.Bulk = Bulk{Supported: false, MaxOperations: 0, MaxPayloadSize: 0}
	r.Filter = Filter{Supported: false, MaxResults: 0}
	r.Etag = Supported{false}
	r.ChangePassword = Supported{false}
	r.Sort = Supported{false}
	r.AuthenticationSchemes = []AuthenticationScheme{}

	return r
}

func DefaultResourceTypes(ctx *RequestCtx) *ListResponse {
	r := &ListResponse{
		TotalResults: 5,
		ItemsPerPage: 10,
		StartIndex:   1,
		Schemas:      []string{URN_LIST_RESPONSE},
	}

	r.Resources = append(r.Resources, initResourceTypeSpc(ctx))
	r.Resources = append(r.Resources, initResourceTypeResourceType(ctx))
	r.Resources = append(r.Resources, initResourceTypeSchema(ctx))
	r.Resources = append(r.Resources, initResourceTypeUser(ctx))
	r.Resources = append(r.Resources, initResourceTypeGroup(ctx))

	return r
}

func DefaultSchemas(ctx *RequestCtx) *ListResponse {

	r := &ListResponse{
		TotalResults: 6,
		ItemsPerPage: 10,
		StartIndex:   1,
		Schemas:      []string{URN_LIST_RESPONSE},
	}

	r.Resources = append(r.Resources, schemaSpc(ctx))
	r.Resources = append(r.Resources, schemaResourceType(ctx))
	r.Resources = append(r.Resources, schemaSchema(ctx))
	r.Resources = append(r.Resources, schemaUser(ctx))
	r.Resources = append(r.Resources, schemaGroup(ctx))
	r.Resources = append(r.Resources, schemaEntpUser(ctx))

	return r
}
func initResourceTypeResourceType(ctx *RequestCtx) *ResourceType {
	rt := &ResourceType{}
	rt.Schemas = []string{URN_RESOURCE_TYPE}
	rt.Schema = URN_RESOURCE_TYPE
	rt.Id = NAME_RESOURCETYPE
	rt.Name = NAME_RESOURCETYPE
	rt.Description = "Resource Type"
	rt.Endpoint = PATH_RESOURCETYPES
	rt.Meta = MetaRecord{
		ResourceType: NAME_RESOURCETYPE,
		Location:     ctx.ScimBaseUrl() + PATH_RESOURCETYPES + PATH_RESOURCETYPE,
	}

	return rt
}

func initResourceTypeSpc(ctx *RequestCtx) *ResourceType {
	rt := &ResourceType{}
	rt.Schemas = []string{URN_RESOURCE_TYPE}
	rt.Schema = URN_SP_CONFIG
	rt.Id = NAME_SP_CONFIG
	rt.Name = NAME_SP_CONFIG
	rt.Description = "Service Provider Configuration"
	rt.Endpoint = PATH_SP
	rt.Meta = MetaRecord{
		ResourceType: NAME_RESOURCETYPE,
		Location:     ctx.ScimBaseUrl() + PATH_RESOURCETYPES + PATH_SP,
	}

	return rt
}

func initResourceTypeSchema(ctx *RequestCtx) *ResourceType {
	rt := &ResourceType{}
	rt.Schemas = []string{URN_RESOURCE_TYPE}
	rt.Schema = URN_SCHEMA
	rt.Id = NAME_SCHEMA
	rt.Name = NAME_SCHEMA
	rt.Description = NAME_SCHEMAS
	rt.Endpoint = PATH_SCHEMAS
	rt.Meta = MetaRecord{
		ResourceType: NAME_RESOURCETYPE,
		Location:     ctx.ScimBaseUrl() + PATH_RESOURCETYPES + PATH_SCHEMA,
	}

	return rt
}

func initResourceTypeUser(ctx *RequestCtx) *ResourceType {
	rt := &ResourceType{}
	rt.Schemas = []string{URN_RESOURCE_TYPE}
	rt.Schema = URN_USER
	rt.Id = NAME_USER
	rt.Name = NAME_USER
	rt.Description = NAME_USER + "s"
	rt.Endpoint = PATH_USERS
	rt.Meta = MetaRecord{
		ResourceType: NAME_RESOURCETYPE,
		Location:     ctx.ScimBaseUrl() + PATH_RESOURCETYPES + PATH_USER,
	}

	return rt
}

func initResourceTypeGroup(ctx *RequestCtx) *ResourceType {
	rt := &ResourceType{}
	rt.Schemas = []string{URN_RESOURCE_TYPE}
	rt.Schema = URN_GROUP
	rt.Id = NAME_GROUP
	rt.Name = NAME_GROUP
	rt.Description = NAME_GROUP + "s"
	rt.Endpoint = PATH_GROUPS
	rt.Meta = MetaRecord{
		ResourceType: NAME_RESOURCETYPE,
		Location:     ctx.ScimBaseUrl() + PATH_RESOURCETYPES + PATH_GROUP,
	}

	return rt
}

func initSchema(id string, ctx *RequestCtx) *Schema {
	s := &Schema{}
	s.Id = id
	s.Schemas = append(s.Schemas, URN_SCHEMA)
	s.Meta = MetaRecord{
		ResourceType: NAME_SCHEMA,
		Location:     ctx.ScimBaseUrl() + PATH_SCHEMAS + "/" + id,
		Created:      toJsonDate(1603095350),
		LastModified: toJsonDate(1603095350),
		Version:      `W/"123"`,
	}

	return s
}

func schemaSchema(ctx *RequestCtx) *Schema {
	s := initSchema(URN_SCHEMA, ctx)
	s.Name = NAME_SCHEMA
	s.Description = NAME_SCHEMAS
	s.Attributes = buildSchemaSchemaAttributes()

	return s
}

func schemaGroup(ctx *RequestCtx) *Schema {
	s := initSchema(URN_GROUP, ctx)
	s.Name = NAME_GROUP
	s.Description = NAME_GROUP + "s"
	s.Attributes = buildGroupSchemaAttributes()

	return s
}

func schemaUser(ctx *RequestCtx) *Schema {
	s := initSchema(URN_USER, ctx)
	s.Name = NAME_USER
	s.Description = NAME_USER + "s"
	s.Attributes = buildUserSchemaAttributes()

	return s
}

func schemaResourceType(ctx *RequestCtx) *Schema {
	s := initSchema(URN_RESOURCE_TYPE, ctx)
	s.Name = NAME_RESOURCETYPE
	s.Description = NAME_RESOURCETYPES
	s.Attributes = buildResourceSchemaAttributes()

	return s
}

func schemaSpc(ctx *RequestCtx) *Schema {
	s := initSchema(URN_SP_CONFIG, ctx)
	s.Name = NAME_SP_CONFIG
	s.Description = "Service Provider Configuration"
	s.Attributes = buildSpcSchemaAttributes()

	return s
}

func schemaEntpUser(ctx *RequestCtx) *Schema {
	s := initSchema(URN_ENTERPRISE_USER, ctx)
	s.Name = "EnterpriseUser"
	s.Description = "Enterprise Users"
	s.Attributes = buildEntpUserSchemaAttributes()

	return s
}

func buildUserSchemaAttributes() []*Attribute {
	var attrs []*Attribute

	attrs = append(attrs, NewAttribute("userName").
		SetDescription("Unique identifier for the User, typically used by the user to directly authenticate to the service provider. Each User MUST include a non-empty userName value. This identifier MUST be unique across the service provider's entire set of Users. REQUIRED.").
		SetType("string").
		SetCaseExact(false).
		SetMultiValued(false).
		SetMutability("readWrite").
		SetRequired(true).
		SetReturned("default").
		SetUniqueness("server"))

	attrs = append(attrs, NewAttribute("name").
		SetDescription("The components of the user's real name. Providers MAY return just the full name as a single string in the formatted sub-attribute, or they MAY return just the individual component attributes using the other sub-attributes, or they MAY return both. If both variants are returned, they SHOULD be describing the same name, with the formatted name indicating how the component attributes should be combined.").
		SetType("complex").
		SetCaseExact(false).
		SetMultiValued(false).
		SetMutability("readWrite").
		SetRequired(false).
		SetReturned("default").
		SetUniqueness("none").
		AddSubAttribute(NewAttribute("formatted").
			SetDescription("The full name, including all middle names, titles, and suffixes as appropriate, formatted for display (e.g., 'Ms. Barbara J Jensen, III').").
			SetType("string").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readWrite").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none")).
		AddSubAttribute(NewAttribute("familyName").
			SetDescription("The family name of the User, or last name in most Western languages (e.g., 'Jensen' given the full name 'Ms. Barbara J Jensen, III').").
			SetType("string").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readWrite").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none")).
		AddSubAttribute(NewAttribute("givenName").
			SetDescription("The given name of the User, or first name in most Western languages (e.g., 'Barbara' given the full name 'Ms. Barbara J Jensen, III').").
			SetType("string").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readWrite").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none")).
		AddSubAttribute(NewAttribute("middleName").
			SetDescription("The middle name(s) of the User (e.g., 'Jane' given the full name 'Ms. Barbara J Jensen, III').").
			SetType("string").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readWrite").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none")).
		AddSubAttribute(NewAttribute("honorificPrefix").
			SetDescription("The honorific prefix(es) of the User, or title in most Western languages (e.g., 'Ms.' given the full name 'Ms. Barbara J Jensen, III').").
			SetType("string").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readWrite").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none")).
		AddSubAttribute(NewAttribute("honorificSuffix").
			SetDescription("The honorific suffix(es) of the User, or suffix in most Western languages (e.g., 'III' given the full name 'Ms. Barbara J Jensen, III').").
			SetType("string").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readWrite").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none")))

	attrs = append(attrs, NewAttribute("displayName").
		SetDescription("The name of the User, suitable for display to end-users. The name SHOULD be the full name of the User being described, if known.").
		SetType("string").
		SetCaseExact(false).
		SetMultiValued(false).
		SetMutability("readWrite").
		SetRequired(false).
		SetReturned("default").
		SetUniqueness("none"))

	attrs = append(attrs, NewAttribute("nickName").
		SetDescription("The casual way to address the user in real life, e.g., 'Bob' or 'Bobby' instead of 'Robert'. This attribute SHOULD NOT be used to represent a User's username (e.g., 'bjensen' or 'mpepperidge').").
		SetType("string").
		SetCaseExact(false).
		SetMultiValued(false).
		SetMutability("readWrite").
		SetRequired(false).
		SetReturned("default").
		SetUniqueness("none"))

	attrs = append(attrs, NewAttribute("profileUrl").
		SetDescription("A fully qualified URL pointing to a page representing the User's online profile.").
		SetType("reference").
		SetCaseExact(false).
		SetMultiValued(false).
		SetMutability("readWrite").
		SetRequired(false).
		SetReturned("default").
		SetUniqueness("none"))

	attrs = append(attrs, NewAttribute("title").
		SetDescription("The user's title, such as \"Vice President\".").
		SetType("string").
		SetCaseExact(false).
		SetMultiValued(false).
		SetMutability("readWrite").
		SetRequired(false).
		SetReturned("default").
		SetUniqueness("none"))

	attrs = append(attrs, NewAttribute("userType").
		SetDescription("Used to identify the relationship between the organization and the user. Typical values used might be 'Contractor', 'Employee', 'Intern', 'Temp', 'External', and 'Unknown', but any value may be used.").
		SetType("string").
		SetCaseExact(false).
		SetMultiValued(false).
		SetMutability("readWrite").
		SetRequired(false).
		SetReturned("default").
		SetUniqueness("none"))

	attrs = append(attrs, NewAttribute("preferredLanguage").
		SetDescription("Indicates the User's preferred written or spoken language. Generally used for selecting a localized user interface; e.g., 'en-US' specifies the language English and country US.").
		SetType("string").
		SetCaseExact(false).
		SetMultiValued(false).
		SetMutability("readWrite").
		SetRequired(false).
		SetReturned("default").
		SetUniqueness("none"))

	attrs = append(attrs, NewAttribute("locale").
		SetDescription("Used to indicate the User's default location for purposes of localizing items such as currency, date time format, or numerical representations.").
		SetType("string").
		SetCaseExact(false).
		SetMultiValued(false).
		SetMutability("readWrite").
		SetRequired(false).
		SetReturned("default").
		SetUniqueness("none"))

	attrs = append(attrs, NewAttribute("timezone").
		SetDescription("The User's time zone in the 'Olson' time zone database format, e.g., 'America/Los_Angeles'.").
		SetType("string").
		SetCaseExact(false).
		SetMultiValued(false).
		SetMutability("readWrite").
		SetRequired(false).
		SetReturned("default").
		SetUniqueness("none"))

	attrs = append(attrs, NewAttribute("active").
		SetDescription("A Boolean value indicating the User's administrative status.").
		SetType("boolean").
		SetCaseExact(false).
		SetMultiValued(false).
		SetMutability("readWrite").
		SetRequired(false).
		SetReturned("default").
		SetUniqueness("none"))

	attrs = append(attrs, NewAttribute("password").
		SetDescription("The User's cleartext password. This attribute is intended to be used as a means to specify an initial password when creating a new User or to reset an existing User's password.").
		SetType("string").
		SetCaseExact(false).
		SetMultiValued(false).
		SetMutability("writeOnly").
		SetRequired(false).
		SetReturned("never").
		SetUniqueness("none"))

	attrs = append(attrs, NewAttribute("emails").
		SetDescription("Email addresses for the user. The value SHOULD be canonicalized by the service provider, e.g., 'bjensen@example.com' instead of 'bjensen@EXAMPLE.COM'.").
		SetType("complex").
		SetCaseExact(false).
		SetMultiValued(true).
		SetMutability("readWrite").
		SetRequired(false).
		SetReturned("default").
		SetUniqueness("none").
		AddSubAttribute(NewAttribute("value").
			SetDescription("Email addresses for the user. The value SHOULD be canonicalized by the service provider, e.g., 'bjensen@example.com' instead of 'bjensen@EXAMPLE.COM'.").
			SetType("string").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readWrite").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none")).
		AddSubAttribute(NewAttribute("display").
			SetDescription("Display value for an email").
			SetType("string").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readOnly").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none")).
		AddSubAttribute(NewAttribute("type").
			SetDescription("A label indicating the attribute's function, e.g., 'work' or 'home'.").
			SetType("string").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readWrite").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none")).
		AddSubAttribute(NewAttribute("primary").
			SetDescription("A Boolean value indicating the 'primary' or preferred attribute value for this attribute, e.g., the preferred mailing address or primary email address. The primary attribute value 'true' MUST appear no more than once.").
			SetType("boolean").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readWrite").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none")))

	attrs = append(attrs, NewAttribute("phoneNumbers").
		SetDescription("Phone numbers for the User. The value SHOULD be canonicalized by the service provider according to the format specified in RFC 3966, e.g., 'tel:+1-201-555-0123'.").
		SetType("complex").
		SetCaseExact(false).
		SetMultiValued(true).
		SetMutability("readWrite").
		SetRequired(false).
		SetReturned("default").
		SetUniqueness("none").
		AddSubAttribute(NewAttribute("value").
			SetDescription("Phone number of the User.").
			SetType("string").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readWrite").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none")).
		AddSubAttribute(NewAttribute("display").
			SetDescription("Display value for a phone number").
			SetType("string").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readOnly").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none")).
		AddSubAttribute(NewAttribute("type").
			SetDescription("A label indicating the attribute's function, e.g., 'work', 'home', 'mobile'.").
			SetType("string").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readWrite").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none")).
		AddSubAttribute(NewAttribute("primary").
			SetDescription("A Boolean value indicating the 'primary' or preferred attribute value for this attribute, e.g., the preferred phone number or primary phone number. The primary attribute value 'true' MUST appear no more than once.").
			SetType("boolean").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readWrite").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none")))

	attrs = append(attrs, NewAttribute("ims").
		SetDescription("Instant messaging addresses for the User.").
		SetType("complex").
		SetCaseExact(false).
		SetMultiValued(true).
		SetMutability("readWrite").
		SetRequired(false).
		SetReturned("default").
		SetUniqueness("none").
		AddSubAttribute(NewAttribute("value").
			SetDescription("Instant messaging address for the User.").
			SetType("string").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readWrite").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none")).
		AddSubAttribute(NewAttribute("display").
			SetDescription("Display value for an IMS").
			SetType("string").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readOnly").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none")).
		AddSubAttribute(NewAttribute("type").
			SetDescription("A label indicating the attribute's function, e.g., 'aim', 'gtalk', 'xmpp'.").
			SetType("string").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readWrite").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none")).
		AddSubAttribute(NewAttribute("primary").
			SetDescription("A Boolean value indicating the 'primary' or preferred attribute value for this attribute, e.g., the preferred messenger or primary messenger. The primary attribute value 'true' MUST appear no more than once.").
			SetType("boolean").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readWrite").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none")))

	attrs = append(attrs, NewAttribute("photos").
		SetDescription("URLs of photos of the User.").
		SetType("complex").
		SetCaseExact(false).
		SetMultiValued(true).
		SetMutability("readWrite").
		SetRequired(false).
		SetReturned("default").
		SetUniqueness("none").
		AddSubAttribute(NewAttribute("value").
			SetDescription("URL of a photo of the User.").
			SetType("reference").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readWrite").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none")).
		AddSubAttribute(NewAttribute("type").
			SetDescription("A label indicating the attribute's function, i.e., 'photo' or 'thumbnail'.").
			SetType("string").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readWrite").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none")).
		AddSubAttribute(NewAttribute("primary").
			SetDescription("A Boolean value indicating the 'primary' or preferred attribute value for this attribute, e.g., the preferred photo or thumbnail. The primary attribute value 'true' MUST appear no more than once.").
			SetType("boolean").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readWrite").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none")))

	attrs = append(attrs, NewAttribute("addresses").
		SetDescription("A physical mailing address for this User.").
		SetType("complex").
		SetCaseExact(false).
		SetMultiValued(true).
		SetMutability("readWrite").
		SetRequired(false).
		SetReturned("default").
		SetUniqueness("none").
		AddSubAttribute(NewAttribute("formatted").
			SetDescription("The full mailing address, formatted for display or use with a mailing label. This attribute MAY contain newlines.").
			SetType("string").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readWrite").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none")).
		AddSubAttribute(NewAttribute("streetAddress").
			SetDescription("The full street address component, which may include house number, street name, P.O. box, and multi-line extended street address information. This attribute MAY contain newlines.").
			SetType("string").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readWrite").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none")).
		AddSubAttribute(NewAttribute("locality").
			SetDescription("The city or locality component.").
			SetType("string").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readWrite").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none")).
		AddSubAttribute(NewAttribute("region").
			SetDescription("The state or region component.").
			SetType("string").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readWrite").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none")).
		AddSubAttribute(NewAttribute("postalCode").
			SetDescription("The zip code or postal code component.").
			SetType("string").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readWrite").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none")).
		AddSubAttribute(NewAttribute("country").
			SetDescription("The country name component.").
			SetType("string").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readWrite").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none")).
		AddSubAttribute(NewAttribute("type").
			SetDescription("A label indicating the attribute's function, e.g., 'work' or 'home'.").
			SetType("string").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readWrite").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none")).
		AddSubAttribute(NewAttribute("primary").
			SetDescription("A Boolean value indicating the 'primary' or preferred attribute value for this attribute. The primary attribute value 'true' MUST appear no more than once.").
			SetType("boolean").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readWrite").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none")))

	attrs = append(attrs, NewAttribute("groups").
		SetDescription("A list of groups to which the user belongs, either through direct membership, through nested groups, or dynamically calculated.").
		SetType("complex").
		SetCaseExact(false).
		SetMultiValued(true).
		SetMutability("readOnly").
		SetRequired(false).
		SetReturned("default").
		SetUniqueness("none").
		AddSubAttribute(NewAttribute("value").
			SetDescription("The identifier of the User's group.").
			SetType("string").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readOnly").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none")).
		AddSubAttribute(NewAttribute("$ref").
			SetDescription("The URI of the corresponding 'Group' resource to which the user belongs.").
			SetType("reference").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readOnly").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none")).
		AddSubAttribute(NewAttribute("display").
			SetDescription("A human-readable name, primarily used for display purposes. READ-ONLY.").
			SetType("string").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readOnly").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none")).
		AddSubAttribute(NewAttribute("type").
			SetDescription("A label indicating the attribute's function, e.g., 'direct' or 'indirect'.").
			SetType("string").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readOnly").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none").
			SetCanonicalValues(strArrToIface("direct", "indirect"))))

	attrs = append(attrs, NewAttribute("entitlements").
		SetDescription("A list of entitlements for the User that represent a thing the User has.").
		SetType("complex").
		SetCaseExact(false).
		SetMultiValued(true).
		SetMutability("readWrite").
		SetRequired(false).
		SetReturned("default").
		SetUniqueness("none").
		AddSubAttribute(NewAttribute("value").
			SetDescription("The value of an entitlement.").
			SetType("string").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readWrite").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none")).
		AddSubAttribute(NewAttribute("display").
			SetDescription("The display value of an entitlement.").
			SetType("string").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readOnly").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none")).
		AddSubAttribute(NewAttribute("type").
			SetDescription("A label indicating the attribute's function.").
			SetType("string").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readWrite").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none")).
		AddSubAttribute(NewAttribute("primary").
			SetDescription("A Boolean value indicating the 'primary' or preferred attribute value for this attribute. The primary attribute value 'true' MUST appear no more than once.").
			SetType("boolean").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readWrite").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none")))

	attrs = append(attrs, NewAttribute("roles").
		SetDescription("A list of roles for the User that collectively represent who the User is, e.g., 'Student', 'Faculty'.").
		SetType("complex").
		SetCaseExact(false).
		SetMultiValued(true).
		SetMutability("readWrite").
		SetRequired(false).
		SetReturned("default").
		SetUniqueness("none").
		AddSubAttribute(NewAttribute("value").
			SetDescription("The value of a role.").
			SetType("string").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readWrite").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none")).
		AddSubAttribute(NewAttribute("display").
			SetDescription("The display value of a role.").
			SetType("string").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readOnly").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none")).
		AddSubAttribute(NewAttribute("type").
			SetDescription("A label indicating the attribute's function.").
			SetType("string").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readWrite").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none")).
		AddSubAttribute(NewAttribute("primary").
			SetDescription("A Boolean value indicating the 'primary' or preferred attribute value for this attribute. The primary attribute value 'true' MUST appear no more than once.").
			SetType("boolean").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readWrite").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none")))

	attrs = append(attrs, NewAttribute("x509Certificates").
		SetDescription("A list of certificates issued to the User.").
		SetType("complex").
		SetCaseExact(false).
		SetMultiValued(true).
		SetMutability("readWrite").
		SetRequired(false).
		SetReturned("default").
		SetUniqueness("none").
		AddSubAttribute(NewAttribute("value").
			SetDescription("The value of an X.509 certificate.").
			SetType("binary").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readWrite").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none")).
		AddSubAttribute(NewAttribute("display").
			SetDescription("The display value of a certificate.").
			SetType("string").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readWrite").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none")).
		AddSubAttribute(NewAttribute("type").
			SetDescription("A label indicating the attribute's function.").
			SetType("string").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readWrite").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none")).
		AddSubAttribute(NewAttribute("primary").
			SetDescription("A Boolean value indicating the 'primary' or preferred attribute value for this attribute. The primary attribute value 'true' MUST appear no more than once.").
			SetType("boolean").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readWrite").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none")))

	return attrs
}
func buildGroupSchemaAttributes() []*Attribute {
	var attrs []*Attribute

	attrs = append(attrs, NewAttribute("displayName").
		SetDescription("A human-readable name for the Group. REQUIRED.").
		SetType("string").
		SetCaseExact(false).
		SetMultiValued(false).
		SetMutability("readWrite").
		SetRequired(true).
		SetReturned("default").
		SetUniqueness("none"))

	attrs = append(attrs, NewAttribute("members").
		SetDescription("A list of members of the Group.").
		SetType("complex").
		SetCaseExact(false).
		SetMultiValued(true).
		SetMutability("readWrite").
		SetRequired(false).
		SetReturned("default").
		SetUniqueness("none").
		AddSubAttribute(NewAttribute("type").
			SetDescription("A label indicating the type of resource, e.g., 'User' or 'Group'.").
			SetType("string").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("immutable").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none").
			SetCanonicalValues(strArrToIface("User", "Group"))).
		AddSubAttribute(NewAttribute("display").
			SetDescription("Display name for the member").
			SetType("string").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("immutable").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none")).
		AddSubAttribute(NewAttribute("value").
			SetDescription("Identifier of the member of this Group.").
			SetType("string").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("immutable").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none")).
		AddSubAttribute(NewAttribute("$ref").
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

func buildSchemaSchemaAttributes() []*Attribute {
	var attrs []*Attribute

	attrs = append(attrs, NewAttribute("name").
		SetDescription("The schema's human-readable name. When applicable, service providers MUST specify the name, e.g., 'User'.").
		SetType("string").
		SetCaseExact(false).
		SetMultiValued(false).
		SetMutability("readOnly").
		SetRequired(false).
		SetReturned("default").
		SetUniqueness("none"))

	attrs = append(attrs, NewAttribute("description").
		SetDescription("The schema's human-readable description. When applicable, service providers MUST specify the description.").
		SetType("string").
		SetCaseExact(false).
		SetMultiValued(false).
		SetMutability("readOnly").
		SetRequired(false).
		SetReturned("default").
		SetUniqueness("none"))

	attrs = append(attrs, NewAttribute("attributes").
		SetDescription("A complex attribute that includes the attributes of a schema.").
		SetType("complex").
		SetCaseExact(false).
		SetMultiValued(true).
		SetMutability("readOnly").
		SetRequired(true).
		SetReturned("default").
		SetUniqueness("none").
		AddSubAttribute(NewAttribute("name").
			SetDescription("The attribute's name.").
			SetType("string").
			SetCaseExact(true).
			SetMultiValued(false).
			SetMutability("readOnly").
			SetRequired(true).
			SetReturned("default").
			SetUniqueness("none")).
		AddSubAttribute(NewAttribute("type").
			SetDescription("The attribute's data type. Valid values include 'string', 'complex', 'boolean', 'decimal', 'integer', 'dateTime', 'reference', 'binary'.").
			SetType("string").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readOnly").
			SetRequired(true).
			SetReturned("default").
			SetUniqueness("none").
			SetCanonicalValues(strArrToIface("string", "complex", "boolean", "decimal", "integer", "dateTime", "reference", "binary"))).
		AddSubAttribute(NewAttribute("description").
			SetDescription("A human-readable description of the attribute.").
			SetType("string").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readOnly").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none")).
		AddSubAttribute(NewAttribute("multiValued").
			SetDescription("A Boolean value indicating an attribute's plurality.").
			SetType("boolean").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readOnly").
			SetRequired(true).
			SetReturned("default").
			SetUniqueness("none")).
		AddSubAttribute(NewAttribute("required").
			SetDescription("A Boolean value indicating whether or not the attribute is required.").
			SetType("boolean").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readOnly").
			SetRequired(true).
			SetReturned("default").
			SetUniqueness("none")).
		AddSubAttribute(NewAttribute("canonicalValues").
			SetDescription("A collection of canonical values. When applicable, service providers MUST specify the canonical types, e.g., 'work', 'home'.").
			SetType("complex").
			SetCaseExact(false).
			SetMultiValued(true).
			SetMutability("readOnly").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none")).
		AddSubAttribute(NewAttribute("caseExact").
			SetDescription("A Boolean value indicating whether or not a string attribute is case sensitive.").
			SetType("boolean").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readOnly").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none")).
		AddSubAttribute(NewAttribute("mutability").
			SetDescription("Indicates whether or not an attribute is modifiable.").
			SetType("string").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readOnly").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none").
			SetCanonicalValues(strArrToIface("readOnly", "readWrite", "immutable", "writeOnly"))).
		AddSubAttribute(NewAttribute("returned").
			SetDescription("Indicates when an attribute is returned in a response (e.g., to a query).").
			SetType("string").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readOnly").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none").
			SetCanonicalValues(strArrToIface("always", "never", "default", "request"))).
		AddSubAttribute(NewAttribute("uniqueness").
			SetDescription("Indicates how unique a value must be.").
			SetType("string").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readOnly").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none").
			SetCanonicalValues(strArrToIface("none", "server", "global"))).
		AddSubAttribute(NewAttribute("referenceTypes").
			SetDescription("Used only with an attribute of type 'reference'. Specifies a SCIM resourceType that a reference attribute MAY refer to, e.g., 'User'.").
			SetType("string").
			SetCaseExact(false).
			SetMultiValued(true).
			SetMutability("readOnly").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none").
			SetCanonicalValues(strArrToIface("scim", "external", "uri", "User", "Group"))).
		AddSubAttribute(NewAttribute("subAttributes").
			SetDescription("Used to define the sub-attributes of a complex attribute.").
			SetType("complex").
			SetCaseExact(false).
			SetMultiValued(true).
			SetMutability("readOnly").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none").
			AddSubAttribute(NewAttribute("name").
				SetDescription("The attribute's name.").
				SetType("string").
				SetCaseExact(true).
				SetMultiValued(false).
				SetMutability("readOnly").
				SetRequired(true).
				SetReturned("default").
				SetUniqueness("none")).
			AddSubAttribute(NewAttribute("type").
				SetDescription("The attribute's data type. Valid values include 'string', 'complex', 'boolean', 'decimal', 'integer', 'dateTime', 'reference', 'binary'.").
				SetType("string").
				SetCaseExact(false).
				SetMultiValued(false).
				SetMutability("readOnly").
				SetRequired(true).
				SetReturned("default").
				SetUniqueness("none").
				SetCanonicalValues(strArrToIface("string", "complex", "boolean", "decimal", "integer", "dateTime", "reference", "binary"))).
			AddSubAttribute(NewAttribute("description").
				SetDescription("A human-readable description of the attribute.").
				SetType("string").
				SetCaseExact(false).
				SetMultiValued(false).
				SetMutability("readOnly").
				SetRequired(false).
				SetReturned("default").
				SetUniqueness("none")).
			AddSubAttribute(NewAttribute("multiValued").
				SetDescription("A Boolean value indicating an attribute's plurality.").
				SetType("boolean").
				SetCaseExact(false).
				SetMultiValued(false).
				SetMutability("readOnly").
				SetRequired(true).
				SetReturned("default").
				SetUniqueness("none")).
			AddSubAttribute(NewAttribute("required").
				SetDescription("A Boolean value indicating whether or not the attribute is required.").
				SetType("boolean").
				SetCaseExact(false).
				SetMultiValued(false).
				SetMutability("readOnly").
				SetRequired(true).
				SetReturned("default").
				SetUniqueness("none")).
			AddSubAttribute(NewAttribute("canonicalValues").
				SetDescription("A collection of canonical values. When applicable, service providers MUST specify the canonical types, e.g., 'work', 'home'.").
				SetType("complex").
				SetCaseExact(false).
				SetMultiValued(true).
				SetMutability("readOnly").
				SetRequired(false).
				SetReturned("default").
				SetUniqueness("none")).
			AddSubAttribute(NewAttribute("caseExact").
				SetDescription("A Boolean value indicating whether or not a string attribute is case sensitive.").
				SetType("boolean").
				SetCaseExact(false).
				SetMultiValued(false).
				SetMutability("readOnly").
				SetRequired(false).
				SetReturned("default").
				SetUniqueness("none")).
			AddSubAttribute(NewAttribute("mutability").
				SetDescription("Indicates whether or not an attribute is modifiable.").
				SetType("string").
				SetCaseExact(false).
				SetMultiValued(false).
				SetMutability("readOnly").
				SetRequired(false).
				SetReturned("default").
				SetUniqueness("none").
				SetCanonicalValues(strArrToIface("readOnly", "readWrite", "immutable", "writeOnly"))).
			AddSubAttribute(NewAttribute("returned").
				SetDescription("Indicates when an attribute is returned in a response (e.g., to a query).").
				SetType("string").
				SetCaseExact(false).
				SetMultiValued(false).
				SetMutability("readOnly").
				SetRequired(false).
				SetReturned("default").
				SetUniqueness("none").
				SetCanonicalValues(strArrToIface("always", "never", "default", "request"))).
			AddSubAttribute(NewAttribute("uniqueness").
				SetDescription("Indicates how unique a value must be.").
				SetType("string").
				SetCaseExact(false).
				SetMultiValued(false).
				SetMutability("readOnly").
				SetRequired(false).
				SetReturned("default").
				SetUniqueness("none").
				SetCanonicalValues(strArrToIface("none", "server", "global"))).
			AddSubAttribute(NewAttribute("referenceTypes").
				SetDescription("Used only with an attribute of type 'reference'. Specifies a SCIM resourceType that a reference attribute MAY refer to, e.g., 'User'.").
				SetType("string").
				SetCaseExact(false).
				SetMultiValued(true).
				SetMutability("readOnly").
				SetRequired(false).
				SetReturned("default").
				SetUniqueness("none").
				SetCanonicalValues(strArrToIface("scim", "external", "uri", "User", "Group")))))

	return attrs
}

func buildResourceSchemaAttributes() []*Attribute {
	var attrs []*Attribute

	attrs = append(attrs, NewAttribute("name").
		SetDescription("The resource type name. When applicable, service providers MUST specify the name, e.g., 'User'.").
		SetType("string").
		SetCaseExact(false).
		SetMultiValued(false).
		SetMutability("readOnly").
		SetRequired(true).
		SetReturned("default").
		SetUniqueness("none"))

	attrs = append(attrs, NewAttribute("description").
		SetDescription("The resource type's human-readable description. When applicable, service providers MUST specify the description.").
		SetType("string").
		SetCaseExact(false).
		SetMultiValued(false).
		SetMutability("readOnly").
		SetRequired(false).
		SetReturned("default").
		SetUniqueness("none"))

	attrs = append(attrs, NewAttribute("endpoint").
		SetDescription("The resource type's HTTP-addressable endpoint relative to the Base URL of the service provider, e.g., \"Users\".").
		SetType("string").
		SetCaseExact(false).
		SetMultiValued(false).
		SetMutability("readOnly").
		SetRequired(true).
		SetReturned("default").
		SetUniqueness("none"))

	attrs = append(attrs, NewAttribute("schema").
		SetDescription("The resource type's primary/base schema URI.").
		SetType("string").
		SetCaseExact(false).
		SetMultiValued(false).
		SetMutability("readOnly").
		SetRequired(true).
		SetReturned("default").
		SetUniqueness("none"))

	attrs = append(attrs, NewAttribute("schemaExtensions").
		SetDescription("A list of URIs of the resource type's schema extensions.").
		SetType("complex").
		SetCaseExact(false).
		SetMultiValued(true).
		SetMutability("readOnly").
		SetRequired(false).
		SetReturned("default").
		SetUniqueness("none").
		AddSubAttribute(NewAttribute("schema").
			SetDescription("The URI of an extended schema, e.g., \"urn:edu:2.0:Staff\". This MUST be equal to the \"id\" attribute of a \"Schema\" resource.").
			SetType("string").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readOnly").
			SetRequired(true).
			SetReturned("default").
			SetUniqueness("none")).
		AddSubAttribute(NewAttribute("required").
			SetDescription("A Boolean value that specifies whether or not the schema extension is required for the resource type. If true, a resource of this type MUST include this schema extension and also include any attributes declared as required in this schema extension. If false, a resource of this type MAY omit this schema extension.").
			SetType("boolean").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readOnly").
			SetRequired(true).
			SetReturned("default").
			SetUniqueness("none")))

	return attrs
}

func buildSpcSchemaAttributes() []*Attribute {
	var attrs []*Attribute

	attrs = append(attrs, NewAttribute("documentationUri").
		SetDescription("An HTTP-addressable URL pointing to the service provider's human-consumable help documentation.").
		SetType("reference").
		SetCaseExact(false).
		SetMultiValued(false).
		SetMutability("readOnly").
		SetRequired(false).
		SetReturned("default").
		SetUniqueness("none"))

	attrs = append(attrs, NewAttribute("patch").
		SetDescription("A complex type that specifies PATCH configuration options.").
		SetType("complex").
		SetCaseExact(false).
		SetMultiValued(false).
		SetMutability("readOnly").
		SetRequired(true).
		SetReturned("default").
		SetUniqueness("none").
		AddSubAttribute(NewAttribute("supported").
			SetDescription("A Boolean value specifying whether or not the operation is supported.").
			SetType("boolean").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readOnly").
			SetRequired(true).
			SetReturned("default").
			SetUniqueness("none")))

	attrs = append(attrs, NewAttribute("bulk").
		SetDescription("A complex type that specifies bulk configuration options.").
		SetType("complex").
		SetCaseExact(false).
		SetMultiValued(false).
		SetMutability("readOnly").
		SetRequired(true).
		SetReturned("default").
		SetUniqueness("none").
		AddSubAttribute(NewAttribute("supported").
			SetDescription("A Boolean value specifying whether or not the operation is supported.").
			SetType("boolean").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readOnly").
			SetRequired(true).
			SetReturned("default").
			SetUniqueness("none")).
		AddSubAttribute(NewAttribute("maxOperations").
			SetDescription("An integer value specifying the maximum number of operations.").
			SetType("integer").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readOnly").
			SetRequired(true).
			SetReturned("default").
			SetUniqueness("none")).
		AddSubAttribute(NewAttribute("maxPayloadSize").
			SetDescription("An integer value specifying the maximum payload size in bytes.").
			SetType("integer").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readOnly").
			SetRequired(true).
			SetReturned("default").
			SetUniqueness("none")))

	attrs = append(attrs, NewAttribute("filter").
		SetDescription("A complex type that specifies FILTER options.").
		SetType("complex").
		SetCaseExact(false).
		SetMultiValued(false).
		SetMutability("readOnly").
		SetRequired(true).
		SetReturned("default").
		SetUniqueness("none").
		AddSubAttribute(NewAttribute("supported").
			SetDescription("A Boolean value specifying whether or not the operation is supported.").
			SetType("boolean").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readOnly").
			SetRequired(true).
			SetReturned("default").
			SetUniqueness("none")).
		AddSubAttribute(NewAttribute("maxResults").
			SetDescription("An integer value specifying the maximum number of resources returned in a response.").
			SetType("integer").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readOnly").
			SetRequired(true).
			SetReturned("default").
			SetUniqueness("none")))

	attrs = append(attrs, NewAttribute("changePassword").
		SetDescription("A complex type that specifies configuration options related to changing a password.").
		SetType("complex").
		SetCaseExact(false).
		SetMultiValued(false).
		SetMutability("readOnly").
		SetRequired(true).
		SetReturned("default").
		SetUniqueness("none").
		AddSubAttribute(NewAttribute("supported").
			SetDescription("A Boolean value specifying whether or not the operation is supported.").
			SetType("boolean").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readOnly").
			SetRequired(true).
			SetReturned("default").
			SetUniqueness("none")))

	attrs = append(attrs, NewAttribute("sort").
		SetDescription("A complex type that specifies sort configuration options.").
		SetType("complex").
		SetCaseExact(false).
		SetMultiValued(false).
		SetMutability("readOnly").
		SetRequired(true).
		SetReturned("default").
		SetUniqueness("none").
		AddSubAttribute(NewAttribute("supported").
			SetDescription("A Boolean value specifying whether or not the operation is supported.").
			SetType("boolean").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readOnly").
			SetRequired(true).
			SetReturned("default").
			SetUniqueness("none")))

	attrs = append(attrs, NewAttribute("etag").
		SetDescription("A complex type that specifies etag configuration options.").
		SetType("complex").
		SetCaseExact(false).
		SetMultiValued(false).
		SetMutability("readOnly").
		SetRequired(true).
		SetReturned("default").
		SetUniqueness("none").
		AddSubAttribute(NewAttribute("supported").
			SetDescription("A Boolean value specifying whether or not the operation is supported.").
			SetType("boolean").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readOnly").
			SetRequired(true).
			SetReturned("default").
			SetUniqueness("none")))

	attrs = append(attrs, NewAttribute("authenticationSchemes").
		SetDescription("A multi-valued complex type that specifies supported authentication scheme properties.").
		SetType("complex").
		SetCaseExact(false).
		SetMultiValued(true).
		SetMutability("readOnly").
		SetRequired(true).
		SetReturned("default").
		SetUniqueness("none").
		AddSubAttribute(NewAttribute("type").
			SetDescription("Specifies the type of the authentication scheme").
			SetType("string").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readOnly").
			SetRequired(true).
			SetReturned("default").
			SetUniqueness("none").
			SetCanonicalValues(strArrToIface("oauth2", "httpbasic"))).
		AddSubAttribute(NewAttribute("name").
			SetDescription("The common authentication scheme name, e.g., HTTP Basic.").
			SetType("string").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readOnly").
			SetRequired(true).
			SetReturned("default").
			SetUniqueness("none")).
		AddSubAttribute(NewAttribute("description").
			SetDescription("A description of the authentication scheme.").
			SetType("string").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readOnly").
			SetRequired(true).
			SetReturned("default").
			SetUniqueness("none")).
		AddSubAttribute(NewAttribute("specUri").
			SetDescription("An HTTP-addressable URL pointing to the authentication scheme's specification.").
			SetType("reference").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readOnly").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none")).
		AddSubAttribute(NewAttribute("documentationUri").
			SetDescription("An HTTP-addressable URL pointing to the authentication scheme's usage documentation.").
			SetType("reference").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readOnly").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none")).
		AddSubAttribute(NewAttribute("primary").
			SetDescription("A boolean indicates whether this authentication schema is primary").
			SetType("boolean").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readOnly").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none")))

	return attrs
}

func buildEntpUserSchemaAttributes() []*Attribute {
	var attrs []*Attribute

	attrs = append(attrs, NewAttribute("employeeNumber").
		SetDescription("Numeric or alphanumeric identifier assigned to a person, typically based on order of hire or association with an organization.").
		SetType("string").
		SetCaseExact(false).
		SetMultiValued(false).
		SetMutability("readWrite").
		SetRequired(false).
		SetReturned("default").
		SetUniqueness("none"))

	attrs = append(attrs, NewAttribute("costCenter").
		SetDescription("Identifies the name of a cost center.").
		SetType("string").
		SetCaseExact(false).
		SetMultiValued(false).
		SetMutability("readWrite").
		SetRequired(false).
		SetReturned("default").
		SetUniqueness("none"))

	attrs = append(attrs, NewAttribute("organization").
		SetDescription("Identifies the name of an organization.").
		SetType("string").
		SetCaseExact(false).
		SetMultiValued(false).
		SetMutability("readWrite").
		SetRequired(false).
		SetReturned("default").
		SetUniqueness("none"))

	attrs = append(attrs, NewAttribute("division").
		SetDescription("Identifies the name of a division.").
		SetType("string").
		SetCaseExact(false).
		SetMultiValued(false).
		SetMutability("readWrite").
		SetRequired(false).
		SetReturned("default").
		SetUniqueness("none"))

	attrs = append(attrs, NewAttribute("department").
		SetDescription("Identifies the name of a department.").
		SetType("string").
		SetCaseExact(false).
		SetMultiValued(false).
		SetMutability("readWrite").
		SetRequired(false).
		SetReturned("default").
		SetUniqueness("none"))

	attrs = append(attrs, NewAttribute("manager").
		SetDescription("The User's manager.  A complex type that optionally allows service providers to represent organizational hierarchy by referencing the 'id' attribute of another User.").
		SetType("complex").
		SetCaseExact(false).
		SetMultiValued(false).
		SetMutability("readWrite").
		SetRequired(false).
		SetReturned("default").
		SetUniqueness("none").
		AddSubAttribute(NewAttribute("value").
			SetDescription("The id of the SCIM resource representing the User's manager. REQUIRED.").
			SetType("string").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readWrite").
			SetRequired(true).
			SetReturned("default").
			SetUniqueness("none")).
		AddSubAttribute(NewAttribute("$ref").
			SetDescription("The URI of the SCIM resource representing the User's manager.").
			SetType("reference").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readWrite").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none")).
		AddSubAttribute(NewAttribute("displayName").
			SetDescription("The displayName of the User's manager.").
			SetType("string").
			SetCaseExact(false).
			SetMultiValued(false).
			SetMutability("readOnly").
			SetRequired(false).
			SetReturned("default").
			SetUniqueness("none")))

	return attrs
}
