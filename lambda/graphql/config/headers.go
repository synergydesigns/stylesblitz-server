package config

func Getheaders() map[string]string {
	return map[string]string{
		"Content-Type":                 "application/json",
		"Control-Allow-Credentials":    "true",
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Methods": "GET,POST,OPTIONS",
		"Access-Control-Allow-Headers": "Connection, Host, Origin, Referer, Access-Control-Request-Method, Access-Control-Request-Headers, User-Agent, Accept, Content-Type, Authorization, Content-Length, X-Requested-With, Accept-Encoding, Accept-Language",
	}
}
