<?xml version="1.0" encoding="UTF-8"?>{{ $dataset := index . "dataset" }}
<wfs:WFS_Capabilities version="1.1.0"
                      xmlns:gml="http://www.opengis.net/gml"
                      xmlns:wfs="http://www.opengis.net/wfs"
                      xmlns:ows="http://www.opengis.net/ows"
                      xmlns:xlink="http://www.w3.org/1999/xlink"
                      xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
                      xmlns:ogc="http://www.opengis.net/ogc"
                      xmlns:{{ $dataset.Name}}="http://{{$dataset.Name}}.geonovum.nl"
                      xmlns="http://www.opengis.net/wfs"  xsi:schemaLocation="http://www.opengis.net/wfs http://schemas.opengis.net/wfs/1.1.0/wfs.xsd">
    {{ index . "service-identification.tpl" }}
    {{ index . "service-provider.tpl" }}
    {{ index . "operations_metadata_1_1_0.tpl" }}
    {{ index . "feature_list_1_1_0.tpl" }}
    {{ index . "filter_capabilities_1_1_0.tpl" }}
</wfs:WFS_Capabilities>
