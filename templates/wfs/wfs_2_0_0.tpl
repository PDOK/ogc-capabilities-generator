<?xml version="1.0" encoding="UTF-8"?>{{ $dataset := index . "dataset" }}
<wfs:WFS_Capabilities version="2.0.0"
                      xmlns:gml="http://www.opengis.net/gml/3.2" xmlns:wfs="http://www.opengis.net/wfs/2.0"
                      xmlns:ows="http://www.opengis.net/ows/1.1" xmlns:xlink="http://www.w3.org/1999/xlink"
                      xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:fes="http://www.opengis.net/fes/2.0"
                      xmlns:{{ $dataset.Name}}="http://{{ $dataset.Name}}.geonovum.nl"
                      xmlns:inspire_common="http://inspire.ec.europa.eu/schemas/common/1.0"
                      xmlns:inspire_dls="http://inspire.ec.europa.eu/schemas/inspire_dls/1.0"
                      xmlns="http://www.opengis.net/wfs/2.0"
                      xsi:schemaLocation="http://www.opengis.net/wfs/2.0 http://schemas.opengis.net/wfs/2.0/wfs.xsd http://inspire.ec.europa.eu/schemas/inspire_dls/1.0 http://inspire.ec.europa.eu/schemas/inspire_dls/1.0/inspire_dls.xsd http://inspire.ec.europa.eu/schemas/common/1.0 http://inspire.ec.europa.eu/schemas/common/1.0/common.xsd">
    {{ index . "service-identification.tpl" }}
    {{ index . "service-provider.tpl" }}
    {{ index . "operations_metadata_2_0_0.tpl" }}
    {{ index . "feature_list_2_0_0.tpl" }}
    {{ index . "filter_capabilities_2_0_0.tpl" }}
</wfs:WFS_Capabilities>