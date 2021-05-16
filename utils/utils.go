package utils

import "os"

func GenAbsoluteHomeDirPathWithConfig(configName string) string {
	homedir, _ := os.UserHomeDir()
	config := homedir + string(os.PathSeparator) + configName
	return config
}
