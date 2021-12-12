package main

import (
	"github.com/pkg/errors"
)

//官方处理ErrNoRows部分代码
//type Row struct {
//	err  error
//	rows *Rows
//}
//
//func (r *Row) Scan(dest ...interface{}) error {
//	if r.err != nil {
//		return r.err
//	}
//	//.............
//	//.............其他逻辑
//	//这里
//	if !r.rows.Next() {
//		if err := r.rows.Err(); err != nil {
//			return err
//		}
//		return ErrNoRows
//	}
//	//.............
//	//.............其他逻辑
//	return r.rows.Close()
//}
//下面处理Scan时候遇到报错处理
func Demo()  error {
	//.....
	err := mysqlDB.QueryRow("").Scan(dest)
	if err != nil {
		return errors.Wrap(err, "Demo failed")
	}
	//.......
	return  nil
}