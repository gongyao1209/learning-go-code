package config

const env = "local"

// Get ...
func Get(item string) string {
	type confMap map[string]string
	conf := make(map[string]confMap)

	conf["rdtest"] = map[string]string{
		"Port":  "9001",
		"Mysql": "root:S12p_w99Q@tcp(127.0.0.1:3306)/phalcon_test?charset=utf8",
		"Redis": "127.0.0.1:6379",
	}

	conf["staging"] = map[string]string{
		"Port":  "9001",
		"Mysql": "root:S12p_w99Q@tcp(172.17.79.19:3306)/staging?charset=utf8",
		"Redis": "172.17.79.19:6379",
	}

	conf["pre"] = map[string]string{
		"Port":  "9002",
		"Mysql": "xddwrite:isd@kslI45ws@tcp(119.254.108.200:4041)/phalcon?charset=utf8",
		"Redis": "119.254.108.200:6379",
	}

	conf["local"] = map[string]string{
		"Port":  "9002",
		"Mysql": "root:S12p_w99Q@tcp(139.198.5.192:3306)/xiangxin0828?charset=utf8",
		"Redis": "119.254.108.200:6379",
	}

	return conf[env][item]
}
