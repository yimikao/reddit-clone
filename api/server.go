package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/yimikao/reddit-clone/db/sqlc"
	"github.com/yimikao/reddit-clone/token"
	"github.com/yimikao/reddit-clone/util"
)

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(cfg util.Config, s db.Store) (server *Server, err error) {
	tm, err := token.NewJWTMaker(cfg.TokenSecretKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server = &Server{
		config:     cfg,
		store:      s,
		tokenMaker: tm,
	}

	server.setupRouter()
	return
}

func (s *Server) status(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "server working well",
	})
	return
}

func (s *Server) Start(addr string) error {
	return s.router.Run(addr)
}

func (s *Server) setupRouter() {
	r := gin.Default()

	r.GET("/", s.status)

	s.router = r
}
