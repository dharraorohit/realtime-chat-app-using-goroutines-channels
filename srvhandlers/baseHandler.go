package srvhandlers

import (
	"net/http"

	"github.com/dharraorohit/scalable-chat-app/utils/responseUtils"
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
