package builder

type Config struct {
	Global  Global  `yaml:"global"`
	Service Service `yaml:"service"`
}

type Global struct {
	Onlineresourceurl string           `yaml:"onlineresourceurl"`
	Organization      Organization     `yaml:"organization"`
	WGS84Boundingbox  WGS84Boundingbox `yaml:"WGS84Boundingbox"`
	Title             string           `yaml:"title"`
	Abstract          string           `yaml:"abstract"`
	Path              string           `yaml:"path"`
}

type WGS84Boundingbox struct {
	WestBoundLongitude float64 `yaml:"westBoundLongitude"`
	EastBoundLongitude float64 `yaml:"eastBoundLongitude"`
	SouthBoundLatitude float64 `yaml:"southBoundLatitude"`
	NorthBoundLatitude float64 `yaml:"northBoundLatitude"`
}

type Organization struct {
	Name       string `yaml:"name"`
	Individual string `yaml:"individual"`
	Position   string `yaml:"position"`
	City       string `yaml:"city"`
	Country    string `yaml:"country"`
	Email      string `yaml:"email"`
}

type Service struct {
	WFS []WFS `yaml:"wfs"`
	WMS []WMS `yaml:"wms"`
	// WMTS []WMTS `yaml:"wmts"`
}

type WFS struct {
	CRS          CRS            `yaml:"crs"`
	Version      string         `yaml:"version"`
	Outputformat []Outputformat `yaml:"outputformat"`
	Features     []Feature      `yaml:"features"`
}

type WMS struct {
	CRS          CRS            `yaml:"crs"`
	Version      string         `yaml:"version"`
	Boundingbox  []Boundingbox  `yaml:"boundingbox"`
	Outputformat []Outputformat `yaml:"outputformat"`
	Layers       []Layer        `yaml:"layers"`
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
	Name     string   `yaml:"name"`
	Title    string   `yaml:"title"`
	Abstract string   `yaml:"abstract"`
	Keywords []string `yaml:"keywords"`
}

type Boundingbox struct {
	Crs  string  `yaml:"crs"`
	Minx float64 `yaml:"minx"`
	Miny float64 `yaml:"miny"`
	Maxx float64 `yaml:"maxx"`
	Maxy float64 `yaml:"maxy"`
}

type Layer struct {
	Name        string   `yaml:"name"`
	Title       string   `yaml:"title", omitempty`
	Abstract    string   `yaml:"abstract", omitempty`
	Keywords    []string `yaml:"keywords", omitempty`
	Boundingbox string   `yaml:"bounding-box", omitempty`
	Crs         []string `yaml:"crs", omitempty`
	Srs         []string `yaml:"srs", omitempty`
	Styles      []Style  `yaml:"styles", omitempty`
	Layers      []Layer  `yaml:"layers", omitempty`
	Nested      string
}

type Style struct {
	Name  string `yaml:"name"`
	Title string `yaml:"title"`
}
