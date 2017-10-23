package common

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// ServiceError models all potential errors generated the service call
type ServiceError interface {
	// The http status code of the error
	GetHttpStatusCode() int

	// The message of the error as sent by the service
	GetMessage() string

	// The code of the error as sent by the service
	GetCode() string
}

type servicefailure struct {
	StatusCode int
	Code       string `json:"code,omitempty"`
	Message    string `json:"message,omitempty"`
}

func newServiceFailureFromResponse(response *http.Response) error {
	var err error
	var bReader io.ReadCloser
	bReader, response.Body, err = drainBody(response.Body)
	if err != nil {
		return servicefailure{
			StatusCode: response.StatusCode,
			Code:       "BadErrorResponse",
			Message:    fmt.Sprintf("The body of the response was not readable, due to :%s", err.Error()),
		}
	}

	body, err := ioutil.ReadAll(bReader)
	if err != nil {
		return servicefailure{
			StatusCode: response.StatusCode,
			Code:       "BadErrorResponse",
			Message:    fmt.Sprintf("The body of the response was not readable, due to :%s", err.Error()),
		}
	}

	se := servicefailure{StatusCode: response.StatusCode}
	err = json.Unmarshal(body, &se)
	if err != nil {
		return servicefailure{
			StatusCode: response.StatusCode,
			Code:       "BadErrorResponse",
			Message:    fmt.Sprintf("Error while parsing failure from response: %s", err.Error()),
		}
	}
	return se
}

func (se servicefailure) Error() string {
	return fmt.Sprintf("Service error:%s. %s. http status code: %d",
		se.Code, se.Message, se.StatusCode)
}

func (se servicefailure) GetHttpStatusCode() int {
	return se.StatusCode

}

func (se servicefailure) GetMessage() string {
	return se.Message
}

func (se servicefailure) GetCode() string {
	return se.Code
}

// IsServiceError returns false if the error is not service side, othwerise true
// additionally it returns an interface representing the ServiceError
func IsServiceError(err error) (failure ServiceError, ok bool) {
	failure, ok = err.(servicefailure)
	return
}