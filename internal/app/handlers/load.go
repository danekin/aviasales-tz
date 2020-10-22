package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/valyala/fasthttp"
)

type LoadDictionaryUsecase interface {
	Do(dict []string)
}

type LoadDictionary struct {
	loadDictionaryUsecase LoadDictionaryUsecase
}

func NewLoadDictionary(loadDictionaryUsecase LoadDictionaryUsecase) *LoadDictionary {
	return &LoadDictionary{
		loadDictionaryUsecase: loadDictionaryUsecase,
	}
}

func (h *LoadDictionary) Handle(ctx *fasthttp.RequestCtx) {
	var dict []string

	if err := json.Unmarshal(ctx.Request.Body(), &dict); err != nil {
		ctx.Error(err.Error(), http.StatusBadRequest)

		return
	}

	h.loadDictionaryUsecase.Do(dict)
}
