#!/usr/bin/env sh
start=`date +%s`

# debug all the things:
# export XML_DEBUG_CATALOG=
export XML_CATALOG_FILES=xml-catalog/ogc-catalog.xml

success_state=SUCCESS

echo "Starting catalog test"
# Validation should be fast and be able to run without an internet connection.
echo
echo "### test WFS"
xmllint --load-trace --noout --schema http://schemas.opengis.net/wfs/2.0/wfs.xsd examples/capabilities/wfs_capabilities_2_0_0.xml
if ! [ $? -eq 0 ]
then
  success_state=FAILURE
fi
echo
echo "### test WFS Inspire"
xmllint --load-trace --noout --schema examples/xsds/wfs_2_0_0_inspire.xsd examples/capabilities/wfs_capabilities_2_0_0_inspire.xml
if ! [ $? -eq 0 ]
then
  success_state=FAILURE
fi
echo
echo "### test WMS"
xmllint --load-trace --noout --schema http://schemas.opengis.net/wms/1.3.0/capabilities_1_3_0.xsd examples/capabilities/wms_capabilities_1_3_0.xml
if ! [ $? -eq 0 ]
then
  success_state=FAILURE
fi
echo
echo "### test WMS Inspire"
xmllint --load-trace --noout --schema examples/xsds/wms_1_3_0_inspire.xsd examples/capabilities/wms_capabilities_1_3_0_inspire.xml
if ! [ $? -eq 0 ]
then
  success_state=FAILURE
fi
echo
echo "## test WMTS"
xmllint --load-trace --noout --schema http://schemas.opengis.net/wmts/1.0/wmts.xsd examples/capabilities/wmts_capabilities_1_0_0.xml
if ! [ $? -eq 0 ]
then
  success_state=FAILURE
fi

end=`date +%s`
runtime=$((end-start))

if [ "$runtime" -gt "0" ]
then
  success_state=FAILURE
fi

echo
echo test $success_state: finished in $runtime seconds