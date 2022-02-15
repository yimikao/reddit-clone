package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/yimikao/reddit-clone/db/sqlc"
)

type createSubRequest struct {
	CreatorID int64  `json:"creator_id"`
	Name      string `json:"name"`
}

type getSubRequest struct {
	// ID int64 `uri:"id" binding:"required,min=1`
	Name string `uri:"name" binding:"required,alphanum"`
}

func (s *Server) createSub(ctx *gin.Context) {
	var r createSubRequest

	if err := ctx.ShouldBindJSON(&r); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	p, err := s.store.CreateSub(ctx, db.CreateSubParams{
		CreatorID: r.CreatorID,
		Name:      r.Name,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, p)

}

func (s *Server) getSub(ctx *gin.Context) {
	var r getSubRequest
	if err := ctx.ShouldBindUri(&r); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	p, err := s.store.GetSub(ctx, r.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, p)
}
