package utils

import (
	"net/url"
	"path"

	"github.com/chazool/serendib_asia_service/pkg/config"
	"github.com/chazool/serendib_asia_service/pkg/log"
	"github.com/chazool/serendib_asia_service/pkg/utils/constant"

	"go.uber.org/zap"
)

func ParseBaseURL(commonLogFields []zap.Field, baseURL string) (parsedURL *url.URL, err error) {
	log.Logger.Debug(log.TraceMsgFuncStart(constant.ParseBaseURLMethod), log.TraceMethodInputs(commonLogFields, baseURL)...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(constant.ParseBaseURLMethod), log.TraceMethodOutputWithErr(commonLogFields, parsedURL, err)...)

	parsedURL, err = url.Parse(baseURL)
	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(constant.ParseBaseURLMethod), log.TraceError(commonLogFields, err)...)
		return nil, err
	}
	return parsedURL, nil
}

func GenerateTicketURL(commonLogFields []zap.Field, ticketNumber string) (finalURL string, err error) {
	log.Logger.Debug(log.TraceMsgFuncStart(constant.GenerateTicketURLMethod), log.TraceMethodInputs(commonLogFields, ticketNumber)...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(constant.GenerateTicketURLMethod), log.TraceMethodOutputWithErr(commonLogFields, finalURL, err)...)

	var (
		baseURL *url.URL
		client  = config.GetConfig().DefaultClient
	)

	switch client {
	case constant.SDesk:
		sDeskBaseURL := config.GetConfig().SDeskTicketBaseURL
		baseURL, err = ParseBaseURL(commonLogFields, sDeskBaseURL)
	default:
		// At the moment, the default URL is also the SDesk URL
		commonBaseURL := config.GetConfig().DefaultTicketBaseURL
		baseURL, err = ParseBaseURL(commonLogFields, commonBaseURL)
	}

	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(constant.GenerateTicketURLMethod), log.TraceError(commonLogFields, err)...)
		return "", err
	}

	baseURL.Path = path.Join(baseURL.Path, ticketNumber)
	finalURL = baseURL.String()
	return finalURL, nil
}
