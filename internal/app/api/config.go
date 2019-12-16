package api

import "github.com/panda-next-team/poolrank-agent/internal/pkg"

type Config struct {
	Port  int
	Mysql pkg.MysqlConfig
}
