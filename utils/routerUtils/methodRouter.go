package routerUtils

import (
	"net/http"

	"github.com/dharraorohit/scalable-chat-app/utils/responseUtils"
)

type ApiHandler interface {
	Get(r *http.Request) responseUtils.StandardResponse
	Put(r *http.Request) responseUtils.StandardResponse
	Post(r *http.Request) responseUtils.StandardResponse
	Delete(r *http.Request) responseUtils.StandardResponse
	Patch(r *http.Request) responseUtils.StandardResponse
}

func RouteMethod(a ApiHandler, r *http.Request) responseUtils.StandardResponse {
	switch r.Method {
	case "GET":
		return a.Get(r)
	case "PUT":
		return a.Put(r)
	case "POST":
		return a.Post(r)
	case "DELETE":
		return a.Delete(r)
	case "PATCH":
		return a.Patch(r)
	default:
		return responseUtils.StandardResponse{
			Message:      "Method not implemented",
			Code:         http.StatusMethodNotAllowed,
			ResponseData: nil,
		}
	}
}
