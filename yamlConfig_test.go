package yamlConfig

import "testing"

func TestConfigEngine_Load(t *testing.T) {
	config := &ConfigEngine{}
	err := config.Load("config.yaml")
	if err != nil {
		t.Fatalf("config load error:%v",err)
	}
	t.Logf("%v", config)
	res := config.Get("mysql.username")
	t.Log(res)
}