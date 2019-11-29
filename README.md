# ogc-capabilities-generator

## TL;DR

```docker
docker build -t pdok/ogc-capabilities-generator
docker run -v `pwd`/config:/config -v `pwd`:/output -e SERVICE_TYPE=wfs -e SERVICE_VERSION=2.0.0 -e SERVICE_CONFIG_PATH=/config/input.yaml -e SERVICE_CAPABILITIES_PATH=/output/capabilities_wfs_200.xml pdok/ogc-capabilities-generator
```

## Config aanroep

```go
typePtr := flag.String("service-type", envString("SERVICE_TYPE", ""), "wfs or wms or wmts")
versionPtr := flag.String("service-version", envString("SERVICE_VERSION", ""), "wfs[1.1.0, 2.0.0] or wms[1.3.0] or wmts[1.0.0]")
serviceConfigPathPtr := flag.String("service-config", envString("SERVICE_CONFIG_PATH", ""), "location of the service config")
serviceDefConfigPathPtr := flag.String("service-def-config", envString("SERVICE_DEF_CONFIG_PATH", "config/serviceDef.yaml"), "location of the service definition config")
outputCapabilitiesPtr := flag.String("service-output-path", envString("SERVICE_CAPABILITIES_PATH", ""), "location of service config")
```

service definition file bevat alle lookups

input.yaml is een configuratie voor het maken van de capabilities

```yaml
organization: PDOK # ref staat in config/serviceDef
identification: kadastralekaart # ref staat in config/serviceDef
dataset: kadastralekaartv4 # ref staat in config/serviceDef
inspire: true
url: https://geodata.nationaalgeoregister.nl/kadastralekaart/wmts/v4_0?
boundingbox: WGS84 # ref staat in config/serviceDef
metadata-resource-id: service-metadata-id
features:
  - name: kadastralekaartv4:kadastralegrens
    title: kadastralegrens
    abstract: Een Kadastrale Grens is de weergave van een grens op de kadastrale kaart die door de dienst van
      het Kadaster tussen percelen vastgesteld wordt, op basis van inlichtingen van belanghebbenden en met
      gebruikmaking van de aan de kadastrale kaart ten grondslag liggende bescheiden die in elk geval de
      landmeetkundige gegevens bevatten van hetgeen op die kaart wordt weergegeven.
    keywords:
      - Kadastrale percelen
      - infoMapAccessService
    bounding-box: WGS84
    crs: # refs staat in config/serviceDef
      - 'EPSG:28992'
      - 'EPSG:3035'
      - 'EPSG:3038'
layers:
  - name: kadastralekaart
    title: kadastralegrens
    abstract: Een Kadastrale Grens is de weergave van een grens op de kadastrale kaart die door de dienst van
      het Kadaster tussen percelen vastgesteld wordt, op basis van inlichtingen van belanghebbenden en met
      gebruikmaking van de aan de kadastrale kaart ten grondslag liggende bescheiden die in elk geval de
      landmeetkundige gegevens bevatten van hetgeen op die kaart wordt weergegeven.
    keywords:
      - Kadastrale percelen
      - infoMapAccessService
    bounding-box: WGS84 # ref staat in config/serviceDef

    crs: # refs staat in config/serviceDef
      - 'EPSG:28992'
      - 'EPSG:25831'
      - 'EPSG:25832'
      - 'EPSG:3034'
      - 'EPSG:3035'
      - 'EPSG:3857'
      - 'EPSG:4258'
      - 'EPSG:4326'
      - 'CRS:84'
    srs: # refs staat in config/serviceDef
      - 'EPSG:28992'
      - 'EPSG:25831'
      - 'EPSG:25832'
      - 'EPSG:3034'
      - 'EPSG:3035'
      - 'EPSG:3857'
      - 'EPSG:4258'
      - 'EPSG:4326'
      - 'CRS:84'

    style_name: kadastralekaartv4:perceel_print
    style_title: Visualisatie van het perceel ten behoeve van afdrukken op 180 dpi.

```
