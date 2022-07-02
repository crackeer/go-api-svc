package define

// AppConfig Config
type AppConfig struct {
	Port            int64  `yaml:"port"`
	PageSize        int64  `yaml:"page_size"`
	Env             string `yaml:"env"`
	DatabaseConfDir string `yaml:"database_conf_dir"`
}
