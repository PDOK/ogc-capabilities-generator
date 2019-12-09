package builder

type Config struct {
	Global  Global  `yaml:"global"`
	Service Service `yaml:"service"`
}

type Global struct {
	Organization Organization `yaml:"organization"`
	Boundingbox  []float64    `yaml:"boundingbox"`
	Title        string       `yaml:"title"`
	Abstract     string       `yaml:"abstract"`
	Path         string       `yaml:"path"`
}

type Organization struct {
	Name       string `yaml:"name"`
	Individual string `yaml:"individual"`
	Position   string `yaml:"position"`
	City       string `yaml:"city"`
	Country    string `yaml:"country"`
	Email      string `yaml:"e-mail"`
}

type Service struct {
	WFS []WFS `yaml:"wfs"`
	// WMS  []WMS  `yaml:"wms"`
	// WMTS []WMTS `yaml:"wmts"`
}

type WFS struct {
	CRS          CRS            `yaml:"crs"`
	Version      string         `yaml:"version"`
	Outputformat []Outputformat `yaml:"outputformat"`
	Features     []Feature      `yaml:"features"`
}

type CRS struct {
	Default string   `yaml:"default"`
	Other   []string `yaml:"other"`
}

type Outputformat struct {
	Name     string `yaml:"name"`
	Mimetype string `yaml:"mime_type"`
}

type Feature struct {
	Name        string    `yaml:"name"`
	Title       string    `yaml:"title"`
	Abstract    string    `yaml:"abstract"`
	Boundingbox []float64 `yaml:"boundingbox"`
	Keywords    []string  `yaml:"keywords"`
}
