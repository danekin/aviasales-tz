package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/valyala/fasthttp"
)

type CheckWordUsecase interface {
	Do(word string) []string
}

type CheckWordHandler struct {
	checkWordUsecase CheckWordUsecase
}

func NewCheckWordHandler(checkWordHandler CheckWordUsecase) *CheckWordHandler {
	return &CheckWordHandler{
		checkWordUsecase: checkWordHandler,
	}
}

func (h *CheckWordHandler) Handle(ctx *fasthttp.RequestCtx) {
	word := string(ctx.QueryArgs().Peek("word"))
	if len(word) == 0 {
		ctx.Error("you must specify anagram", http.StatusBadRequest)

		return
	}

	anagrams := h.checkWordUsecase.Do(word)

	body, err := json.Marshal(anagrams)
	if err != nil {
		ctx.Error(err.Error(), http.StatusInternalServerError)

		return
	}

	ctx.SetBody(body)
}
