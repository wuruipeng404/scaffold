/*
* @Author: Rumple
* @Email: ruipeng.wu@cyclone-robotics.com
* @DateTime: 2022/8/12 11:01
 */

package orm

import (
	"fmt"
)

// some default sort
const (
	createTime = "created_at"
	updateTime = "updated_at"
	deleteTime = "deleted_at"
	id         = "id"
)

var (
	CreateTimeAsc  = Asc(createTime)
	UpdateTimeAsc  = Asc(updateTime)
	CreateTimeDesc = Desc(createTime)
	UpdateTimeDesc = Desc(updateTime)
	IdDesc         = Desc(id)
)

func Asc(field string) string {
	return fmt.Sprintf("%s asc", field)
}

func Desc(field string) string {
	return fmt.Sprintf("%s desc", field)
}
