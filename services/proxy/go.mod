module github.com/sazonovItas/proxy-manager/services/proxy

go 1.22.2

require (
	github.com/google/uuid v1.6.0
	github.com/sazonovItas/proxy-manager/pkg v0.0.0-00010101000000-000000000000
)

require (
	github.com/BurntSushi/toml v1.2.1 // indirect
	github.com/fatih/color v1.17.0 // indirect
	github.com/ilyakaznacheev/cleanenv v1.5.0 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/sazonovItas/go-simple-proxy v0.0.0-20240509164113-5aeb6e9f14db // indirect
	golang.org/x/sys v0.19.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	olympos.io/encoding/edn v0.0.0-20201019073823-d3554ca0b0a3 // indirect
)

replace github.com/sazonovItas/proxy-manager/pkg => ../../pkg
