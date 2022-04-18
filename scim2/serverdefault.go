package scim2

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Default Server Handler
type DefaultHandler struct {
}

func (d DefaultHandler) HandleSPC(ctx *RequestCtx) Error {
	b, err := json.MarshalIndent(DefaultSpConfig(ctx), "", "  ")
	if err != nil {
		return NewError(http.StatusInternalServerError, err.Error(), "")
	}

	ctx.Response.WriteHeader(http.StatusOK)
	fmt.Fprint(ctx.Response, string(b))
	return NoError
}

func (d DefaultHandler) HandleSchemas(ctx *RequestCtx) Error {
	b, err := json.MarshalIndent(DefaultSchemas(ctx), "", "  ")
	if err != nil {
		return NewError(http.StatusInternalServerError, err.Error(), "")
	}

	ctx.Response.WriteHeader(http.StatusOK)
	fmt.Fprint(ctx.Response, string(b))
	return NoError
}

func (d DefaultHandler) HandleSchema(ctx *RequestCtx, id string) Error {
	var obj interface{}

	switch id {
	case URN_RESOURCE_TYPE:
		obj = schemaResourceType(ctx)
	case URN_SCHEMA:
		obj = schemaSchema(ctx)
	case URN_USER:
		obj = schemaUser(ctx)
	case URN_GROUP:
		obj = schemaGroup(ctx)
	case URN_SP_CONFIG:
		obj = schemaSpc(ctx)
	case URN_ENTERPRISE_USER:
		obj = schemaEntpUser(ctx)
	default:
		return BadRequest("Inavlid ResourceType Id")
	}

	b, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return NewError(http.StatusInternalServerError, err.Error(), "")
	}

	ctx.Response.WriteHeader(http.StatusOK)
	fmt.Fprint(ctx.Response, string(b))
	return NoError
}

func (d DefaultHandler) HandleResourceTypes(ctx *RequestCtx) Error {
	b, err := json.MarshalIndent(DefaultResourceTypes(ctx), "", "  ")
	if err != nil {
		return NewError(http.StatusInternalServerError, err.Error(), "")
	}

	ctx.Response.WriteHeader(http.StatusOK)
	fmt.Fprint(ctx.Response, string(b))
	return NoError
}

func (d DefaultHandler) HandleResourceType(ctx *RequestCtx, id string) Error {
	var obj interface{}

	switch id {
	case NAME_RESOURCETYPE:
		obj = initResourceTypeResourceType(ctx)
	case NAME_SCHEMA:
		obj = initResourceTypeSchema(ctx)
	case NAME_USER:
		obj = initResourceTypeUser(ctx)
	case NAME_GROUP:
		obj = initResourceTypeGroup(ctx)
	case NAME_SP_CONFIG:
		obj = initResourceTypeSpc(ctx)
	default:
		return BadRequest("Inavlid ResourceType Id")
	}

	b, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return NewError(http.StatusInternalServerError, err.Error(), "")
	}

	ctx.Response.WriteHeader(http.StatusOK)
	fmt.Fprint(ctx.Response, string(b))
	return NoError
}

func (d DefaultHandler) PreHandle(*RequestCtx) (bool, Error) {
	return true, NoError
}
func (d DefaultHandler) PostHandle(*RequestCtx) {
}

// Not Implemeneted Here, but developer must implement it in their projects!
func (d DefaultHandler) HandleBulk(ctx *RequestCtx) Error {
	return NotImplemented(http.StatusText(http.StatusNotImplemented))
}
func (d DefaultHandler) HandleSearch(ctx *RequestCtx) Error {
	return NotImplemented(http.StatusText(http.StatusNotImplemented))
}
func (d DefaultHandler) HandleUserCreate(ctx *RequestCtx, user UserRecord) Error {
	log.Println("I am here")
	return NotImplemented(http.StatusText(http.StatusNotImplemented))
}
func (d DefaultHandler) HandleUserUpdate(ctx *RequestCtx, id string, user UserRecord) Error {
	return NotImplemented(http.StatusText(http.StatusNotImplemented))
}
func (d DefaultHandler) HandleUserPatch(ctx *RequestCtx, id string, patch UserPatchRequest) Error {
	return NotImplemented(http.StatusText(http.StatusNotImplemented))
}
func (d DefaultHandler) HandleUserDelete(ctx *RequestCtx, id string) Error {
	return NotImplemented(http.StatusText(http.StatusNotImplemented))
}
func (d DefaultHandler) HandleUserRead(ctx *RequestCtx, id string) Error {
	return NotImplemented(http.StatusText(http.StatusNotImplemented))
}
func (d DefaultHandler) HandleUserSearch(ctx *RequestCtx, search SearchRequest) Error {
	return NotImplemented(http.StatusText(http.StatusNotImplemented))
}
func (d DefaultHandler) HandleGroupCreate(ctx *RequestCtx, group GroupRecord) Error {
	return NotImplemented(http.StatusText(http.StatusNotImplemented))
}
func (d DefaultHandler) HandleGroupUpdate(ctx *RequestCtx, id string, group GroupRecord) Error {
	return NotImplemented(http.StatusText(http.StatusNotImplemented))
}
func (d DefaultHandler) HandleGroupPatch(ctx *RequestCtx, id string, patch GroupPatchRequest) Error {
	return NotImplemented(http.StatusText(http.StatusNotImplemented))
}
func (d DefaultHandler) HandleGroupDelete(ctx *RequestCtx, id string) Error {
	return NotImplemented(http.StatusText(http.StatusNotImplemented))
}
func (d DefaultHandler) HandleGroupRead(ctx *RequestCtx, id string) Error {
	return NotImplemented(http.StatusText(http.StatusNotImplemented))
}
func (d DefaultHandler) HandleGroupSearch(ctx *RequestCtx, search SearchRequest) Error {
	return NotImplemented(http.StatusText(http.StatusNotImplemented))
}
