package models

type GlobalConfig struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	// Could be int/float/bool depending on config structure
}
