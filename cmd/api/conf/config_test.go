package conf

import "testing"

func TestGetConfig(t *testing.T) {
	config := GetConfig()
	t.Log(config)
}
