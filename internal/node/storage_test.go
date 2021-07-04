package node

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
)

func TestLoadStoarage(t *testing.T) {
	var storageInfo map[string]interface{}
	file, err := os.Open("storage-info")
	if err != nil {
		t.Log(err)
	}
	context, err := ioutil.ReadAll(file)
	if err != nil {
		t.Log(err)
	}
	err = json.Unmarshal(context, &storageInfo)
	if err != nil {
		t.Log(err)
	}
	t.Log(storageInfo)
}
