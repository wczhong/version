package main

import (
	"errors"
	"fmt"
	"time"

	"example/mysql/version"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	null "gopkg.in/guregu/null.v3"
)

type WxUser struct {
	I_ID       int64     `db:"I_ID" json:"I_ID"`
	D_BIRTHDAY null.Time `db:"D_BIRTHDAY" json:"D_BIRTHDAY"`
}

// GetWxUserByWxUserID 根据 wx_user_id 查找用户
func GetWxUserByWxUserID(db *sqlx.DB) (result WxUser, err error) {
	query := fmt.Sprintf("SELECT * FROM wx_user_1 WHERE I_ID=290292725876527119")
	err = db.Get(&result, query)
	return
}

// TimeStrParseLocalWithFormat 指定格式的时间转换
func TimeStrParseLocalWithFormat(tStr string, format string) (t time.Time, err error) {
	if tStr == "" {
		err = errors.New("时间不可为空")
		return
	}
	loc, _ := time.LoadLocation("Local")
	t, err = time.ParseInLocation(format, tStr, loc)
	return
}

func main() {
	fmt.Println(version.String())
	_db, err := sqlx.Connect("mysql", "root:@tcp(127.0.0.1:3306)/broker?timeout=200ms&parseTime=true&loc=Local")
	if err != nil {
		fmt.Println("Initialization mysql driver failed: ", _db, err)
	}
	// r, e := GetWxUserByWxUserID(_db)
	// fmt.Println(r.I_ID)
	// fmt.Println(r.D_BIRTHDAY.Time.Unix())
	// fmt.Println(r, e)

	t, e := TimeStrParseLocalWithFormat("1986-09-11 00:00:00", "2006-01-02")
	fmt.Println(t.Unix())
	fmt.Println(e)
}
