package verso

type Changelog struct {
	Path     string   `json:"path"`
	Versions []Semver `json:"versions"`
}
