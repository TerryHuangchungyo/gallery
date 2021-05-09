package config

var GalleryDB = struct {
	DatabaseType string
	Host         string
	User         string
	Password     string
	DatabaseName string
}{
	"mysql",
	"gallery-db",
	"root",
	"root",
	"gallery",
}

var App = struct {
	Port string
}{
	":8080",
}
