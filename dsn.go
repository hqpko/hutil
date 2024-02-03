package hutils

import (
	"fmt"
	"time"
)

type MysqlDSNCreator struct {
	user         string
	password     string
	host         string
	port         string
	dbName       string
	parseTime    bool
	charset      string
	loc          string
	timeout      time.Duration
	readTimeout  time.Duration
	writeTimeout time.Duration
}

func NewMysqlDSNCreator(user, password, host, port, dbName string) *MysqlDSNCreator {
	return &MysqlDSNCreator{
		user:      user,
		password:  password,
		host:      host,
		port:      port,
		dbName:    dbName,
		parseTime: true,
		charset:   "utf8mb4",
		loc:       "Local",
	}
}

func (opts *MysqlDSNCreator) DataSourceName() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=%t&charset=%s&loc=%s", opts.user, opts.password, opts.host, opts.port, opts.dbName, opts.parseTime, opts.charset, opts.loc)
}

func (opts *MysqlDSNCreator) SetParseTime(parseTime bool) *MysqlDSNCreator {
	opts.parseTime = parseTime
	return opts
}

func (opts *MysqlDSNCreator) SetCharset(charset string) *MysqlDSNCreator {
	opts.charset = charset
	return opts
}

func (opts *MysqlDSNCreator) SetLoc(loc string) *MysqlDSNCreator {
	opts.loc = loc
	return opts
}

func (opts *MysqlDSNCreator) SetTimeout(timeout time.Duration) *MysqlDSNCreator {
	opts.timeout = timeout
	return opts
}

func (opts *MysqlDSNCreator) SetReadTimeout(readTimeout time.Duration) *MysqlDSNCreator {
	opts.readTimeout = readTimeout
	return opts
}

func (opts *MysqlDSNCreator) SetWriteTimeout(writeTimeout time.Duration) *MysqlDSNCreator {
	opts.writeTimeout = writeTimeout
	return opts
}
