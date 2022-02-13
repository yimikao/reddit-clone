package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	mockdb "github.com/yimikao/reddit-clone/db/mock"
	db "github.com/yimikao/reddit-clone/db/sqlc"

	"github.com/yimikao/reddit-clone/util"
)

func testUser(t *testing.T) (u db.User, password string) {
	password = util.RandomString(6)
	hpwd, err := util.HashPassword(password)
	require.NoError(t, err)

	u = db.User{
		ID:             0,
		Username:       util.RandomOwner(),
		HashedPassword: hpwd,
		Email:          util.RandomString(6),
	}
	return
}

func requireBodyMatchUser(t *testing.T, body *bytes.Buffer, u *db.User) {
	bytes, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var userGotten db.User
	err = json.Unmarshal(bytes, &userGotten)
	require.NoError(t, err)

	require.Equal(t, userGotten, u)
}

func TestGetUserAPI(t *testing.T) {
	u, _ := testUser(t)

	testCases := []struct {
		name          string
		userID        int64
		buildStubs    func(s *mockdb.MockStore)
		checkResponse func(t *testing.T, r *httptest.ResponseRecorder)
	}{
		{
			name:   "OK",
			userID: u.ID,
			buildStubs: func(s *mockdb.MockStore) {
				s.EXPECT().
					GetUser(gomock.Any(), gomock.Eq(u.ID)).
					Times(1).
					Return(u, nil)
			},
			checkResponse: func(t *testing.T, r *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, r.Code)
				requireBodyMatchUser(t, r.Body, &u)
			},
		},
		{
			name:   "Not Found",
			userID: u.ID,
			buildStubs: func(s *mockdb.MockStore) {
				s.EXPECT().
					GetUser(gomock.Any(), gomock.Eq(u.ID)).
					Times(1).
					Return(db.User{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, r *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, r.Code)
			},
		},
		{
			name:   "Bad Request",
			userID: u.ID,
			buildStubs: func(s *mockdb.MockStore) {
				s.EXPECT().
					GetUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, r *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, r.Code)
			},
		},
		{
			name:   "Internal Error",
			userID: u.ID,
			buildStubs: func(s *mockdb.MockStore) {
				s.EXPECT().
					GetUser(gomock.Any(), gomock.Eq(u.ID)).
					Times(1).
					Return(db.User{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, r *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, r.Code)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			url := fmt.Sprintf("/users/%d", tc.userID)
			req, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			r := httptest.NewRecorder()

			s := newTestServer(t, store)
			s.router.ServeHTTP(r, req)

			tc.checkResponse(t, r)

		})
	}

}
