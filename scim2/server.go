package scim2

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Error struct {
	Code     int
	Detail   string
	ScimType string
}

var NoError = Error{Code: 0}

func NewError(code int, detail string, scimType string) Error {
	return Error{Code: code, Detail: detail, ScimType: scimType}
}

func AddScim2DefaultHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/scim+json; charset=UTF-8")
}

func ErrorResponse(w http.ResponseWriter, err Error) {
	w.WriteHeader(err.Code)
	var body string = `{"schemas":["` + URN_ERROR + `"],`

	if err.ScimType == "" {
		body += fmt.Sprintf(`"detail":"%s","status":"%d"}`,
			strings.ReplaceAll(err.Detail, `"`, `\"`), err.Code)
	} else {
		body += fmt.Sprintf(`"scimType":"%s","detail":"%s","status":"%d"}`, err.ScimType,
			strings.ReplaceAll(err.Detail, `"`, `\"`), err.Code)
	}
	fmt.Fprint(w, body)
}

func BadRequest(detail string) Error {
	return NewError(http.StatusBadRequest, detail, "")
}
func NotImplemented(detail string) Error {
	return NewError(http.StatusNotImplemented, detail, "")
}

type PathIdType uint16

const (
	PathIdSpc PathIdType = iota
	PathIdSchemas
	PathIdSchema
	PathIdResoruceTypes
	PathIdResoruceType
	PathIdSearch
	PathIdBulk
	PathIdUserCreate
	PathIdUserUpdate
	PathIdUserPatch
	PathIdUserDelete
	PathIdUserRead
	PathIdUserSearch
	PathIdGroupCreate
	PathIdGroupUpdate
	PathIdGroupPatch
	PathIdGroupDelete
	PathIdGroupRead
	PathIdGroupSearch
)

type RequestCtx struct {
	Response http.ResponseWriter
	Request  *http.Request
	PathId   PathIdType
}

type ProtocolHandler interface {
	// SPC = ServiceProviderConfig
	HandleSPC(*RequestCtx) Error
	HandleSchemas(*RequestCtx) Error
	HandleSchema(*RequestCtx, string) Error
	HandleResourceTypes(*RequestCtx) Error
	HandleResourceType(*RequestCtx, string) Error
}

type BulkHandler interface {
	HandleBulk(*RequestCtx) Error
}

type SearchHandler interface {
	HandleSearch(*RequestCtx) Error
}

type UserHandler interface {
	HandleUserCreate(*RequestCtx, UserRecord) Error
	HandleUserUpdate(*RequestCtx, string, UserRecord) Error
	HandleUserPatch(*RequestCtx, string, UserPatchRequest) Error
	HandleUserDelete(*RequestCtx, string) Error
	HandleUserRead(*RequestCtx, string) Error
	HandleUserSearch(*RequestCtx, SearchRequest) Error
}
type GroupHandler interface {
	HandleGroupCreate(*RequestCtx, GroupRecord) Error
	HandleGroupUpdate(*RequestCtx, string, GroupRecord) Error
	HandleGroupPatch(*RequestCtx, string, GroupPatchRequest) Error
	HandleGroupDelete(*RequestCtx, string) Error
	HandleGroupRead(*RequestCtx, string) Error
	HandleGroupSearch(*RequestCtx, SearchRequest) Error
}

type RequestHandler interface {
	PreHandle(*RequestCtx) (bool, Error)
	PostHandle(*RequestCtx)
	ProtocolHandler
	BulkHandler
	SearchHandler
	UserHandler
	GroupHandler
}

func makeSearch(r *http.Request) (SearchRequest, Error) {
	sr := SearchRequest{
		Filter:    r.FormValue("filter"),
		SortBy:    r.FormValue("sortBy"),
		SortOrder: r.FormValue("sortOrder"),
	}

	return sr, NoError
}

var pathHandleMap = map[PathIdType]interface{}{
	PathIdSpc: func(ctx *RequestCtx, rh RequestHandler) Error {
		return rh.HandleSPC(ctx)
	},
	PathIdSchemas: func(ctx *RequestCtx, rh RequestHandler) Error {
		return rh.HandleSchemas(ctx)
	},
	PathIdSchema: func(ctx *RequestCtx, rh RequestHandler) Error {
		vars := mux.Vars(ctx.Request)
		if _, ok := vars["Id"]; ok {
			return BadRequest("Could not find Schema ID in the URL Path")
		}

		return rh.HandleSchema(ctx, vars["Id"])
	},
	PathIdResoruceTypes: func(ctx *RequestCtx, rh RequestHandler) Error {
		return rh.HandleResourceTypes(ctx)
	},
	PathIdResoruceType: func(ctx *RequestCtx, rh RequestHandler) Error {
		vars := mux.Vars(ctx.Request)
		if _, ok := vars["Id"]; ok {
			return BadRequest("Could not find ResourceType ID in the URL Path")
		}

		return rh.HandleResourceType(ctx, vars["Id"])
	},
	PathIdSearch: func(ctx *RequestCtx, rh RequestHandler) Error {
		return rh.HandleSearch(ctx)
	},
	PathIdBulk: func(ctx *RequestCtx, rh RequestHandler) Error {
		return rh.HandleBulk(ctx)
	},
	PathIdUserCreate: func(ctx *RequestCtx, rh RequestHandler) Error {
		var user UserRecord

		decoder := json.NewDecoder(ctx.Request.Body)
		err := decoder.Decode(&user)
		if err != nil {
			return BadRequest(err.Error())
		}
		return rh.HandleUserCreate(ctx, user)
	},
	PathIdUserUpdate: func(ctx *RequestCtx, rh RequestHandler) Error {
		vars := mux.Vars(ctx.Request)
		if _, ok := vars["Id"]; ok {
			return BadRequest("Could not find User ID in the URL Path")
		}

		decoder := json.NewDecoder(ctx.Request.Body)
		var user UserRecord
		err := decoder.Decode(&user)
		if err != nil {
			return BadRequest(err.Error())
		}

		return rh.HandleUserUpdate(ctx, vars["Id"], user)
	},
	PathIdUserPatch: func(ctx *RequestCtx, rh RequestHandler) Error {
		vars := mux.Vars(ctx.Request)
		if _, ok := vars["Id"]; ok {
			return BadRequest("Could not find User ID in the URL Path")
		}

		decoder := json.NewDecoder(ctx.Request.Body)
		var patch UserPatchRequest
		err := decoder.Decode(&patch)
		if err != nil {
			return BadRequest(err.Error())
		}

		return rh.HandleUserPatch(ctx, vars["Id"], patch)
	},
	PathIdUserDelete: func(ctx *RequestCtx, rh RequestHandler) Error {
		vars := mux.Vars(ctx.Request)
		if _, ok := vars["Id"]; ok {
			return BadRequest("Could not find User ID in the URL Path")
		}

		return rh.HandleUserDelete(ctx, vars["Id"])
	},
	PathIdUserRead: func(ctx *RequestCtx, rh RequestHandler) Error {
		vars := mux.Vars(ctx.Request)
		if _, ok := vars["Id"]; ok {
			return BadRequest("Could not find User ID in the URL Path")
		}

		return rh.HandleUserRead(ctx, vars["Id"])
	},
	PathIdUserSearch: func(ctx *RequestCtx, rh RequestHandler) Error {
		sr, err := makeSearch(ctx.Request)
		if err.Code != 0 {
			return err
		}
		return rh.HandleUserSearch(ctx, sr)
	},
	PathIdGroupCreate: func(ctx *RequestCtx, rh RequestHandler) Error {
		var user GroupRecord

		decoder := json.NewDecoder(ctx.Request.Body)
		err := decoder.Decode(&user)
		if err != nil {
			return BadRequest(err.Error())
		}

		return rh.HandleGroupCreate(ctx, user)
	},
	PathIdGroupUpdate: func(ctx *RequestCtx, rh RequestHandler) Error {
		vars := mux.Vars(ctx.Request)
		if _, ok := vars["Id"]; ok {
			return BadRequest("Could not find Group ID in the URL Path")
		}

		decoder := json.NewDecoder(ctx.Request.Body)
		var user GroupRecord
		err := decoder.Decode(&user)
		if err != nil {
			return BadRequest(err.Error())
		}

		return rh.HandleGroupUpdate(ctx, vars["Id"], user)
	},
	PathIdGroupPatch: func(ctx *RequestCtx, rh RequestHandler) Error {
		vars := mux.Vars(ctx.Request)
		if _, ok := vars["Id"]; ok {
			return BadRequest("Could not find Group ID in the URL Path")
		}

		decoder := json.NewDecoder(ctx.Request.Body)
		var patch GroupPatchRequest
		err := decoder.Decode(&patch)
		if err != nil {
			return BadRequest(err.Error())
		}

		return rh.HandleGroupPatch(ctx, vars["Id"], patch)
	},
	PathIdGroupDelete: func(ctx *RequestCtx, rh RequestHandler) Error {
		vars := mux.Vars(ctx.Request)
		if _, ok := vars["Id"]; ok {
			return BadRequest("Could not find Group ID in the URL Path")
		}

		return rh.HandleGroupDelete(ctx, vars["Id"])
	},
	PathIdGroupRead: func(ctx *RequestCtx, rh RequestHandler) Error {
		vars := mux.Vars(ctx.Request)
		if _, ok := vars["Id"]; ok {
			return BadRequest("Could not find Group ID in the URL Path")
		}

		return rh.HandleGroupRead(ctx, vars["Id"])
	},
	PathIdGroupSearch: func(ctx *RequestCtx, rh RequestHandler) Error {
		sr, err := makeSearch(ctx.Request)
		if err.Code != 0 {
			return err
		}
		return rh.HandleGroupSearch(ctx, sr)
	},
}

type requestServer struct {
	pathId PathIdType
	rh     RequestHandler
}

func (s *requestServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := &RequestCtx{Response: w, Request: r, PathId: s.pathId}

	log.Println("Serving SCIM2 request:", s.pathId, r.RequestURI)

	AddScim2DefaultHeaders(w)

	success, err := s.rh.PreHandle(ctx)
	if err.Code != 0 {
		log.Error(err.Detail)
		ErrorResponse(ctx.Response, err)
		return
	} else if !success {
		log.Println("Operation skipped by PreHandle call for path:", s.pathId, r.RequestURI)
		return
	}

	f := pathHandleMap[s.pathId]

	err = f.(func(ctx *RequestCtx, rh RequestHandler) Error)(ctx, s.rh)

	if err.Code != 0 {
		log.Error(err.Detail)
		ErrorResponse(ctx.Response, err)
		return
	}

	s.rh.PostHandle(ctx)
}

func newReqHandler(pathId PathIdType, rh RequestHandler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		rs := requestServer{pathId, rh}
		rs.ServeHTTP(w, r)
	}
}

func Server(pathPrefix string, r *mux.Router, rh RequestHandler) {
	log.Out = os.Stdout
	log.Formatter = &logrus.TextFormatter{}
	log.Level = logrus.InfoLevel

	pathPrefix = "/" + strings.Trim(pathPrefix, `/`)

	// 1. Schema Paths
	r.HandleFunc(pathPrefix+PATH_SP, newReqHandler(PathIdSpc, rh)).Methods(http.MethodGet)
	r.HandleFunc(pathPrefix+PATH_SCHEMAS, newReqHandler(PathIdSchemas, rh)).Methods(http.MethodGet)
	r.HandleFunc(pathPrefix+PATH_SCHEMA+"/{Id}", newReqHandler(PathIdSchema, rh)).Methods(http.MethodGet)
	r.HandleFunc(pathPrefix+PATH_RESOURCETYPES, newReqHandler(PathIdResoruceTypes, rh)).Methods(http.MethodGet)
	r.HandleFunc(pathPrefix+PATH_RESOURCETYPE+"/Id", newReqHandler(PathIdResoruceType, rh)).Methods(http.MethodGet)

	// 2. Misc Paths
	r.HandleFunc(pathPrefix+"/"+PATH_SEARCH, newReqHandler(PathIdSearch, rh)).Methods(http.MethodPost)
	r.HandleFunc(pathPrefix+PATH_BULK, newReqHandler(PathIdBulk, rh)).Methods(http.MethodPost)

	// 3.User Paths
	r.HandleFunc(pathPrefix+PATH_USERS, newReqHandler(PathIdUserCreate, rh)).Methods(http.MethodPost)
	r.HandleFunc(pathPrefix+PATH_USERS+"/{Id}", newReqHandler(PathIdUserUpdate, rh)).Methods(http.MethodPut)
	r.HandleFunc(pathPrefix+PATH_USERS+"/{Id}", newReqHandler(PathIdUserPatch, rh)).Methods(http.MethodPatch)
	r.HandleFunc(pathPrefix+PATH_USERS+"/{Id}", newReqHandler(PathIdUserDelete, rh)).Methods(http.MethodDelete)
	r.HandleFunc(pathPrefix+PATH_USERS+"/{Id}", newReqHandler(PathIdUserRead, rh)).Methods(http.MethodGet)
	r.HandleFunc(pathPrefix+PATH_USERS+"/"+PATH_SEARCH, newReqHandler(PathIdUserSearch, rh)).Methods(http.MethodGet)

	// 4. Group Paths
	r.HandleFunc(pathPrefix+PATH_GROUPS, newReqHandler(PathIdGroupCreate, rh)).Methods(http.MethodPost)
	r.HandleFunc(pathPrefix+PATH_GROUPS+"/{Id}", newReqHandler(PathIdGroupUpdate, rh)).Methods(http.MethodPut)
	r.HandleFunc(pathPrefix+PATH_GROUPS+"/{Id}", newReqHandler(PathIdGroupPatch, rh)).Methods(http.MethodPatch)
	r.HandleFunc(pathPrefix+PATH_GROUPS+"/{Id}", newReqHandler(PathIdGroupDelete, rh)).Methods(http.MethodDelete)
	r.HandleFunc(pathPrefix+PATH_GROUPS+"/{Id}", newReqHandler(PathIdGroupRead, rh)).Methods(http.MethodGet)
	r.HandleFunc(pathPrefix+PATH_GROUPS+"/"+PATH_SEARCH, newReqHandler(PathIdGroupSearch, rh)).Methods(http.MethodGet)

	log.Println("SCIM2 Path(s) Added")
}

// Default Server Handler
type DefaultHandler struct {
}

func (d DefaultHandler) HandleSPC(ctx *RequestCtx) Error {
	return NoError
}
func (d DefaultHandler) HandleSchemas(ctx *RequestCtx) Error {
	b, err := json.MarshalIndent(DefaultSchemas(), "", "  ")
	if err != nil {
		return NewError(http.StatusInternalServerError, err.Error(), "")
	}

	ctx.Response.WriteHeader(http.StatusOK)
	fmt.Fprint(ctx.Response, string(b))
	return NoError
}
func (d DefaultHandler) HandleSchema(ctx *RequestCtx, id string) Error {
	return NoError
}
func (d DefaultHandler) HandleResourceTypes(ctx *RequestCtx) Error {
	return NoError
}
func (d DefaultHandler) HandleResourceType(ctx *RequestCtx, id string) Error {
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
