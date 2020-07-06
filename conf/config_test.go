package conf

import (
	"github.com/ahviplc/GoJustToolc/UConsole"
	"testing"
)

func TestConfig(t *testing.T) {
	UConsole.Log(Mysql_Link_Info) // root:root@(192.168.0.10:3306)/bubble?charset=utf8&parseTime=True&loc=Local
}
