package config_test

import (
	"testing"

	"github.com/gonethopper/libs/config"
)

type TConf struct {
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	EnableUDP bool   `yaml:"enableudp"`
}

type AKConf struct {
	Test *TConf `yaml:"test"`
}

func init() {

}
func TestParsePath(t *testing.T) {
	path := "./conf/test.yml"
	pc := config.ParsePath(path)
	if pc.Dir != "./conf/" {
		t.Errorf("not Expect dir result %s", pc.Dir)
	}
	if pc.Ext != ".yml" {
		t.Errorf("not Expect ext result %s", pc.Ext)
	}
	if pc.Name != "test" {
		t.Errorf("not Expect name result %s", pc.Name)
	}
	if pc.FileName != "test.yml" {
		t.Errorf("not Expect filename result %s", pc.FileName)
	}
}
func TestParseConfig(t *testing.T) {
	c := new(AKConf)
	c.Test = new(TConf)
	err := config.ParseConfig(c, "./conf/test.yml")
	if err != nil {
		t.Error(err.Error())
	} else {
		if c.Test.Host != "127.0.0.1" {
			t.Errorf("Expect 127.0.0.1 get %s", c.Test.Host)
		}
		if c.Test.Port != 443 {
			t.Errorf("Expect 443 get %d", c.Test.Port)
		}
		if c.Test.EnableUDP {
			t.Errorf("Expect false get true")
		}
	}

}

func TestSaveConfig(t *testing.T) {
	c := new(AKConf)
	c.Test = new(TConf)

	err := config.ParseConfig(c, "./conf/test.yml")
	if err != nil {
		t.Error("read config error")
	}
	config.SaveConfig(c, "./conf/test2.yml")

	except := new(AKConf)
	except.Test = new(TConf)
	err = config.ParseConfig(except, "./conf/test2.yml")
	if err != nil {
		t.Error("read expect config error")
	}

	if c.Test.Host != except.Test.Host {
		t.Errorf("Expect %s get %s", c.Test.Host, except.Test.Host)
	}
	if c.Test.Port != except.Test.Port {
		t.Errorf("Expect %d get %d", c.Test.Port, except.Test.Port)
	}
	if c.Test.EnableUDP != except.Test.EnableUDP {
		t.Errorf("Expect %t get %t", c.Test.EnableUDP, except.Test.EnableUDP)
	}

}
