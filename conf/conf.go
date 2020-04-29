package conf

type appConf struct {
	ManagerList []string
	Version     []int
	Id          []int
}

var AppConf = &appConf{}

func InitConf() error {
	return nil
}
