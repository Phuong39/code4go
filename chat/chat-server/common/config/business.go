package config

import "github.com/spf13/viper"

type Business struct {
	KeepAlivePeriod uint

	EventCount uint

	Reply bool

	OnlineDataClearTime OnlineDataClearTime
}

type OnlineDataClearTime struct {
	Hour   string
	Minute string
	Second string
}

func InitBusiness(cfg *viper.Viper) *Business {
	m := cfg.Get("OnlineDataClearTime").(map[string]interface{})
	onlineDataClearTime := OnlineDataClearTime{
		Hour:   m["hour"].(string),
		Minute: m["minute"].(string),
		Second: m["second"].(string),
	}
	db := &Business{
		KeepAlivePeriod:     cfg.GetUint("KeepAlivePeriod"),
		EventCount:          cfg.GetUint("EventCount"),
		Reply:               cfg.GetBool("Reply"),
		OnlineDataClearTime: onlineDataClearTime,
	}
	return db
}

var BusinessConfig *Business = nil
