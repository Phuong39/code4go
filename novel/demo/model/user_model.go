/**
* 作者：刘时明
* 时间：2019/10/21-22:21
* 作用：
 */
package model

type User struct {
	Id         int64  `json:"id"`
	Uid        int64  `json:"uid"`
	UserName   string `json:"user_name"`
	Password   string `json:"password"`
	NickName   string `json:"nick_name"`
	HeadImg    string `json:"head_img"`
	Salt       string `json:"salt"`
	UpdateTime string `json:"update_time"`
	CreateTime string `json:"create_time"`
}

func InsertUser(user *User) (int64, error) {
	return mysqlEngine.Insert(user)
}

func FindUserByWhere() {

}
