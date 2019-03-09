package config

const (
	// ConfigBasePath 基础配置文件config
	ConfigBasePath = "config"
	// MainConfigFile 主配置文件路径
	MainConfigFile = "app.conf"
	// DataBaseConfigFile 数据库配置文件路径
	DataBaseConfigFile = "database.json"
	// CacheConfigFile  缓存配置文件路径
	CacheConfigFile = "cache.json"
	// LoggerConfigFile 日志配置文件路径
	LoggerConfigFile = "logger.conf"
)

type MySQLConfig struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	IP       string `json:"ip"`
	DB       string `json:"db"`
}

type CacheConfig struct {
	IP string `json:"ip"`
}

type RPCClient struct {
	ServiceName         string   `json:"service_name"`
	IP                  []string `json:"ip"`
	URI                 string   `json:"url"`
	ConnTimeoutMs       int      `json:"conn_time_out_ms"`
	RetryCount          int      `json:"retry_count"`
	RetryTimeoutMs      int      `json:"retry_timeout_ms"`
	MaxIdleConnsPerHost int      `json:"max_idle_conns_per_host"`
}
