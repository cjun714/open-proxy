package open

import (
	"errors"
	"io/ioutil"
	"net/http"
	"util/json"
	"util/log"
)

// {"bizDesc":"Repeatedly apply for a token in a short time","bizCode":"1","expire":25,"token":"KuQjRerX2EJ6WjABMou3p7I2TIeJgqqW"}

type LoginInfo struct {
	BizDesc string `json:"bizDesc"`
	BizCode string `json:"bizCode"`
	Expire  int    `json:"expire"`
	Token   string `json:"token"`
}

//{"respCode":"06","respDesc":"No message available"}

type LoginError struct {
	RespCode string `json:"respCode"`
	RespDesc string `json:"respDesc"`
}

//"http://20.26.33.122:32020/open/token?appKey=1000000007&secret=V4ML6U3GRZFZ6WPOEL3NGNDZRD1UD73FM0UZGXYK8FJIWMB8F5NS5DUZ8PKWW1N9&timestamp=''"
func GetToken(appKey, secret string) (string, error) {
	url := "http://20.26.33.122:32020/open/token?appKey=" + appKey + "&secret=" + secret + "&timestamp=''"
	log.I("get token:", url)

	client := &http.Client{}
	request, e := http.NewRequest("GET", url, nil)

	if e != nil {
		log.E(e)
		return "", e
	}

	resp, e := client.Do(request)
	if e != nil {
		return "", e
	}
	defer resp.Body.Close()

	log.I("resp: ", resp.Status)

	body, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return "", e
	}
	if len(body) != 0 {
		log.I(string(body))
	}

	if resp.StatusCode == 200 {
		result := LoginInfo{}
		json.ToObj(body, &result)
		log.H("token: ", result.Token)
		return result.Token, nil
	} else if resp.StatusCode == 400 {
		err := LoginError{}
		json.ToObj(body, &err)
		log.E("error: ", err.RespDesc)
		return "", errors.New(err.RespDesc)
	}

	return "", nil
}
