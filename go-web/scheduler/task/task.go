package main

import (
	"github.com/jiaofanting/go/go-web/api/dao"
	"log"
	"errors"
	"go/ast"
)

func VideoClearDispatcher(dc dataChan) error {
	res,err := dao.ReadVideoDeletionRecord(3)
	if err != nil {
		log.Printf("clear dispatcher error:%v",err)
		return err
	}
	if len(res) ==0 {
		return errors.New("all tasks finished")
	}
	for _,id:=range res{
		dc <- id
	}
	return nil
}

func VideoClearExecutor(dc dataChan) error {
	forloop:
		for  {
			select {
			case vid:= <-dc:
				go func(id interface{}) {
					err := dao.delVedio
					if
				}()

			}
		}
}