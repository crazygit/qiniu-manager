//
// 生成上传应用到七牛的TOKEN
//
package main

import (
	"encoding/json"
	"fmt"
	"github.com/crazygit/qiniu-manager/Godeps/_workspace/src/qiniupkg.com/api.v7/kodo"
	"net/http"
	"os"
)

var QiniuClient *kodo.Client

type Uptoken struct {
	Value string `json:"uptoken"`
}

func init() {
	qiniuAccessKey := os.Getenv("QINIU_ACCESS_KEY")
	qiniuSecretKey := os.Getenv("QINIU_SECRET_KEY")
	kodo.SetMac(qiniuAccessKey, qiniuSecretKey)
	zone := 0
	QiniuClient = kodo.New(zone, nil)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/uptoken", makeUptoken)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	fmt.Printf("Start server and listen on *:%s...", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println(err)
	}
}
func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("building..."))
}

func makeUptoken(w http.ResponseWriter, r *http.Request) {
	upload_key := r.FormValue("key")
	if upload_key == "" {
		http.Error(w, "missing required parameter: key", http.StatusBadRequest)
		return
	}
	qiniuBucketName := os.Getenv("QINIU_BUCKET_NAME")
	putpolicy := kodo.PutPolicy{
		Scope:      qiniuBucketName + ":" + upload_key,
		Expires:    3600,
		InsertOnly: 1, // 不允许覆盖或修改已经存在的同名文件
	}
	token, err := json.Marshal(Uptoken{Value: QiniuClient.MakeUptoken(&putpolicy)})
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(token)
}
