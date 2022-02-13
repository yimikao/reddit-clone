package api

import (
	"fmt"

	db "github.com/yimikao/reddit-clone/db/sqlc"
	"github.com/yimikao/reddit-clone/util"
)

func randomSub(u db.User) db.Sub {
	return db.Sub{
		ID:        0,
		CreatorID: u.ID,
		Name:      fmt.Sprintf("r/%s", util.RandomString(6)),
	}
}
