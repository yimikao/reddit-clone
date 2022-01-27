// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
	"time"
)

type Comment struct {
	ID          int64 `json:"id"`
	CommenterID int64 `json:"commenter_id"`
	// may be null for comments that are top level
	ParentCommentID sql.NullInt64 `json:"parent_comment_id"`
	PostID          int64         `json:"post_id"`
	Points          sql.NullInt64 `json:"points"`
	CreatedAt       time.Time     `json:"created_at"`
}

type Mod struct {
	ID        int64     `json:"id"`
	SubID     int64     `json:"sub_id"`
	UserID    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type Post struct {
	ID          int64     `json:"id"`
	PosterID    int64     `json:"poster_id"`
	SubID       int64     `json:"sub_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type Sub struct {
	ID        int64     `json:"id"`
	CreatorID int64     `json:"creator_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type Subscriber struct {
	ID        int64     `json:"id"`
	SubID     int64     `json:"sub_id"`
	UserID    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type User struct {
	ID                int64     `json:"id"`
	Username          string    `json:"username"`
	HashedPassword    string    `json:"hashed_password"`
	Email             string    `json:"email"`
	Avatar            string    `json:"avatar"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
	Isblocked         bool      `json:"isblocked"`
}
