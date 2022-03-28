package config

type KongConfig struct {
	Path   string
	Reopen bool
}

type AppModuleConfig struct {
	Kong KongConfig
}

func (conf KongConfig) Log(kong *pdk.PDK)
