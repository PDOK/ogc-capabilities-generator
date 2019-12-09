package builder

type Dataset struct {
	Name               string `yaml:"name"`
	MetadataResourceId string `yaml:"metadata-resource-id"`
}

type Filter struct {
}

type InspireCommon struct {
}

type Tilematrixset struct {
}

// type Crs struct {
// 	Name  int    `yaml:"name"`
// 	Value string `yaml:"value"`
// }

type Srs struct {
	Name  int    `yaml:"name"`
	Value string `yaml:"value"`
}

type Format struct {
	Name     string `yaml:"name"`
	MimeType string `yaml:"mime_type"`
}

// type Boundingbox struct {
// 	Coordinates []string `yaml:"coordinates"`
// 	Srs         int      `yaml:"srs"`
// }

type Operation struct {
	Key   string   `yaml:"key"`
	Value []string `yaml:"value"`
}
type Constraints struct {
	Name         string `yaml:"name"`
	DefaultValue bool   `yaml:"default-value"`
}

// type Service struct {
// 	Dataset            string        `yaml:"dataset"`
// 	Organization       string        `yaml:"organization"`
// 	Identification     string        `yaml:"identification"`
// 	Inspire            bool          `yaml:"inspire"`
// 	UrlWms             string        `yaml:"urlwms"`
// 	UrlWfs             string        `yaml:"urlwfs"`
// 	UrlWmts            string        `yaml:"urlwmts"`
// 	PathWMTS           string        `yaml:"pathwmts"`
// 	MetadataResourceId string        `yaml:"metadata-resource-id"`
// 	Operations         []string      `yaml:"operations"`
// 	Boundingbox        string        `yaml:"boundingbox"`
// 	Constraints        []Constraints `yaml:"constraints"`
// 	Features           []Feature     `yaml:"features"`
// 	Layers             []Layer       `yaml:"layers"`
// }

// type Identification struct {
// 	Title    string   `yaml:"title"`
// 	Abstract string   `yaml:"abstract"`
// 	Keywords []string `yaml:"keywords"`
// 	Version  string
// }

// type ServiceDef struct {
// 	Datasets       map[string]Dataset        `yaml:"datasets"`
// 	Organizations  map[string]Organization   `yaml:"organizations"`
// 	Identification map[string]Identification `yaml:"identification"`
// 	Crs            map[string]Crs            `yaml:"crs"`
// 	Srs            map[string]Srs            `yaml:"srs"`
// 	Format         []Format                  `yaml:"format"`
// 	Boundingbox    map[string]Boundingbox    `yaml:"boundingbox"`
// 	Operations     map[string]Operation      `yaml:"operations"`
// 	Constraints    []Constraints             `yaml:"constraints"`
// }
