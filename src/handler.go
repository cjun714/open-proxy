package main

import (
	"fmt"
	"net/http"
	"unsafe"
	"util/json"
	"util/log"

	"open"

	"github.com/julienschmidt/httprouter"
)

func token(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.H("/open/token -------------------------------------------------")
	r.ParseForm()

	appKey := r.Form.Get("appKey")
	log.I("appKey:", appKey)

	secret := r.Form.Get("secret")
	log.I("secret:", secret)

	// result := make(map[string]string)
	// if key != "1000000007" || secret != "V4ML6U3GRZFZ6WPOEL3NGNDZRD1UD73FM0UZGXYK8FJIWMB8F5NS5DUZ8PKWW1N9" {
	// 	result["token"] = "Access denied"
	// 	writeObj(w, result)
	// 	return
	// }

	// result["bizDesc"] = "success"
	// result["bizCode"] = "1"
	// result["expire"] = "25"
	// result["token"] = "uOGlzaiq6C2Y0YJjJAk3jonjNEr7OA6x"

	token, e := open.GetToken(appKey, secret)
	if e != nil {
		log.E(e)
	}
	writeObj(w, token)
}

func hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.H("/poc/hello  -------------------------------------------------")
	token := r.Header.Get("token")
	log.I("token:", token)
	writeObj(w, "Hello world!")
}

func writeObj(w http.ResponseWriter, obj interface{}) {
	bytes := json.ToJSON(obj)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, token")
	w.Header().Set("content-type", "application/json")

	fmt.Fprintf(w, *(*string)(unsafe.Pointer(&bytes)))
}
