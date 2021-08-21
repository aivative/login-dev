package utils

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// tError data structure to pass an error data
type tError struct {
	Code    int     `json:"code"`
	Message string  `json:"message"`
	Refs    *string `json:"refs,omitempty"`
}

// TReqResp is a base data structure to pass request/response data
type TReqResp struct {
	c *gin.Context

	Kind    string      `json:"kind"`
	Success bool        `json:"success"`
	Values  interface{} `json:"values,omitempty"`
	Error   *tError     `json:"error,omitempty"`
}

// NewReqRespStruct creates TReqResp object
func NewReqRespStruct(c *gin.Context, kind string) TReqResp {
	return TReqResp{c: c, Kind: kind}
}

// Ok will pass the value then wraps gin json response
func (r *TReqResp) Ok(values interface{}) {
	r.Success = true
	r.Values = values

	r.c.JSON(http.StatusOK, *r)
	return
}

// BadReq will pass the error message and error reference then wraps gin json response
func (r *TReqResp) BadReq(message string, errRef error) {
	r.Success = false
	r.Error = &tError{
		Code:    http.StatusBadRequest,
		Message: message,
	}

	if errRef != nil {
		er := errRef.Error()
		r.Error.Refs = &er

		// Custom mongodb unauthorized
		if er == "(Unauthorized) command find requires authentication" {
			r.Error.Code = http.StatusInternalServerError
		}
	}

	r.c.JSON(r.Error.Code, r)
}

// Unauthorized will pass the error message then wraps gin json response
func (r *TReqResp) Unauthorized(message string, errRef error) {
	r.Success = false
	r.Error = &tError{
		Code:    401,
		Message: message,
	}

	if errRef != nil {
		er := errRef.Error()
		r.Error.Refs = &er

		// Custom mongodb unauthorized
		if er == "(Unauthorized) command find requires authentication" {
			r.Error.Code = http.StatusInternalServerError
		}
	}

	r.c.JSON(r.Error.Code, *r)
}

// InternalErr will pass the error message and error reference then wraps gin json response
func (r *TReqResp) InternalErr(message string, errRef error) {
	r.Success = false
	r.Error = &tError{
		Code:    http.StatusInternalServerError,
		Message: message,
	}

	if errRef != nil {
		er := errRef.Error()
		r.Error.Refs = &er
	}

	r.c.JSON(r.Error.Code, *r)
}

func ParseKindAndBody(c *gin.Context, kind string, body interface{}) error {
	var req TReqResp

	if err := c.BindJSON(&req); err != nil {
		log.Println("didn't parse correctly:", err.Error())
		return err
	}

	if req.Kind != kind {
		return errors.New("unknown kind")
	}

	byteReqValues, err := json.Marshal(req.Values)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(byteReqValues, body); err != nil {
		return err
	}

	return validator.New().Struct(body)
}
