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
