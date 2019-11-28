{{ $service:= index . "service" }}{{if $service.Inspire}}<inspire_vs:ExtendedCapabilities>
    <inspire_common:MetadataUrl xsi:type="inspire_common:resourceLocatorType">
      <inspire_common:URL>http://www.nationaalgeoregister.nl/geonetwork/srv/eng/csw?service=CSW&amp;version=2.0.2&amp;request=GetRecordById&amp;outputschema=http://www.isotc211.org/2005/gmd&amp;elementsetname=full&amp;id={{ $service.MetadataResourceId}}</inspire_common:URL>
      <inspire_common:MediaType>application/vnd.ogc.csw.GetRecordByIdResponse_xml</inspire_common:MediaType>
    </inspire_common:MetadataUrl>
    <inspire_common:SupportedLanguages>
      <inspire_common:DefaultLanguage><inspire_common:Language>dut</inspire_common:Language></inspire_common:DefaultLanguage>
    </inspire_common:SupportedLanguages>
    <inspire_common:ResponseLanguage><inspire_common:Language>dut</inspire_common:Language></inspire_common:ResponseLanguage>
  </inspire_vs:ExtendedCapabilities>{{end}}