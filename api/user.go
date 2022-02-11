package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	db "github.com/yimikao/reddit-clone/db/sqlc"
	"github.com/yimikao/reddit-clone/util"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required, alphanum"`
	Password string `json:"password" binding:"required, min=6"`
	Email    string `json:"email" binding:"required, email"`
}

type userResponse struct {
	Username          string    `json:"username"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

type getUserRequest struct {
	ID int64 `uri:"id" binding:"required, min=1"`
}

func newUserResponse(u *db.User) userResponse {
	return userResponse{
		Username:          u.Username,
		Email:             u.Email,
		PasswordChangedAt: u.PasswordChangedAt,
		CreatedAt:         u.CreatedAt,
	}
}

func (s *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashed_pw, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	args := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashed_pw,
		Email:          req.Email,
	}

	u, err := s.store.CreateUser(ctx, args)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, newUserResponse(&u))
}

func (s *Server) getUser(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	u, err := s.store.GetUser(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, userResponse{
		Username:          u.Username,
		Email:             u.Email,
		PasswordChangedAt: u.PasswordChangedAt,
		CreatedAt:         u.CreatedAt,
	})
}

func (s *Server) listUsers(ctx *gin.Context) {
	args := db.ListUsersParams{
		Limit:  5,
		Offset: 5,
	}

	u, err := s.store.ListUsers(ctx, args)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	var ur []userResponse
	for _, usr := range u {
		ur = append(ur, newUserResponse(&usr))
	}

	ctx.JSON(http.StatusOK, ur)
}
