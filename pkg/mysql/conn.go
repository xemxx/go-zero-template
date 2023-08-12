package mysql

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/xemxx/go-zero-template/internal/config"
)

var connIns sqlx.SqlConn

func Init(c config.Config) {
	conn := sqlx.NewMysql(c.Mysql.Dsn)
	connIns = conn
	// todo 检测db是否可用然后决定能否启动
}

func GetConn() sqlx.SqlConn {
	return connIns
}
