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
	Response   http.ResponseWriter
	Request    *http.Request
	PathId     PathIdType
	PathPrefix string
}

func (ctx *RequestCtx) ScimBaseUrl() string {
	protocol := "https://"
	if ctx.Request.TLS == nil {
		protocol = "http://"
	}
	if ctx.PathPrefix == "/" {
		return protocol + ctx.Request.Host
	}
	return protocol + ctx.Request.Host + ctx.PathPrefix
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
		if _, ok := vars["Id"]; !ok {
			return BadRequest("Could not find Schema ID in the URL Path")
		}

		return rh.HandleSchema(ctx, vars["Id"])
	},
	PathIdResoruceTypes: func(ctx *RequestCtx, rh RequestHandler) Error {
		return rh.HandleResourceTypes(ctx)
	},
	PathIdResoruceType: func(ctx *RequestCtx, rh RequestHandler) Error {
		vars := mux.Vars(ctx.Request)
		if _, ok := vars["Id"]; !ok {
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
		if _, ok := vars["Id"]; !ok {
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
		if _, ok := vars["Id"]; !ok {
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
		if _, ok := vars["Id"]; !ok {
			return BadRequest("Could not find User ID in the URL Path")
		}

		return rh.HandleUserDelete(ctx, vars["Id"])
	},
	PathIdUserRead: func(ctx *RequestCtx, rh RequestHandler) Error {
		vars := mux.Vars(ctx.Request)
		if _, ok := vars["Id"]; !ok {
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
		if _, ok := vars["Id"]; !ok {
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
		if _, ok := vars["Id"]; !ok {
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
		if _, ok := vars["Id"]; !ok {
			return BadRequest("Could not find Group ID in the URL Path")
		}

		return rh.HandleGroupDelete(ctx, vars["Id"])
	},
	PathIdGroupRead: func(ctx *RequestCtx, rh RequestHandler) Error {
		vars := mux.Vars(ctx.Request)
		if _, ok := vars["Id"]; !ok {
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
	pathId     PathIdType
	rh         RequestHandler
	pathPrefix string
}

func (s *requestServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := &RequestCtx{Response: w, Request: r, PathId: s.pathId, PathPrefix: s.pathPrefix}

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

func newReqHandler(
	pathId PathIdType,
	rh RequestHandler,
	pathPrefix string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		rs := requestServer{pathId: pathId, rh: rh, pathPrefix: pathPrefix}
		rs.ServeHTTP(w, r)
	}
}

func Server(pathPrefix string, r *mux.Router, rh RequestHandler) {
	r.MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		AddScim2DefaultHeaders(w)
		ErrorResponse(w, NewError(http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed), ""))
	})

	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		AddScim2DefaultHeaders(w)
		ErrorResponse(w, NewError(http.StatusNotFound, http.StatusText(http.StatusNotFound), ""))
	})

	log.Out = os.Stdout
	log.Formatter = &logrus.TextFormatter{}
	log.Level = logrus.InfoLevel

	pf := "/" + strings.Trim(pathPrefix, `/`)

	// 1. Schema Paths
	r.HandleFunc(pf+PATH_SP, newReqHandler(PathIdSpc, rh, pf)).Methods(http.MethodGet)
	r.HandleFunc(pf+PATH_SCHEMAS, newReqHandler(PathIdSchemas, rh, pf)).Methods(http.MethodGet)
	r.HandleFunc(pf+PATH_SCHEMAS+"/{Id}", newReqHandler(PathIdSchema, rh, pf)).Methods(http.MethodGet)
	r.HandleFunc(pf+PATH_RESOURCETYPES, newReqHandler(PathIdResoruceTypes, rh, pf)).Methods(http.MethodGet)
	r.HandleFunc(pf+PATH_RESOURCETYPES+"/{Id}", newReqHandler(PathIdResoruceType, rh, pf)).Methods(http.MethodGet)

	// 2. Misc Paths
	r.HandleFunc(pf+"/"+PATH_SEARCH, newReqHandler(PathIdSearch, rh, pf)).Methods(http.MethodPost)
	r.HandleFunc(pf+PATH_BULK, newReqHandler(PathIdBulk, rh, pf)).Methods(http.MethodPost)

	// 3.User Paths
	r.HandleFunc(pf+PATH_USERS, newReqHandler(PathIdUserCreate, rh, pf)).Methods(http.MethodPost)
	r.HandleFunc(pf+PATH_USERS+"/{Id}", newReqHandler(PathIdUserUpdate, rh, pf)).Methods(http.MethodPut)
	r.HandleFunc(pf+PATH_USERS+"/{Id}", newReqHandler(PathIdUserPatch, rh, pf)).Methods(http.MethodPatch)
	r.HandleFunc(pf+PATH_USERS+"/{Id}", newReqHandler(PathIdUserDelete, rh, pf)).Methods(http.MethodDelete)
	r.HandleFunc(pf+PATH_USERS+"/{Id}", newReqHandler(PathIdUserRead, rh, pf)).Methods(http.MethodGet)
	r.HandleFunc(pf+PATH_USERS+"/"+PATH_SEARCH, newReqHandler(PathIdUserSearch, rh, pf)).Methods(http.MethodGet)

	// 4. Group Paths
	r.HandleFunc(pf+PATH_GROUPS, newReqHandler(PathIdGroupCreate, rh, pf)).Methods(http.MethodPost)
	r.HandleFunc(pf+PATH_GROUPS+"/{Id}", newReqHandler(PathIdGroupUpdate, rh, pf)).Methods(http.MethodPut)
	r.HandleFunc(pf+PATH_GROUPS+"/{Id}", newReqHandler(PathIdGroupPatch, rh, pf)).Methods(http.MethodPatch)
	r.HandleFunc(pf+PATH_GROUPS+"/{Id}", newReqHandler(PathIdGroupDelete, rh, pf)).Methods(http.MethodDelete)
	r.HandleFunc(pf+PATH_GROUPS+"/{Id}", newReqHandler(PathIdGroupRead, rh, pf)).Methods(http.MethodGet)
	r.HandleFunc(pf+PATH_GROUPS+"/"+PATH_SEARCH, newReqHandler(PathIdGroupSearch, rh, pf)).Methods(http.MethodGet)

	log.Println("SCIM2 Path(s) Added")
}
