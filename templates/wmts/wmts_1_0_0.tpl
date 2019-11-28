<?xml version="1.0"?>
<Capabilities xmlns="http://www.opengis.net/wmts/1.0" xmlns:ows="http://www.opengis.net/ows/1.1" xmlns:xlink="http://www.w3.org/1999/xlink" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:gml="http://www.opengis.net/gml" xsi:schemaLocation="http://www.opengis.net/wmts/1.0 http://schemas.opengis.net/wmts/1.0/wmtsGetCapabilities_response.xsd" version="1.0.0">
    <ows:ServiceIdentification>
        <ows:Title></ows:Title>
        <ows:Abstract></ows:Abstract>
        <ows:ServiceType>OGC WMTS</ows:ServiceType>
        <ows:ServiceTypeVersion>1.0.0</ows:ServiceTypeVersion>
        <ows:Fees>none</ows:Fees>
        <ows:AccessConstraints>none</ows:AccessConstraints>
    </ows:ServiceIdentification>
    {{index . "operations_metadata.tpl" }}
    <Contents>
    {{index . "layer.tpl" }}
    {{ index . "tilematrixset_28992.tpl" }}
    {{ index . "tilematrixset_3857.tpl" }}
    {{ index . "tilematrixset_25831.tpl" }}
  </Contents>
    <ServiceMetadataURL xlink:href="https://geodata.nationaalgeoregister.nl/kadastralekaart/wmts/v4_0/WMTSCapabilities.xml"/>
</Capabilities>
