package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/yimikao/reddit-clone/db/sqlc"
)

type createPostReq struct {
	PosterID    int64  `json:"poster_id"`
	SubID       int64  `json:"sub_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (s *Server) createPost(ctx *gin.Context) {
	var req createPostReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	p, err := s.store.CreatePost(ctx, db.CreatePostParams{
		PosterID:    req.PosterID,
		SubID:       req.SubID,
		Title:       req.Title,
		Description: req.Description,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, p)
}
