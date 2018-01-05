package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/dotaspirit/dotaspirit/types"
)

func LoadAppConfig() types.AppConfig {
	buf := bytes.NewBuffer(nil)
	f, _ := os.Open("config/app.json")
	io.Copy(buf, f)
	f.Close()

	var jsonobject types.AppConfig

	err := json.Unmarshal(buf.Bytes(), &jsonobject)

	if err != nil {
		fmt.Println("error:", err)
	}

	return jsonobject
}
