module ogc-capabilities-generator

go 1.14

replace (
	github.com/pdok/ogc-specifications => /home/mark/repos/ogc-specifications
	github.com/pdok/ogc-specifications/pkg/wcs201 => /home/mark/repos/ogc-specifications/pkg/wcs201
	github.com/pdok/ogc-specifications/pkg/wfs200 => /home/mark/repos/ogc-specifications/pkg/wfs200
	github.com/pdok/ogc-specifications/pkg/wms130 => /home/mark/repos/ogc-specifications/pkg/wms130
	github.com/pdok/ogc-specifications/pkg/wmts100 => /home/mark/repos/ogc-specifications/pkg/wmts100
)

require (
	github.com/ajankovic/xdiff v0.0.1
	github.com/imdario/mergo v0.3.11
	github.com/pdok/ogc-specifications v0.2.3
	gopkg.in/yaml.v2 v2.3.0
)
