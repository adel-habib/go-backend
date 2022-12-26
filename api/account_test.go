package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	mockdb "github.com/adel-habib/go-backend/db/mock"
	db "github.com/adel-habib/go-backend/db/sqlc"
	"github.com/adel-habib/go-backend/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetAccountApi(t *testing.T) {

	account := randomAccount()

	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)

	store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account.ID)).Times(1).Return(account, nil)

	server := NewServer(store)

	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/accounts/%d", account.ID)

	req, err := http.NewRequest(http.MethodGet, url, nil)

	require.NoError(t, err)

	server.router.ServeHTTP(recorder, req)

	require.Equal(t, http.StatusOK, recorder.Code)

}

func randomAccount() db.Account {
	return db.Account{
		ID:       util.RandomInt(1, 1000),
		Owner:    util.RandomOwner(),
		Currency: util.RandomCurrency(),
		Balance:  util.RandomMoney(),
	}
}
