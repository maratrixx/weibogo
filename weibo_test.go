package weibogo

import (
	"net/http"
	"path/filepath"
	"testing"
	"time"
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

func TestWeibo_Upload(t *testing.T) {
	img := "https://qny.smzdm.com/201908/20/5cc4d1658d5489495.jpg_d250.jpg"
	accessToken := ""
	var result interface{}

	cli := http.Client{Timeout: 3 * time.Second}
	r, e := cli.Get(img)
	if e != nil {
		t.Fatal(r, e)
	}

	err := new(Weibo).Upload(accessToken, Params{"status": "this is a new weibo. http://***"}, r.Body, filepath.Ext(img), &result)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v\n", result)
}
