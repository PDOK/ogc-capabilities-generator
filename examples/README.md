# README - Examples

Folders:

- `examples/spec_capabilities`: contains the example capabilities documents from
  the official OGC specification.
- `examples/capabilities`: altered versions of the official OGC specification
  examples, since the PDOK/OGC-Capabilities-Generator does not support all the
  features showcased in the official OGC specification examples. Examples with 
  INSPIRE extended capabilities can be found in the files with `_inspire` postfix.
- `examples/config`: configuration files that can be used to generate the
  altered capabilities documents (stored in `examples/capabilities`). These are
  used in the integrations tests (`integration_test.go`)

Missing examples:

- WCS 2.0.0 example, both regular and INSPIRE

To only run the tests generating the examples run:

```go
go test -run TestIntegrationW -v
```

> **NOTE:** For each example an integration test case should be present in the 
[integration_test.go](/integration_test.go) file.


## WMS 1.3.0

OGC Spec Example:
- [external OGC link](http://schemas.opengis.net/wms/1.3.0/capabilities_1_3_0.xml)
- [repo link](spec_capabilities/wms_capabilities_1_3_0.xml)

Not Implemented (but validates against spec):

- Layer attributes:
  - cascaded
  - noSubsets
  - fixedWidth
  - fixedHeight

## WFS 2.0

OGC Spec Example:
- [external OGC link](http://schemas.opengis.net/wfs/2.0/examples/GetCapabilities/GetCapabilities_Res_01.xml)
- [repo link](spec_capabilities/wfs_capabilities_2_0_0.xml)

Not Implemented (but validates against spec):

- FeatureType/ExtendedDescription
- Optional FeatureType/OutputFormats element. FeatureType/OutputFormats is
  always specified.

> TODO: Volgens NL Profiel WFS moet "Minimal Temporal Filter" ondersteund worden
> indien WFS data uitserveert met temporele aspecten. Uitzoeken of we dit
> ondersteunen. "Minimal Temporal Filter" wordt ondersteund door MapServer, dus
> vooralsnog zetten we dit op True
> (<https://www.geonovum.nl/uploads/standards/downloads/Nederlands%20profiel%20op%20ISO%2019142%20WFS%202%200%20-%20versie%201%201%20-%20definitief.pdf>),
> zie ook
> <https://github.com/MapServer/MapServer/blob/5dbbf8a19abb2f30b4852b72c8661faac14ae4c9/msautotest/wxs/expected/wfs_200_caps_sections_filter_capabilities.xml#L114>

## WMTS 1.0.0

OGC Spec Example:
- [external OGC link](http://schemas.opengis.net/wmts/1.0/examples/wmtsGetCapabilities_response.xml)
- [repo link](spec_capabilities/wmts_capabilities_1_0_0.xml)

Not Implemented (but validates against spec):

- Optional `Dimension` element not supported  (`Layer/Dimension`).

    ```xml
    <Dimension>
        <ows:Title>Time</ows:Title>
        <ows:Abstract>Monthly datasets</ows:Abstract>
        <ows:Identifier>TIME</ows:Identifier>
        <Value>2007-05</Value>
        <Value>2007-06</Value>
        <Value>2007-07</Value>
    </Dimension>
    ```

- Optional `Themes` element not supported (`Capabilities/Themes`).

Not tested in example (but available in spec):

Elements:

- Layer/ResourceURL
