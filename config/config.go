package config

type Configuration struct {
	Server struct {
		KeepAlivePeriodSeconds int
		ListenAddr             string `default:""`
		Port                   int    `default:"8080"`

		SSL struct {
			Enabled         *bool  `default:"false"`
			RedirectToHTTPS *bool  `default:"true"`
			ListenAddr      string `default:""`
			Port            int    `default:"443"`
			CertFile        string `default:""`
			CertKey         string `default:""`
			LetsEncrypt     struct {
				Enabled   *bool  `default:"false"`
				AcceptTOS *bool  `default:"false"`
				Cache     string `default:"data/certs"`
				Hosts     []string
			}
		}
		ResponseHeaders map[string]string
		Stream          struct {
			PingPeriodSeconds int `default:"45"`
			AllowedOrigins    []string
		}
		Cors struct {
			AllowOrigins []string
			AllowMethods []string
			AllowHeaders []string
		}

		TrustedProxies []string
	}
	Database struct {
		Dialect    string `default:"postgres"`
		Connection string `default:"data/kingdom.db"`
	}
	DefaultUser struct {
		Name string `default:"admin"`
		Pass string `default:"admin"`
	}
	PassStrength      int    `default:"10"`
	UploadedImagesDir string `default:"data/images"`
	PluginsDir        string `default:"data/plugins"`
	Registration      bool   `default:"false"`
}

func Get() *Configuration {
	conf := new(Configuration)
	return conf
}
