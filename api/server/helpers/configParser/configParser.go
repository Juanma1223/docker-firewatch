package configParser

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// This function receives a fileDir and the corresponding golang struct (pointer) to map those configurations
func ParseConfigFile(fileDir string, fileStruct interface{}) {
	configFile, err := os.Open(fileDir)
	if err != nil {
		fmt.Println(err)
	}

	configFileBytes, err := io.ReadAll(configFile)
	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(configFileBytes, &fileStruct)
}
