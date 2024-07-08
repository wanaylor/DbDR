package config

type ServiceConfig struct {
    Host string `json:"host"`
    Port int `json:"port"`
    Db string `json:db`
    Num_rows int `json:num_rows`
    Ms_delay int `json:ms_dalay`
}
