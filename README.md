# ogc-capabilities-generator

## TL;DR

```docker
docker build -t pdok/ogc-capabilities-generator .
docker run -d -v `pwd`/examples:/config -v `pwd`:/output -e SERVICECONFIG=/config/natura2000.yaml --name ocg pdok/ogc-capabilities-generator
```

https://geodata.nationaalgeoregister.nl/kadastralekaart/wms/v4_0?request=getcapabilities&service=wms
https://geodata.nationaalgeoregister.nl/kadastralekaart/wmts/v4_0?request=GetCapabilities&service=WMTS
https://geodata.nationaalgeoregister.nl/kadastralekaart/wmts/v4_0/WMTSCapabilities.xml
https://geodata.nationaalgeoregister.nl/kadastralekaart/wfs/v4_0?request=getcapabilities&service=wfs&version=2.0.0
https://geodata.nationaalgeoregister.nl/kadastralekaart/wfs/v4_0?request=getcapabilities&service=wfs&version=1.1.0

https://www.onlinetool.io/xmltogo/

| | wms 1.3.0 | wmts 1.0.0 | wfs 1.1.0 | wfs 2.0.0 |
|---|---|---|---|---|
| Service | X | - | - | - |
| Capability | X | - | - | - |
| ServiceIdentification | - | X | X | X |
| ServiceProvider | - | - | X | X |
| OperationsMetadata | - | X | X | X |
| FeatureTypeList | - | - | X | X |
| Filter_Capabilities | - | - | X (ogc) | X (fes) |
| ServiceMetadataURL | - | X | - | - |

## Config aanroep

```go
go run . -c ./examples/kadastralekaart.yaml
go run . -c ./examples/natura2000.yaml
```
