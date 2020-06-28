package db

import (
	"github.com/treeforest/renju-server/global"
	"gopkg.in/mgo.v2/bson"
)

var (
	// 用户服
	mgUserDatabase 			= "user" 		// 数据库名
	mgIndexTableName 		= "index_id" 	// 索引表
	mgUserInfoTableName 	= "user_info" 	// 用户信息表
	mgBlackList 			= "black_list" 	// 黑名单
)

// 数据定义
type UserInfo struct {
	UserID      uint64 `bson:"_id"`
	UserName    string `bson:"name"`
	Password    string `bson:"pwd"`
	PhoneNumber string `bson:"pn"`
	Email 		string `bson:"email"`
}

// 自增序号（id）接口
type indexTable struct {
	ID   uint64 `bson:"id"`
	Name string `bson:"Name"`
}

// -----------------
// 用户数据库一些列操作工作(Mongodb)
type user struct {}

func (u *user) GetDatabase() string {
	return mgUserDatabase
}

// 判断电话号是否在UserInfo表中
func (u *user) ExistPhoneNumber(phoneNumber string) bool {
	ok, _ := global.Mongodb().Exists(u.GetDatabase(), u.getUserInfoTableName(), bson.M{"pn":phoneNumber})
	return ok
}

// 判断邮箱号是否存在UserInfo表中
func (u *user) ExistEmail(email string) bool {
	ok, _ := global.Mongodb().Exists(u.GetDatabase(), u.getUserInfoTableName(), bson.M{"email":email})
	return ok
}

// 判断电话号是否存在于黑名单中
func (u *user) ExistBlackListPhoneNumber(account string) bool {
	ok, _ := global.Mongodb().Exists(u.GetDatabase(), u.getBlackListTableName(), bson.M{"pn":account})
	return ok
}

// 判断邮箱号是否存在于黑名单中
func (u *user) ExistBlackListEmail(email string) bool {
	ok, _ := global.Mongodb().Exists(u.GetDatabase(), u.getBlackListTableName(), bson.M{"email":email})
	return ok
}

/*
 * 插入一条用户的注册信息
 */
func (u *user) InsertUserInfo(info *UserInfo) error {
	return global.Mongodb().Insert(u.GetDatabase(), u.getUserInfoTableName(), info)
}

/*
 * 获取用户表中的自增ID
 */
func (u *user) GetUserInfoIncID() uint64 {
	return u.GetIncIDBase("UserInfo", 100000)
}

// 获取自增后的ID
func (u *user) GetIncID(Name string) uint64 {
	return u.GetIncIDBase(Name, 0)
}

/*
 * 获取自增后的ID
 *
 * 参数：
 *  Name：关键key值
 *  baseId：基础ID，若获取的ID小于baseId，则重置ID
 */
func (u *user) GetIncIDBase(Name string, baseId uint64) uint64 {
	index := indexTable{}
	err := global.Mongodb().FindAndInc(u.GetDatabase(), u.getIndexTableName(), bson.M{"Name": Name}, bson.M{"id": 1}, &index)
	if err != nil {
		return 0
	}

	if baseId > index.ID {
		global.Mongodb().FindAndInc(u.GetDatabase(), u.getIndexTableName(), bson.M{"Name": Name}, bson.M{"id": baseId}, &index)
	}

	return index.ID
}

func (u user) getIndexTableName() string {
	return mgIndexTableName
}

func (u user) getBlackListTableName() string {
	return mgBlackList
}

func (u *user) getUserInfoTableName() string {
	return mgUserInfoTableName
}