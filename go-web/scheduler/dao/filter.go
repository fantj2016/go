package dao

import (
	"log"
)

func ReadVideoDeletionRecord(count int) ([]string, error) {
	stmtOut ,err := dbConn.Prepare("select video_id from video_info")

	var ids []string
	if err != nil {
		return ids,err
	}
	rows,err := stmtOut.Query(count)
	if err != nil {
		log.Printf("query video error")
		return ids,err
	}
	for rows.Next() {
		var id string
		err := rows.Scan(&id)
		if err != nil {
			return ids,err
		}
		ids = append(ids,id)
	}
	defer stmtOut.Close()
	return ids,err
}

func DelvideoDeletionRecord(vid string) error {
	stmtDel,err := dbConn.Prepare("delete from video_del_rec where id=?")
	if err != nil {
		return err
	}
	_,err = stmtDel.Exec(vid)
	if err != nil {
		log.Printf("Delete error:%v",err)
		return err
	}
	defer stmtDel.Close()
	return nil
}

