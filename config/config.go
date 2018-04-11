package config

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	//log "github.com/Sirupsen/logrus"

	"github.com/spf13/viper"
	yaml "gopkg.in/yaml.v2"
)

//PathConfig 路径解析结构
type PathConfig struct {
	FileName string
	Name     string
	Ext      string
	Dir      string
}

//ParsePath 解析路径结构，获取文件名，文件路径，文件扩展名
func ParsePath(path string) *PathConfig {
	pc := &PathConfig{}

	base := filepath.Base(path)
	ext := filepath.Ext(path)
	name := strings.TrimSuffix(base, ext)
	dir, file := filepath.Split(path)

	pc.FileName = file
	pc.Name = name
	pc.Ext = ext
	pc.Dir = dir

	return pc
}

//ParseConfig returns the config struct
//反射字段定义需要大写
//search config file from global path:/etc/$name  user path:$HOME/.conf/$name
//on server current directory ./ and param path from where you want
func ParseConfig(c interface{}, path string) error {

	pc := ParsePath(path)

	if pc.Ext != ".yml" && pc.Ext != ".yaml" {
		return fmt.Errorf("file ext need set to .yml or .yaml")
	}
	// file
	viper.SetConfigName(pc.Name)
	viper.SetConfigType("yaml")

	// config dir
	viper.AddConfigPath("/etc/" + pc.Name)        // global
	viper.AddConfigPath("$HOME/.conf/" + pc.Name) // user
	viper.AddConfigPath(".")                      // on directory
	if len(path) > 0 {
		viper.AddConfigPath(pc.Dir)
	}
	// read config
	if err := viper.ReadInConfig(); err != nil {

		return err
	}

	err := viper.Unmarshal(&c)
	if err != nil {
		return err
	}

	return nil
}

//SaveConfig save config from struct
func SaveConfig(c interface{}, path string) error {
	pc := ParsePath(path)

	if pc.Ext != ".yml" && pc.Ext != ".yaml" {
		return fmt.Errorf("file ext need set to .yml or .yaml")
	}

	dirPath := pc.Dir
	if len(dirPath) <= 0 {
		usr, err := user.Current()
		if err != nil {
			return err
		}
		dirPath = filepath.Join(usr.HomeDir, ".conf/"+pc.Name)
	}

	// create config file if file not exist.
	if _, err := os.Stat(dirPath); err != nil {
		e := os.MkdirAll(dirPath, 0755)
		if e != nil {
			fmt.Println(e.Error())
		}
	}

	filePath := filepath.Join(dirPath, pc.FileName)
	if _, fileFindErr := os.Stat(filePath); fileFindErr != nil {
		// create config file if file not exist.
		file, err := os.Create(filePath)
		defer file.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
		if d, err := yaml.Marshal(c); err == nil {
			file.Write(d)
		}
	}
	return nil
}
