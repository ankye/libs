package db

//DBConfig db config
type DBConfig struct {
	Driver   string `yaml:"driver"`
	DSM      string `yaml:"dsm"`
	Timezone string `yaml:"timezone"`
}

//DBGroup db group
type DBGroup struct {
	MaxConnection int      `yaml:"maxConnection"`
	Master        DBConfig `yaml:"master"`
	Slaver        DBConfig `yaml:"slaver"`
}
