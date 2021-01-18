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
	userAuthFieldNames          = builderx.FieldNames(&UserAuth{})
	userAuthRows                = strings.Join(userAuthFieldNames, ",")
	userAuthRowsExpectAutoSet   = strings.Join(stringx.Remove(userAuthFieldNames, "`create_time`", "`update_time`"), ",")
	userAuthRowsWithPlaceHolder = strings.Join(stringx.Remove(userAuthFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheUserAuthIdPrefix = "cache#UserAuth#id#"
)

type (
	UserAuthModel interface {
		Insert(data UserAuth) (sql.Result, error)
		FindOne(id int64) (*UserAuth, error)
		Update(data UserAuth) error
		Delete(id int64) error
		Trans(fn func(session sqlx.Session)error)error
		TranInsert(session sqlx.Session,data UserAuth)error
	}

	defaultUserAuthModel struct {
		sqlc.CachedConn
		table string
	}

	UserAuth struct {
		UserId     int64  `db:"user_id"`
		AuthType   int64  `db:"auth_type"` // 授权类型
		Id         int64  `db:"id"`
		CreateTime string `db:"create_time"`
		UpdateTime string `db:"update_time"`
		Version    int64  `db:"version"`
		AuthKey    string `db:"auth_key"`    // 唯一值（手机、openid）
		AuthSecret string `db:"auth_secret"` // 站内密码，站外token
	}
)


func NewUserAuthModel(conn sqlx.SqlConn, c cache.CacheConf) UserAuthModel {
	return &defaultUserAuthModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      UserAuthTable(),
	}
}

func UserAuthTable() string {
	return "user_auth"
}

func (m *defaultUserAuthModel) Insert(data UserAuth) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, userAuthRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.UserId, data.AuthType, data.Id, data.Version, data.AuthKey, data.AuthSecret)

	return ret, err
}

func (m *defaultUserAuthModel) FindOne(id int64) (*UserAuth, error) {
	userAuthIdKey := fmt.Sprintf("%s%v", cacheUserAuthIdPrefix, id)
	var resp UserAuth
	err := m.QueryRow(&resp, userAuthIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where id = ? limit 1", userAuthRows, m.table)
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

func (m *defaultUserAuthModel) Update(data UserAuth) error {
	userAuthIdKey := fmt.Sprintf("%s%v", cacheUserAuthIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where id = ?", m.table, userAuthRowsWithPlaceHolder)
		return conn.Exec(query, data.UserId, data.AuthType, data.Version, data.AuthKey, data.AuthSecret, data.Id)
	}, userAuthIdKey)
	return err
}

func (m *defaultUserAuthModel) Delete(id int64) error {

	userAuthIdKey := fmt.Sprintf("%s%v", cacheUserAuthIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where id = ?", m.table)
		return conn.Exec(query, id)
	}, userAuthIdKey)
	return err
}

func (m *defaultUserAuthModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheUserAuthIdPrefix, primary)
}

func (m *defaultUserAuthModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", userAuthRows, m.table)
	return conn.QueryRow(v, query, primary)
}

func (m *defaultUserAuthModel) Trans(fn func(session sqlx.Session)error) error  {
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
func (m *defaultUserAuthModel) TranInsert(session sqlx.Session,data UserAuth) (err error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", UserAuthTable(), userAuthRowsExpectAutoSet)
	_,err = session.Exec(query,data.UserId, data.AuthType, data.Id, data.Version, data.AuthKey, data.AuthSecret)
	return
}
