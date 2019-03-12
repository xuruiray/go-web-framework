package config

const (
	// ConfigBasePath 基础配置文件config
	ConfigBasePath = "config"
	// MainConfigFile 主配置文件路径
	AppConfigFile = "app.json"
	// DataBaseConfigFile 数据库配置文件路径
	DataBaseConfigFile = "database.json"
	// CacheConfigFile  缓存配置文件路径
	CacheConfigFile = "cache.json"
	// RPCConfigFile RPC配置文件路径
	RPCConfigFile = "rpc.json"
	// LoggerConfigFile 日志配置文件路径
	LoggerConfigFile = "logger.conf"
)

var MainConfig AppConfig

type AppConfig struct {
	Name string `json:"name"`
	Port string `json:"port"`
}

type MySQLConfig struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	IP       string `json:"ip"`
	DB       string `json:"db"`
}

type CacheConfig struct {
	IP                string `json:"ip"`
	MaxIdle           int    `json:"max_idle"`
	MaxActive         int    `json:"max_active"`
	IdleTimeoutMs     int    `json:"idle_timeout_ms"`
	MaxConnLifeTimeMs int    `json:"max_conn_life_time_ms"`
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
