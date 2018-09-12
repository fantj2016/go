package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"os"
	"time"
)

func streamHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params)  {
	//获取vid
	vid := p.ByName("vid-id")
	vl := VIDEO_DIR+vid
	//打开
	video ,err := os.Open(vl)
	if err != nil {
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

}