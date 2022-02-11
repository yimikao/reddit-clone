package api

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	db "github.com/yimikao/reddit-clone/db/sqlc"
	"github.com/yimikao/reddit-clone/util"
)

func newTestServer(t *testing.T, s db.Store) *Server {
	c := util.Config{
		TokenSecretKey:      "secret",
		AccessTokenDuration: time.Minute * 5,
	}

	sv, err := NewServer(c, s)
	require.NoError(t, err)

	return sv
}
