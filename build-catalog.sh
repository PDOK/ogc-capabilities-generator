#!/usr/bin/env sh
rm xml-catalog/ogc-catalog.xml
xmlcatalog --create > xml-catalog/ogc-catalog.xml
xmlcatalog --noout --add system http://schemas.opengis.net/wfs/2.0/wfs.xsd schemas/schemas.opengis.net/wfs/2.0/wfs.xsd xml-catalog/ogc-catalog.xml
xmlcatalog --noout --add system http://schemas.opengis.net/wms/1.3.0/capabilities_1_3_0.xsd schemas/schemas.opengis.net/wms/1.3.0/capabilities_1_3_0.xsd xml-catalog/ogc-catalog.xml
xmlcatalog --noout --add system http://schemas.opengis.net/sld/1.1.0/sld_capabilities.xsd schemas/schemas.opengis.net/sld/1.1.0/sld_capabilities.xsd xml-catalog/ogc-catalog.xml
xmlcatalog --noout --add system http://inspire.ec.europa.eu/schemas/inspire_vs/1.0/inspire_vs.xsd schemas/inspire.ec.europa.eu/schemas/inspire_vs/1.0/inspire_vs.xsd xml-catalog/ogc-catalog.xml
xmlcatalog --noout --add system http://inspire.ec.europa.eu/schemas/inspire_dls/1.0/inspire_dls.xsd schemas/inspire.ec.europa.eu/schemas/inspire_dls/1.0/inspire_dls.xsd xml-catalog/ogc-catalog.xml
xmlcatalog --noout --add system http://schemas.opengis.net/wmts/1.0/wmts.xsd schemas/schemas.opengis.net/wmts/1.0/wmts.xsd xml-catalog/ogc-catalog.xml
