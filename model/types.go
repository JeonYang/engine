package model

type EngineInfo struct {
	Version []int `json:"version"`
	Id      []int `json:"id"`
}

type EngineConf struct {
	RefreshInterval   int                    `json:"refreshInterval"`
	Command           string                 `json:"command"` //"command":"stop/restart",
	Upgrade           *EngineUpgrade         `json:"upgrade"`
	PluginVersionList []*PluginVersion       `json:"pluginVersion"`
	PluginConf        map[string]*PluginConf `json:"pluginConf"`
}

type EngineUpgrade struct {
	Url     string `json:"url"`
	Md5     string `json:"md5"`
	Version []int  `json:"version"`
}

type Plugin struct {
	Name    string `json:"name"`
	Version []int  `json:"version"`
	Md5     string `json:"md5"`
}

type PluginFunc struct {
	PluginFunction PluginFunction
	CacheConf      string
	Running        bool
}

type PluginFunction interface {
	Start(conf string)
	Stop()
	Version() []int
	Name() string
}

type PluginConf struct {
	Name string
	Conf string
}
