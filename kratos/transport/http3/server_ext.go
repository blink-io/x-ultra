package http3

import (
	"github.com/blink-io/x/kratos/internal/matcher"
)

type MiddlewareMatcher = matcher.Matcher

func (s *Server) DecodeVars() DecodeRequestFunc {
	return s.decVars
}

func (s *Server) DecodeQuery() DecodeRequestFunc {
	return s.decQuery
}

func (s *Server) DecodeBody() DecodeRequestFunc {
	return s.decBody
}

func (s *Server) EncodeResponse() EncodeResponseFunc {
	return s.encResp
}

func (s *Server) EncodeError() EncodeErrorFunc {
	return s.encErr
}

func (s *Server) Middleware() MiddlewareMatcher {
	return s.middleware
}
