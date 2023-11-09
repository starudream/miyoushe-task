package miyoushe

import (
	"testing"
	"time"

	"github.com/starudream/go-lib/core/v2/utils/testutil"

	"github.com/starudream/miyoushe-task/util"
)

func TestQRCodeLogin(t *testing.T) {
	account := GetAccount(t)

	data1, err := GenQRCode(account)
	testutil.LogNoErr(t, err, data1, data1.Url)

	time.Sleep(5 * time.Second)

loop:

	time.Sleep(5 * time.Second)

	data2, err := QueryQRCode(data1.Ticket, account)
	if IsRetCode(err, RetCodeQRCodeExpired) {
		t.Skip("qrcode expired")
	}
	testutil.LogNoErr(t, err, data2)

	if !data2.Stat.IsConfirmed() {
		goto loop
	}

	account.Uid = data2.Payload.Uid
	account.GToken = data2.Payload.Token

	data3, err := GetSTokenByGToken(GetAccount(t))
	testutil.LogNoErr(t, err, data3)

	account.Mid = data3.UserInfo.Mid
	account.SToken = data3.Token.Token

	SaveAccount(t, account)
}

func TestGenQRCode(t *testing.T) {
	data, err := GenQRCode(GetAccount(t))
	testutil.LogNoErr(t, err, data, data.Url, util.QRCode(data.Url))
}

func TestQueryQRCode(t *testing.T) {
	data, err := QueryQRCode("654b81644b7cf30567d6d20c", GetAccount(t))
	testutil.LogNoErr(t, err, data)
}

func TestGetSTokenByGToken(t *testing.T) {
	data, err := GetSTokenByGToken(GetAccount(t))
	testutil.LogNoErr(t, err, data)
}

func TestGetCTokenBySToken(t *testing.T) {
	data, err := GetCTokenBySToken(GetAccount(t))
	testutil.LogNoErr(t, err, data)
}

func TestGetLTokenBySToken(t *testing.T) {
	data, err := GetLTokenBySToken(GetAccount(t))
	testutil.LogNoErr(t, err, data)
}
