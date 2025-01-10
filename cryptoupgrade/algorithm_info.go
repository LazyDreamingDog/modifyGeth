package cryptoupgrade

import (
	"encoding/json"
	"os"
	"strings"
)

var algoInfoMap = make(map[string]algoInfo)

type algoInfo struct {
	code  string
	gas   uint64
	itype string
	otype string
}

func (c *algoInfo) getTypeList() ([]string, []string) {
	ilist := strings.Split(c.itype, ",")
	olist := strings.Split(c.otype, ",")
	return ilist, olist
}

// When geth exit, need to store @algoInfoMap
func Store() error {
	return storeAlgoMap(algoInfoPath)
}

func storeAlgoMap(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			// File is not exsit, ignore this error
			return nil
		}
		return err
	}

	return json.Unmarshal(file, &algoInfoMap)
}

func loadFromFile(filename string) error {
	data, err := json.MarshalIndent(algoInfoMap, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}
