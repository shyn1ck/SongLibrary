package models

type AppConfig struct {
	AuthParams     AuthParams     `json:"auth"`
	LogParams      LogParams      `json:"log_params"`
	AppParams      AppParams      `json:"app_params"`
	PostgresParams PostgresParams `json:"postgres_params"`
}

type AuthParams struct {
	JwtTtlMinutes int `json:"jwt_ttl_minutes"`
}

type LogParams struct {
	LogDirectory     string `json:"log_directory"`      // Директория для логов
	LogInfo          string `json:"log_info"`           // Лог для информации
	LogError         string `json:"log_error"`          // Лог для ошибок
	LogWarn          string `json:"log_warn"`           // Лог для предупреждений
	LogDebug         string `json:"log_debug"`          // Лог для отладочной информации
	MaxSizeMegabytes int    `json:"max_size_megabytes"` // Максимальный размер лога в мегабайтах
	MaxBackups       int    `json:"max_backups"`        // Максимальное количество резервных копий логов
	MaxAge           int    `json:"max_age"`            // Максимальный возраст логов в днях
	Compress         bool   `json:"compress"`           // Сжимать ли логи
	LocalTime        bool   `json:"local_time"`         // Использовать ли локальное время
}

type AppParams struct {
	GinMode    string `json:"gin_mode"`    // Режим Gin (например, debug или release)
	PortRun    string `json:"port_run"`    // Порт, на котором будет запущен сервер
	ServerURL  string `json:"server_url"`  // URL сервера
	ServerName string `json:"server_name"` // Имя сервера
}

type PostgresParams struct {
	Host     string `json:"host"`     // Хост базы данных
	Port     string `json:"port"`     // Порт базы данных
	User     string `json:"user"`     // Имя пользователя базы данных
	Database string `json:"database"` // Имя базы данных
}
