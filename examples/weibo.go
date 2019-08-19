// 例子程序：调用微博API
package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/imttx/weibogo"
)

var (
	access_token = flag.String("access_token", "", "用户的访问令牌")
	image        = flag.String("image", "", "上传图片的位置")
	random       = rand.New(rand.NewSource(time.Now().UnixNano()))
	weibo        = weibogo.Weibo{}
)

func showUser() {
	fmt.Println("==== 测试 users/show ====")
	var user weibogo.User
	params := weibogo.Params{"screen_name": "人民日报"}
	err := weibo.Call("users/show", "get", *access_token, params, &user)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%#v\n", user)
	}
}

func getFriendsStatuses() {
	fmt.Println("==== 测试 statuses/friends_timeline ====")
	var statuses weibogo.Statuses
	params := weibogo.Params{"count": 10}
	err := weibo.Call("statuses/friends_timeline", "get", *access_token, params, &statuses)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, status := range statuses.Statuses {
			fmt.Println(status.Text)
		}
	}
}

func getUserStatus() {
	fmt.Println("==== 测试 statuses/user_timeline ====")
	var statuses weibogo.Statuses
	params := weibogo.Params{"screen_name": "人民日报", "count": 1}
	err := weibo.Call("statuses/user_timeline", "get", *access_token, params, &statuses)
	if err != nil {
		fmt.Println(err)
	} else if len(statuses.Statuses) > 0 {
		fmt.Printf("%#v\n", statuses.Statuses[0])
	}
}

func updateStatus() {
	fmt.Println("==== 测试 statuses/update ====")
	var status weibogo.Status
	params := weibogo.Params{"status": "测试" + strconv.Itoa(rand.Int())}
	err := weibo.Call("statuses/update", "status", *access_token, params, &status)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%#v\n", status)
	}
}

func uploadStatus() {
	fmt.Println("==== 测试 statuses/upload ====")
	var status weibogo.Status
	params := weibogo.Params{"status": "测试" + strconv.Itoa(rand.Int())}
	img, err := os.Open(*image)
	if err != nil {
		fmt.Println(err)
	}
	err = weibo.Upload(*access_token, params, img, filepath.Ext(*image), &status)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%#v\n", status)
	}
}

func main() {
	flag.Parse()
	showUser()
	getFriendsStatuses()
	getUserStatus()
	//updateStatus()
	//uploadStatus()
}
