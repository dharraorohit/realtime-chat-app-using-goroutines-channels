package srvhandlers

import (
	"net/http"

	"github.com/dharraorohit/realtime-chat-app-using-goroutines-channels/utils/responseUtils"
)

type BaseHandler struct{}

func (b *BaseHandler) Get(r *http.Request) responseUtils.StandardResponse {
	return responseUtils.ResponseNotImplemented()
}

func (b *BaseHandler) Put(r *http.Request) responseUtils.StandardResponse {
	return responseUtils.ResponseNotImplemented()
}

func (b *BaseHandler) Post(r *http.Request) responseUtils.StandardResponse {
	return responseUtils.ResponseNotImplemented()
}

func (b *BaseHandler) Delete(r *http.Request) responseUtils.StandardResponse {
	return responseUtils.ResponseNotImplemented()
}

func (b *BaseHandler) Patch(r *http.Request) responseUtils.StandardResponse {
	return responseUtils.ResponseNotImplemented()
}
