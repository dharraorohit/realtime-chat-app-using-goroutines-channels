package srvhandlers

import (
	"net/http"

	"github.com/dharraorohit/scalable-chat-app/utils/responseUtils"
	"github.com/dharraorohit/scalable-chat-app/utils/routerUtils"
)

type MessageHandler struct {
	BaseHandler
}

func (m *MessageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := routerUtils.RouteMethod(m, r)
	response.RenderResponse(w)
}

func (m *MessageHandler) Get(r *http.Request) responseUtils.StandardResponse {
	return responseUtils.StandardResponse{
		Message:      "Hello World!",
		Code:         http.StatusOK,
		ResponseData: "Hello World!",
	}
}
