package flagvar

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

type DsnVar struct {
	dsn string
}

type dbinfo struct {
	Database string `json:"database"`
	User     string `json:"user"`
	Passwd   string `json:"passwd"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
}

func (dv *DsnVar) Help() string {
	return "user:pass@protocol(ip:port)/db or dbinfo.json. see https://pkg.go.dev/database/sql#Open"
}

func (dv *DsnVar) json2dsn(fname string) (string, error) {
	bytes, err := ioutil.ReadFile(fname)
	if err != nil {
		return "", err
	}

	var db dbinfo
	err = json.Unmarshal(bytes, &db)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", db.User, db.Passwd, db.Host, db.Port, db.Database), nil
}

/**
 * data source として適切か?
 */
func (dv *DsnVar) isDsnFormat(dsn string) bool {
	pattern := `^[a-z0-9]+:.*@tcp\([a-z0-9_.-]+:[0-9]+\)/[a-zA-Z0-9_.-]+$`
	matched, _ := regexp.MatchString(pattern, dsn)
	return matched
}

func (dv *DsnVar) Set(opt string) error {
	dsn := opt
	if dsn == "" {
		dv.dsn = dsn
		return nil
	}

	if strings.HasSuffix(opt, ".json") {
		d, err := dv.json2dsn(opt)
		if err != nil {
			return err
		}
		dsn = d
	}

	if !dv.isDsnFormat(dsn) {
		return fmt.Errorf("invalid format")
	}
	dv.dsn = dsn
	return nil
}

func (dv *DsnVar) String() string {
	return dv.dsn
}

func (dv *DsnVar) Value() string {
	return dv.dsn
}
