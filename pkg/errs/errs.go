package errs

import (
	"errors"
	"net/http"

	"github.com/famesensor/playground-go-fiber-todonotes/pkg/reponseHandler"
	"github.com/gofiber/fiber/v2"
)

var (
	DocumentNotFound    = errors.New("NO_DATA_FOUND")
	CannotParseData     = errors.New("CANNOT_PARSE_DATA")
	InternalServerError = errors.New("INTERNAL_SERVER_ERROR")
	ValidateErrors      = errors.New("Validate_Errors")
)

var (
	Msgs_CannotParseData     string = "data cannot be parsed"
	Msgs_DocumentNotFound    string = "data not found"
	Msgs_UserAlreadyExsiting string = "User alreay exsiting"
	Msgs_ValidateErrors      string = "Validate Errors"
)

func errorDetails(err error, errData interface{}) (int, string, interface{}) {
	switch err {
	case DocumentNotFound:
		return http.StatusNotFound, DocumentNotFound.Error(), []string{Msgs_DocumentNotFound}
	case CannotParseData:
		return http.StatusBadRequest, CannotParseData.Error(), []string{Msgs_CannotParseData}
	case ValidateErrors:
		return http.StatusBadRequest, Msgs_ValidateErrors, errData
	default:
		return http.StatusInternalServerError, InternalServerError.Error(), []string{err.Error()}
	}
}

func ErrorReponse(c *fiber.Ctx, err error, errData interface{}) error {
	code, msg, data := errorDetails(err, errData)
	return reponseHandler.ReponseMsg(c, code, "failed", msg, data)
}
