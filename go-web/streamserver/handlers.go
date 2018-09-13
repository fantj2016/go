package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"os"
	"time"
	"fmt"
	"io/ioutil"
	"log"
	"io"
	"html/template"
)

func streamHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params)  {
	//获取vid
	vid := p.ByName("vid-id")
	vl := VIDEO_DIR+vid
	fmt.Print("拿到的路径："+vl)
	//打开
	video ,err := os.Open(vl)
	if err != nil {
		fmt.Printf("错误是%v",err)
		sendErrorResponse(w,http.StatusInternalServerError,"internal error:500")
	}
	//如果文件没有扩展名，加上header，浏览器会将其用mp4来解析
	w.Header().Set("Content-Type","video/mp4")
	//time.now 保证自由跳跃时间
	http.ServeContent(w,r,"",time.Now(),video)
	//关闭文件流
	defer video.Close()
}
func uploadHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params)  {
	//从 request中read 信息
	r.Body = http.MaxBytesReader(w,r.Body,MAX_UPLOAD_SIZE)
	//ParseMultipartForm 解析reader
	err := r.ParseMultipartForm(MAX_UPLOAD_SIZE);
	if  err != nil {
		sendErrorResponse(w,http.StatusBadRequest,"");
	}
	//通过FormFile来获取input name属性的值
	file ,_,err := r.FormFile("file")
	if err != nil {
		sendErrorResponse(w,http.StatusInternalServerError,"通过FormFile来获取input name属性的值错误")
		return
	}
	data,err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Read file error:%v",err)
		sendErrorResponse(w,http.StatusInternalServerError,"Read file error")
	}
	fileName := p.ByName("vid-id")
	err = ioutil.WriteFile(VIDEO_DIR+fileName,data,0666)
	if err != nil {
		sendErrorResponse(w,http.StatusInternalServerError,"WriteFile error")
	}
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w,"upload successfully")
}

func testPageHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	t, _ := template.ParseFiles(VIDEO_DIR+"upload.html")
	t.Execute(w, nil)
}