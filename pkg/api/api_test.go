package api_test

import (
	"testing"
	"github.com/Berlin-opendb-hack/opendb-mock/pkg/api"
	"net/url"
	"net/http/httptest"
	"github.com/Berlin-opendb-hack/opendb-mock/pkg/api/mock"
	"github.com/golang/mock/gomock"
)


func TestDeutscheBankClient_GetCashAccounts(t *testing.T) {

	mockCtrl := gomock.NewController(t)

	client := mock_api.NewMockRequestor(mockCtrl)
	response := httptest.NewRecorder()
	response.WriteString(`[{"iban":"DE10000000000000000306","balance":100.95,"productDescription":"persönliches Konto"}]`)
	response.Code = 200

	client.EXPECT().Do(gomock.Any()).Times(1).Return(response.Result(), nil)


	gateWay, err := api.NewDeutscheBankClient(
		url.URL{Scheme: "https", Host: "whatever-api.db.com", Path: "/gw/dbapi/v0"},
		client,
	)
	if (nil != err) {
		t.Error(err.Error())
	}
	token := "asdf"
	cashAccounts, err := gateWay.GetCashAccounts(token, "")
	if nil != err {
		t.Error(err.Error())
	}
	if 1 != len(cashAccounts) {
		t.Errorf("Lenght is not ok, got %d",  len(cashAccounts))
	}

}

func TestDeutscheBankClient_GetTransactions(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	client := mock_api.NewMockRequestor(mockCtrl)
	response := httptest.NewRecorder()
	response.WriteString(`[{"amount": -31, "counterPartyName": "BlaBlaCar", "usage": "Rech 8886590 Köln-Berlin", "date": "2016-10-28"},{"amount": -75.5, "counterPartyName": "Studentenwerk Berlin", "usage": "Semesterticket", "date": "2016-10-27"}]`)
	response.Code = 200

	client.EXPECT().Do(gomock.Any()).Times(1).Return(response.Result(), nil)

	gateWay, err := api.NewDeutscheBankClient(
		url.URL{Scheme: "https", Host: "whatever-api.db.com", Path: "/gw/dbapi/v0"},
		client,
	)
	if (nil != err) {
		t.Error(err.Error())
	}
	token := "asdf"
	cashTransactions, err := gateWay.GetTransactions(token, "")
	if nil != err {
		t.Error(err.Error())
	}
	if 2 != len(cashTransactions) {
		t.Errorf("Lenght is not ok, got %d",  len(cashTransactions))
	}
}

