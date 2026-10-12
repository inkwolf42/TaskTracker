package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func getTaskJsonPath() string {
	currentDir, err := os.Getwd()
	PanicIfError(err)
	jsonpath := filepath.Join(currentDir, "tasks.json")
	return jsonpath
}
func Load() saveFile {
	jsonpath := getTaskJsonPath()
	dat, err := os.ReadFile(jsonpath)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return saveFile{
				Current_ID: 0,
				Task_list:  []task{},
			}

		} else {
			errMsg := fmt.Sprintf("Reading Error : %s\n", err)
			panic(errMsg)
		}
	}
	var result saveFile
	err = json.Unmarshal(dat, &result)
	PanicIfError(err)
	return result
}

func Save(data saveFile) {
	jsonpath := getTaskJsonPath()
	dat, err := json.MarshalIndent(data, "", "  ")
	PanicIfError(err)
	err = os.WriteFile(jsonpath, dat, 0644)
	PanicIfError(err)
}
