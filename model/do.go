package model

//配置文件中的变量
type Conf struct {
	DbHost     string `yaml:"dbHost"`
	DbPort     string `yaml:"dbPort"`
	DbDataBase string `yaml:"dbDatabase"`
	DbUser     string `yaml:"dbUser"`
	DbPwd      string `yaml:"dbPwd"`
}

type StremData struct {
	Id     int
	Output string
	Input  string
}



