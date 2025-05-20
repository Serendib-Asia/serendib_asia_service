package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"runtime/debug"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/chazool/serendib_asia_service/pkg/custom"
	"github.com/chazool/serendib_asia_service/pkg/log"
	"github.com/chazool/serendib_asia_service/pkg/utils/constant"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// HTTPClientImplInstance is an instance implementing the HTTPClient interface.
var HTTPClientImplInstance HTTPClient

// HTTPClient represents an interface for making HTTP requests.
type HTTPClient interface {
	// HTTPRequest makes an HTTP request.
	HTTPRequest(commonLogFields []zap.Field, request Request) (Response, error)
}

// HTTPClientImpl is a real implementation of the HTTPClient interface.
type HTTPClientImpl struct{}

// NewRealHTTPClient creates a new instance of RealHTTPClient.
func NewHTTPClientUtil() *HTTPClientImpl {
	return &HTTPClientImpl{}
}

// HTTPMethods is a type for HTTP methods.
type HTTPMethods string

// HTTPMethods constants
const (
	Post HTTPMethods = "POST"
	Get  HTTPMethods = "GET"
	Put  HTTPMethods = "PUT"
)

// Request is a struct for HTTP request.
type Request struct {
	URL         string
	Method      constant.HTTPMethod
	RequestBody any
	TimeOut     time.Duration
	Headers     map[string]string
}

// Response is a struct for HTTP response.
type Response struct {
	StatusCode int
	Body       []byte
	Headers    map[string]string
}

// CallHTTPEndpoint makes an HTTP request and returns the response and any error.
func CallHTTPEndpoint(commonLogFields []zap.Field, request Request, callingMethod string) (response Response, errorResult *custom.ErrorResult) {
	logFields := append(commonLogFields, zap.Any(constant.CallingFromMethod, callingMethod))
	log.Logger.Debug(log.TraceMsgFuncStart(constant.CallHTTPEndpointMethod), logFields...)

	defer func() {
		// Panic handling
		if r := recover(); r != nil {
			log.Logger.Error(constant.PanicOccurred, append(logFields, []zap.Field{zap.String(constant.StackTrace, string(debug.Stack()))}...)...)
			errRes := custom.BuildInternalServerErrResult(constant.UnexpectedErrorCode, fmt.Sprintf(constant.UnexpectedErrorMessage, constant.CallHTTPEndpointMethod), constant.Empty)
			errorResult = &errRes
		}
		log.Logger.Debug(log.TraceMsgFuncEnd(constant.CallHTTPEndpointMethod), logFields...)
	}()

	log.Logger.Debug(log.TraceMsgBeforeInvoke(constant.HTTPRequestMethod), commonLogFields...)
	response, err := HTTPClientImplInstance.HTTPRequest(commonLogFields, request)
	log.Logger.Debug(log.TraceMsgAfterInvoke(constant.HTTPRequestMethod), commonLogFields...)

	if err != nil {
		logFields := append(commonLogFields, zap.Any(constant.ErrorNote, err))
		log.Logger.Error(constant.APIErrorResponse, logFields...)
		errRes := custom.BuildBadReqErrResult(constant.ErrHTTPRequestCode, constant.ErrHTTPRequestMsg, err.Error())
		return response, &errRes
	}

	return response, nil
}

// HTTPRequest makes an HTTP request and returns the response and any error.
func (httpClientImpl *HTTPClientImpl) HTTPRequest(commonLogFields []zap.Field, request Request) (Response, error) {
	log.Logger.Debug(log.TraceMsgFuncStart(constant.HTTPRequestMethod), commonLogFields...)

	var (
		agent    = fiber.AcquireAgent()
		response Response
	)

	defer func() {
		agent.ConnectionClose()
		log.Logger.Debug(log.TraceMsgFuncEnd(constant.HTTPRequestMethod), commonLogFields...)
	}()

	// build and parse the request to fiber agent
	err := buildAndParseRequest(commonLogFields, request, agent)
	if err != nil {
		log.Logger.Error(constant.ErrInvalidInputWhenParse, append(commonLogFields, zap.Any(constant.ErrorNote, err))...)
		return response, err
	}

	resp := fiber.AcquireResponse()
	// http cas and set the response
	statusCode, body, err := setResponseToAgent(commonLogFields, resp, agent)
	if err != nil {
		log.Logger.Error(constant.ErrInvalidInput, append(commonLogFields, zap.Any(constant.ErrorNote, err))...)
		return response, err
	}

	// get response headers from http call and assign
	responseHeader, err := getResponseHeaders(commonLogFields, resp)
	if err != nil {
		log.Logger.Error(constant.ErrUnableToReadHeader, append(commonLogFields, zap.Any(constant.ErrorNote, err))...)
		return response, err
	}

	response = Response{StatusCode: statusCode, Body: body, Headers: responseHeader}
	return response, nil
}

func buildAndParseRequest(commonLogFields []zap.Field, request Request, agent *fiber.Agent) error {
	log.Logger.Debug(log.TraceMsgFuncStart(constant.BuildAndParseRequestMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(constant.BuildAndParseRequestMethod), commonLogFields...)

	if request.RequestBody != nil {
		agent.JSON(request.RequestBody)
	}

	if request.TimeOut > constant.Zero {
		agent = agent.Timeout(request.TimeOut)
	}

	req := agent.Request()
	req.SetRequestURI(request.URL)
	req.Header.SetMethod(string(request.Method))
	for key, value := range request.Headers {
		req.Header.Set(key, value)
	}

	log.Logger.Debug(log.TraceMsgBeforeParse(constant.HTTPRequestMethod), commonLogFields...)
	if err := agent.Parse(); err != nil {
		logFields := append(commonLogFields, zap.Any(constant.ErrorNote, err))
		log.Logger.Error(constant.ErrInvalidInputWhenParse, logFields...)

		return err
	}

	return nil
}

func setResponseToAgent(commonLogFields []zap.Field, resp *fasthttp.Response, agent *fiber.Agent) (statusCode int, body []byte, err error) {
	log.Logger.Debug(log.TraceMsgFuncStart(constant.SetResponseToAgentMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(constant.SetResponseToAgentMethod), commonLogFields...)

	var errs []error
	statusCode, body, errs = agent.SetResponse(resp).Bytes()

	if errs != nil {
		logFields := append(commonLogFields, zap.Any(constant.ErrorNote, errs))
		log.Logger.Error(constant.ErrInvalidInput, logFields...)

		err = errors.New(constant.Error)
		for _, e := range errs {
			err = errors.Wrap(err, e.Error())
		}
		return statusCode, body, err
	}

	return statusCode, body, nil
}

func getResponseHeaders(commonLogFields []zap.Field, resp *fasthttp.Response) (map[string]string, error) {
	log.Logger.Debug(log.TraceMsgFuncStart(constant.GetResponseHeadersMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(constant.GetResponseHeadersMethod), commonLogFields...)

	var (
		reader         = bytes.NewReader(resp.Header.Header())
		bReader        = bufio.NewReader(reader)
		responseHeader = make(map[string]string)
	)

	for {
		line, err := bReader.ReadString(constant.NewlineBytes)
		line = strings.TrimSpace(line)
		if err != nil {
			if err != io.EOF {
				log.Logger.Error(constant.ErrUnableToReadHeader, append(commonLogFields, zap.Any(constant.ErrorNote, err))...)
				return responseHeader, err
			}
			break
		}

		if strings.Contains(line, constant.Colon) {
			data := strings.Split(line, constant.Colon)
			responseHeader[strings.TrimSpace(data[constant.Zero])] = strings.TrimSpace(data[constant.IntOne])
		}
	}

	return responseHeader, nil
}
