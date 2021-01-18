package model

import (
	"database/sql"
	"fmt"
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
	"strings"
)

var (
	userFieldNames          = builderx.FieldNames(&User{})
	userRows                = strings.Join(userFieldNames, ",")
	userRowsExpectAutoSet   = strings.Join(stringx.Remove(userFieldNames, "`create_time`", "`update_time`","`birthday`"), ",")
	userRowsWithPlaceHolder = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheUserInviteCodePrefix = "cache#User#inviteCode#"
	cacheUserIdPrefix         = "cache#User#id#"
	cacheUserMobilePrefix     = "cache#User#mobile#"
)

type (
	UserModel interface {
		Insert(data User) (sql.Result, error)
		FindOne(id int64) (*User, error)
		FindOneByMobile(mobile string) (*User, error)
		FindOneByInviteCode(inviteCode string) (*User, error)
		Update(data User) error
		Delete(id int64) error
		Trans(fn func(session sqlx.Session)error)error
		TranInsert(session sqlx.Session,data User) (err error)
	}

	defaultUserModel struct {
		sqlc.CachedConn
		table string
	}

	User struct {
		Mobile     string `db:"mobile"`      // 手机号
		InviteCode string `db:"invite_code"` // 邀请码
		Id         int64  `db:"id"`
		CreateTime string `db:"create_time"` // 创建时间
		UpdateTime string `db:"update_time"` // 更新时间
		Version    int64  `db:"version"`     // 版本号
		Nickname   string `db:"nickname"`    // 昵称
		Avatar     string `db:"avatar"`      // 头像
		Info       string `db:"info"`        // 个人简介
		Sex        int64  `db:"sex"`         // 0:保密 1:男 2:女
		Birthday   string `db:"birthday"`    // 出生年月日
		AreaId     int64  `db:"area_id"`     // 地区
	}
)

func NewUserModel(conn sqlx.SqlConn, c cache.CacheConf) UserModel {
	return &defaultUserModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      UserTable(),
	}
}

func UserTable() string {
	return `user`
}

func (m *defaultUserModel) Insert(data User) (sql.Result, error) {
	userMobileKey := fmt.Sprintf("%s%v", cacheUserMobilePrefix, data.Mobile)
	userInviteCodeKey := fmt.Sprintf("%s%v", cacheUserInviteCodePrefix, data.InviteCode)
	ret, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, userRowsExpectAutoSet)
		return conn.Exec(query, data.Mobile, data.InviteCode, data.Id, data.Version, data.Nickname, data.Avatar, data.Info, data.Sex, data.AreaId)
	}, userMobileKey, userInviteCodeKey)
	return ret, err
}

func (m *defaultUserModel) FindOne(id int64) (*User, error) {
	userIdKey := fmt.Sprintf("%s%v", cacheUserIdPrefix, id)
	var resp User
	err := m.QueryRow(&resp, userIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where id = ? limit 1", userRows, m.table)
		return conn.QueryRow(v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) FindOneByMobile(mobile string) (*User, error) {
	userMobileKey := fmt.Sprintf("%s%v", cacheUserMobilePrefix, mobile)
	var resp User
	err := m.QueryRowIndex(&resp, userMobileKey, m.formatPrimary, func(conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where mobile = ? limit 1", userRows, m.table)
		if err := conn.QueryRow(&resp, query, mobile); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) FindOneByInviteCode(inviteCode string) (*User, error) {
	userInviteCodeKey := fmt.Sprintf("%s%v", cacheUserInviteCodePrefix, inviteCode)
	var resp User
	err := m.QueryRowIndex(&resp, userInviteCodeKey, m.formatPrimary, func(conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where invite_code = ? limit 1", userRows, m.table)
		if err := conn.QueryRow(&resp, query, inviteCode); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) Update(data User) error {
	userIdKey := fmt.Sprintf("%s%v", cacheUserIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where id = ?", m.table, userRowsWithPlaceHolder)
		return conn.Exec(query, data.Mobile, data.InviteCode, data.Version, data.Nickname, data.Avatar, data.Info, data.Sex, data.Birthday, data.AreaId, data.Id)
	}, userIdKey)
	return err
}

func (m *defaultUserModel) Delete(id int64) error {
	data, err := m.FindOne(id)
	if err != nil {
		return err
	}

	userIdKey := fmt.Sprintf("%s%v", cacheUserIdPrefix, id)
	userMobileKey := fmt.Sprintf("%s%v", cacheUserMobilePrefix, data.Mobile)
	userInviteCodeKey := fmt.Sprintf("%s%v", cacheUserInviteCodePrefix, data.InviteCode)
	_, err = m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where id = ?", m.table)
		return conn.Exec(query, id)
	}, userInviteCodeKey, userIdKey, userMobileKey)
	return err
}

func (m *defaultUserModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheUserIdPrefix, primary)
}

func (m *defaultUserModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", userRows, m.table)
	return conn.QueryRow(v, query, primary)
}

func (m *defaultUserModel) Trans(fn func(session sqlx.Session)error) error  {
	err := m.Transact(func(session sqlx.Session) error {
		err := fn(session)
		if err != nil{
			return err
		}
		return nil
	})
	return err
}

//事务插入
func (m *defaultUserModel) TranInsert(session sqlx.Session,data User) (err error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, userRowsExpectAutoSet)
	_,err = session.Exec(query, data.Mobile, data.InviteCode, data.Id, data.Version, data.Nickname, data.Avatar, data.Info, data.Sex, data.AreaId)
	return
}



//注册
func (m *defaultUserModel) Register(user User,userAuth UserAuth) error  {

	return  m.Transact(func(session sqlx.Session) error {

		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, userRowsExpectAutoSet)
		_,err := session.Exec(query, user.Mobile, user.InviteCode, user.Id, user.Version, user.Nickname, user.Avatar, user.Info, user.Sex, user.AreaId)
		if err != nil{
			return err
		}

		query = fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", UserAuthTable(), userAuthRowsExpectAutoSet)
		_,err = session.Exec(query,userAuth.UserId, userAuth.AuthType, userAuth.Id, userAuth.Version, userAuth.AuthKey, userAuth.AuthSecret)
		if err != nil{
			return err
		}

		return nil
	})

}

