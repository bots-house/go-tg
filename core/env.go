package core

import "strings"

const (
	EnvLocal      = "local"
	EnvStaging    = "staging"
	EnvProduction = "production"
)

func GetEnv(env string) string {
	for _, v := range []string{EnvLocal, EnvProduction, EnvStaging} {
		if v == strings.ToLower(env) {
			return v
		}
	}
	return EnvLocal
}
