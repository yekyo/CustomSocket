package cmd

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
)

func clearConfigFile()  {
	os.Remove(configPath)
}

func TestConfig_SaveConfig(t *testing.T) {
	clearConfigFile()
	config := Config{
		ListenAddr: ":7474",
		RemoteAddr: "do.kyoye.com:10081",
		Password:	"eJ40dXsaJ1ClCkKEwZWRt8jvBYK0THoDqSCOH6zLjeOFO3g4TS88rMpBNqqwv8t0s1S+TiSvvQ8pKttXLni32a/2v12mrL+uPh",
	}
	config.SaveConfig()

	file, err := os.Open(configPath)
	if err != nil {
		t.Errorf("打开配置文件 %s 出错: %s", configPath, err)
	}

	defer file.Close()

	tmp := make(map[string]string)
	err = json.NewDecoder(file).Decode(&tmp)
	if err != nil {
		t.Error(err)
	}

	if tmp["listen"] != ":7474" || tmp["password"] != "eJ40dXsaJ1ClCkKEwZWRt8jvBYK0THoDqSCOH6zLjeOFO3g4TS88rMpBNqqwv8t0s1S+TiSvvQ8pKttXLni32a/2v12mrL+uPh" {
		t.Error("数据保存不一致")

	}
}

func TestConfig_ReadConfig(t *testing.T) {
	clearConfigFile()
	jsonConfig := []byte(`
{
	"listen": ":7474",
	"remote": "do.kyoye.com:10081",
	"password": "kyoyecGfU96a9o/sPM0bJ6xRPJiQCOSmfDF9UYNsAKhsxgfK2HWS1oAvQDXM10fVX5oz8g0Rr5WP2Wk2kwUT6Z7o90w5SJfpxJFnI9+oCwcZYY2QiIf1YixPMSEd2bs4XHk2IERdfc8RcofF6Hu6+Y6LoB1eyxeJ40dXsaJ1ClCkKEwZWRt8jvBYK0THoDqSCOH6zLjeOFO3g4TS88rMpBNqqwv8t0s1S+TiSvvQ8pKttXLni32a/2v12mrL+uPh"
}`)

	err := ioutil.WriteFile(configPath, jsonConfig, 0644)
	if err != nil {
		t.Error(err)
	}

	config := Config{}
	config.ReadConfig()

	if config.ListenAddr != ":7474" {
		t.Error("listen addr 不一致")
	}

	if config.RemoteAddr != "do.kyoye.com:10081" {
		t.Error("remote addr 不一致")
	}

	if config.Password != "kyoyecGfU96a9o/sPM0bJ6xRPJiQCOSmfDF9UYNsAKhsxgfK2HWS1oAvQDXM10fVX5oz8g0Rr5WP2Wk2kwUT6Z7o90w5SJfpxJFnI9+oCwcZYY2QiIf1YixPMSEd2bs4XHk2IERdfc8RcofF6Hu6+Y6LoB1eyxeJ40dXsaJ1ClCkKEwZWRt8jvBYK0THoDqSCOH6zLjeOFO3g4TS88rMpBNqqwv8t0s1S+TiSvvQ8pKttXLni32a/2v12mrL+uPh" {
		t.Error("密码不一致")
	}

	if config.ListenAddr != ":7474" || config.RemoteAddr != "do.kyoye.com:10081" || config.Password != "kyoyecGfU96a9o/sPM0bJ6xRPJiQCOSmfDF9UYNsAKhsxgfK2HWS1oAvQDXM10fVX5oz8g0Rr5WP2Wk2kwUT6Z7o90w5SJfpxJFnI9+oCwcZYY2QiIf1YixPMSEd2bs4XHk2IERdfc8RcofF6Hu6+Y6LoB1eyxeJ40dXsaJ1ClCkKEwZWRt8jvBYK0THoDqSCOH6zLjeOFO3g4TS88rMpBNqqwv8t0s1S+TiSvvQ8pKttXLni32a/2v12mrL+uPh" {
		t.Error("读取数据不一致")
	}
}
