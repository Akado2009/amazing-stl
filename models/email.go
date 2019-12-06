package models

//Email struct
type Email struct {
	Address string `toml:"Address"`
	Login string `toml:"Login"`
	Password string `toml:"Password"`
	Port int `toml:"Port"`
}