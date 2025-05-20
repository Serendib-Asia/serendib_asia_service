package handler

import (
	"fmt"
	"os"
	textTemplate "text/template"

	"github.com/chazool/serendib_asia_service/pkg/log"

	"github.com/chazool/serendib_asia_service/pkg/utils/constant"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// HandleStaticDoc serves static documentation files.
func HandleStaticDoc(ctx *fiber.Ctx) (err error) {
	log.Logger.Info(log.TraceMsgFuncStart(HandleStaticDocMethod))
	defer log.Logger.Info(log.TraceMsgFuncEnd(HandleStaticDocMethod))

	html, err := os.ReadFile(dotFwdSlash + staticDocName)
	if err != nil {
		log.Logger.Debug(fmt.Sprintf(constant.FileReadError, staticDocName), zap.Error(err))
	}

	template, err := textTemplate.New(staticDocName).Parse(string(html))
	if err != nil {
		log.Logger.Debug(constant.ErrorOccurredWhenTemplateParse, zap.Error(err))
	}

	ctx.Type(htmlString)

	return template.Execute(ctx, nil)
}

// HandleDoc serves documentation files.
func HandleDoc(ctx *fiber.Ctx) (err error) {
	html, err := os.ReadFile(dotFwdSlash + docFileName)
	if err != nil {
		log.Logger.Debug(fmt.Sprintf(constant.FileReadError, dotFwdSlash+docFileName), zap.Error(err))
	}

	template, err := textTemplate.New(docFileName).Parse(string(html))
	if err != nil {
		log.Logger.Debug(constant.ErrorOccurredWhenTemplateParse, zap.Error(err))
	}

	ctx.Type(htmlString)

	return template.Execute(ctx, nil)
}
