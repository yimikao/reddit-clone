package api

import (
	"bytes"
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

func randomPost(u db.User, s db.Sub) db.Post {
	return db.Post{
		ID:          0,
		PosterID:    u.ID,
		SubID:       u.ID,
		Title:       util.RandomString(15),
		Description: util.RandomString(30),
	}
}

func requireBodyMatchPost(t *testing.T, bd *bytes.Buffer, p db.Post) {
	var gp db.Post

	bt, err := ioutil.ReadAll(bd)
	require.NoError(t, err)

	err = json.Unmarshal(bt, &gp)
	require.NoError(t, err)
	require.Equal(t, gp, p)
}

func TestGetPostAPI(t *testing.T) {
	u, _ := randomUser(t)
	s := randomSub(u)
	p := randomPost(u, s)

	testCases := []struct {
		name          string
		buildStubs    func(s *mockdb.MockStore)
		checkResponse func(t *testing.T, r *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			buildStubs: func(s *mockdb.MockStore) {
				s.EXPECT().
					GetPost(gomock.Any(), gomock.Eq(p.ID)).
					Times(1).
					Return(p, nil)
			},
			checkResponse: func(t *testing.T, r *httptest.ResponseRecorder) {
				requireBodyMatchPost(t, r.Body, p)
				require.Equal(t, http.StatusOK, r.Code)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			c.Finish()

			st := mockdb.NewMockStore(c)
			tc.buildStubs(st)

			sv := newTestServer(t, st)
			u := fmt.Sprintf("/post/%d", p.ID)

			r, err := http.NewRequest(http.MethodGet, u, nil)
			require.NoError(t, err)
			w := httptest.NewRecorder()

			sv.router.ServeHTTP(w, r)
			tc.checkResponse(t, w)
		})
	}

}
