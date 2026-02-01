package handlers

import (
	"context"
	"encoding/json"
	"green-api-test-assignment/internal/utils"
	"time"

	"green-api-test-assignment/internal/client"
	"green-api-test-assignment/internal/models/greenapi"

	"github.com/valyala/fasthttp"
)

func GetStateInstanceHandler(ctx *fasthttp.RequestCtx) {
	var req greenapi.UserInstanceRequest
	if err := json.Unmarshal(ctx.PostBody(), &req); err != nil {
		utils.WriteError(ctx, err, 400)
		return
	}
	reqCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	greenClient := client.New(req.IdInstance, req.ApiToken)
	result, err := greenClient.GetStateInstance(reqCtx)
	if err != nil {
		utils.WriteError(ctx, err, 500)
		return
	}

	utils.WriteJSON(ctx, result)
}

func GetSettingsHandler(ctx *fasthttp.RequestCtx) {
	var req greenapi.UserInstanceRequest
	if err := json.Unmarshal(ctx.PostBody(), &req); err != nil {
		utils.WriteError(ctx, err, 400)
		return
	}
	reqCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	greenClient := client.New(req.IdInstance, req.ApiToken)
	result, err := greenClient.GetSettings(reqCtx)
	if err != nil {
		utils.WriteError(ctx, err, 500)
		return
	}
	utils.WriteJSON(ctx, result)
}

func SendMessageHandler(ctx *fasthttp.RequestCtx) {
	var req struct {
		IdInstance string `json:"idInstance"`
		ApiToken   string `json:"apiToken"`
		ChatId     string `json:"chatId"`
		Message    string `json:"message"`
	}

	if err := json.Unmarshal(ctx.PostBody(), &req); err != nil {
		utils.WriteError(ctx, err, 400)
		return
	}

	reqCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	greenClient := client.New(req.IdInstance, req.ApiToken)
	msgReq := greenapi.SendMessageRequest{
		ChatIdOrNumber: req.ChatId,
		Message:        req.Message,
	}

	result, err := greenClient.SendMessage(reqCtx, msgReq)
	if err != nil {
		utils.WriteError(ctx, err, 500)
		return
	}
	utils.WriteJSON(ctx, result)
}

func SendFileByUrlHandler(ctx *fasthttp.RequestCtx) {
	var req struct {
		IdInstance string `json:"idInstance"`
		ApiToken   string `json:"apiToken"`
		ChatId     string `json:"chatId"`
		UrlFile    string `json:"urlFile"`
		FileName   string `json:"fileName"`
		Caption    string `json:"caption"`
	}

	if err := json.Unmarshal(ctx.PostBody(), &req); err != nil {
		utils.WriteError(ctx, err, 400)
		return
	}

	reqCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	greenClient := client.New(req.IdInstance, req.ApiToken)
	fileReq := greenapi.SendFileRequest{
		ChatIdOrNumber: req.ChatId,
		UrlFile:        req.UrlFile,
		FileName:       req.FileName,
		Caption:        req.Caption,
	}

	result, err := greenClient.SendFileByUrl(reqCtx, fileReq)
	if err != nil {
		utils.WriteError(ctx, err, 500)
		return
	}
	utils.WriteJSON(ctx, result)
}
