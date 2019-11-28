<?xml version='1.0' encoding="UTF-8" standalone="no" ?>{{ $dataset := index . "dataset" }}
<WMS_Capabilities version="1.3.0"  xmlns="http://www.opengis.net/wms"   xmlns:sld="http://www.opengis.net/sld"   xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"   xmlns:ms="http://mapserver.gis.umn.edu/mapserver"   xmlns:inspire_common="http://inspire.ec.europa.eu/schemas/common/1.0"   xmlns:inspire_vs="http://inspire.ec.europa.eu/schemas/inspire_vs/1.0"   xsi:schemaLocation="http://www.opengis.net/wms http://schemas.opengis.net/wms/1.3.0/capabilities_1_3_0.xsd  http://www.opengis.net/sld http://schemas.opengis.net/sld/1.1.0/sld_capabilities.xsd  http://inspire.ec.europa.eu/schemas/inspire_vs/1.0  http://inspire.ec.europa.eu/schemas/inspire_vs/1.0/inspire_vs.xsd http://mapserver.gis.umn.edu/mapserver https://geodata.nationaalgeoregister.nl/{{ $dataset.Name}}/wms/v4_0?SERVICE=WMS&amp;language=dut&amp;service=WMS&amp;version=1.3.0&amp;request=GetSchemaExtension">
        {{ index . "service_provider.tpl" }}
        <Capability>
        {{ index . "capability_request.tpl" }}
        <Exception>
                    <Format>XML</Format>
                    <Format>INIMAGE</Format>
                    <Format>BLANK</Format>
                </Exception>
         {{ index . "inspire_common.tpl"}}
         {{ index . "layer.tpl" }} </Capability>
</WMS_Capabilities>
