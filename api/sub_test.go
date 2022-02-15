package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	mockdb "github.com/yimikao/reddit-clone/db/mock"
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

func TestGetSubAPI(t *testing.T) {
	u, _ := randomUser(t)
	su := randomSub(u)

	testCases := []struct {
		name          string
		buildStubs    func(s *mockdb.MockStore)
		checkResponse func(t *testing.T, r *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			buildStubs: func(s *mockdb.MockStore) {
				s.EXPECT().
					GetSub(gomock.Any(), gomock.Eq(su.ID)).
					Times(1).
					Return(su, nil)
			},
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, w.Code)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			c := gomock.NewController(t)
			defer c.Finish()

			ms := mockdb.NewMockStore(c)
			tc.buildStubs(ms)

			s := newTestServer(t, ms)

			url := fmt.Sprintf("/subs/%d", su.ID)
			r, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			w := httptest.NewRecorder()

			s.router.ServeHTTP(w, r)
			tc.checkResponse(t, w)

		})
	}
}
