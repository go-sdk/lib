package token

import (
	"testing"
	"time"

	"github.com/go-sdk/lib/testx"
)

func TestNew(t *testing.T) {
	t.Log(GetExpire())

	SetExpire(time.Hour)

	tk1 := New("*", "1", -1, WithID("abc"), WithExtra(Extra{"uid": "aaa", "foo": "bar"}))
	str, err := tk1.Sign()
	testx.AssertNoError(t, err)
	t.Log(str)
	t.Log(tk1.GetID())
	t.Log(tk1.GetIssuer())
	t.Log(tk1.GetUserId())
	t.Log(tk1.GetIssuedAt())
	t.Log(tk1.GetExtra())
	t.Log(tk1.IsExpired())

	tk2, err := Parse(str)
	testx.AssertNoError(t, err)
	t.Log(tk2.GetExtra())

	tk3 := New("*", "2", time.Now().Unix())
	ctx1 := tk3.WithContext()

	testx.AssertEqual(t, tk3.GetID(), GetID(ctx1))

	time.Sleep(time.Second)

	tk4 := tk3.Refresh()

	t.Log(tk3.GetIssuedAt())
	t.Log(tk4.GetIssuedAt())
	t.Log(tk4.IsExpired())

	t.Log(GetID(ctx1))
	t.Log(GetIssuer(ctx1))
	t.Log(GetUserId(ctx1))
	t.Log(GetIssuedAt(ctx1))
	t.Log(GetExtra(ctx1))
}
