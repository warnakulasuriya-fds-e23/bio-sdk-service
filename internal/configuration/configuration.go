package configuration

type Configuration struct {
	ImagesDir     string `toml:"imagesdir"`
	CborDir       string `toml:"cbordir"`
	StorageVolume string `toml:"storagevolume"`
}
