package weibogo

import (
	"testing"
)

func TestWeibo_Call(t *testing.T) {
	accessToken := ""
	var result interface{}
	err := new(Weibo).Call("statuses/share", "post", accessToken, Params{"status": "this is a new weibo2. http://***"}, &result)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v\n", result)
}
