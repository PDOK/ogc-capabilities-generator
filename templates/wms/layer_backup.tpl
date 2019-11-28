{{ $layers:= index . "layers" }}{{ $dataset:= index . "dataset" }}{{ $layerTypeRef:= index . "layertype" }} {{ $srsRef:= index . "srs" }} {{ $service:= index . "service" }}{{ $boundingBoxRef:= index . "boundingbox" }}{{if $service.Inspire}} <inspire_vs:ExtendedCapabilities>
                     <inspire_common:MetadataUrl xsi:type="inspire_common:resourceLocatorType">
                         <inspire_common:URL>http://www.nationaalgeoregister.nl/geonetwork/srv/eng/csw?service=CSW&amp;version=2.0.2&amp;request=GetRecordById&amp;outputschema=http://www.isotc211.org/2005/gmd&amp;elementsetname=full&amp;id={{ $service.MetadataResourceId}}</inspire_common:URL>
                         <inspire_common:MediaType>application/vnd.ogc.csw.GetRecordByIdResponse_xml</inspire_common:MediaType>
                     </inspire_common:MetadataUrl>
                     <inspire_common:SupportedLanguages>
                         <inspire_common:DefaultLanguage><inspire_common:Language>dut</inspire_common:Language></inspire_common:DefaultLanguage>
                     </inspire_common:SupportedLanguages>
                     <inspire_common:ResponseLanguage><inspire_common:Language>dut</inspire_common:Language></inspire_common:ResponseLanguage>
                 </inspire_vs:ExtendedCapabilities>{{ end}} {{range  $layer := $layers}}
          <Layer {{ (index $layerTypeRef $layer.Layertype).Value}}>
            <Name>{{escape $layer.Name}}</Name>
            <Title>{{escape $layer.Title}}</Title>
            <Abstract>{{escape $layer.Abstract}}</Abstract>
            <KeywordList>{{range $keyword := $layer.Keywords}}
                <Keyword>{{ escape $keyword}}</Keyword>{{end}}
            </KeywordList>{{range $crs := $layer.Crs}}
            <CRS>{{escape $crs}}</CRS>{{end}}
            <EX_GeographicBoundingBox>
            {{ $coords:=(index $boundingBoxRef $layer.Boundingbox).Coordinates}}  <westBoundLongitude>{{index $coords 0 }}</westBoundLongitude>
              <eastBoundLongitude>{{index $coords 2 }}</eastBoundLongitude>
              <southBoundLatitude>{{index $coords 1 }}</southBoundLatitude>
              <northBoundLatitude>{{index $coords 3 }}</northBoundLatitude>
            </EX_GeographicBoundingBox>{{range  $srs := $layer.Srs}}
            <BoundingBox {{ (index $srsRef $srs).Value}} />{{end}}
             <AuthorityURL name="Kadaster">
                      <OnlineResource xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="http://www.kadaster.nl"/>
                    </AuthorityURL>
                    <Identifier authority="Kadaster">nl.kad.4</Identifier>
                    <MetadataURL type="TC211">
               <Format>text/plain</Format>
               <OnlineResource xmlns:xlink="http://www.w3.org/1999/xlink" xlink:type="simple" xlink:href="https://www.nationaalgeoregister.nl/geonetwork/srv/dut/xml.metadata.get?uuid=a29917b9-3426-4041-a11b-69bcb2256904"/>
             </MetadataURL>{{range  $style := $layer.Styles}}
            <Style>
                <Name>{{escape $style.Name}}</Name>
                <Title>{{escape $style.Title}}</Title>
                <LegendURL width="183" height="169">
                    <Format>image/png</Format>
                    <OnlineResource xmlns:xlink="http://www.w3.org/1999/xlink" xlink:type="simple"
                                    xlink:href="{{escape $service.UrlWms }}?SERVICE=WMS&amp;language=dut&amp;version=1.3.0&amp;service=WMS&amp;request=GetLegendGraphic&amp;sld_version=1.1.0&amp;layer={{escape $layer.Name}}&amp;format=image/png&amp;STYLE={{escape $style.Name}}"/>
                </LegendURL>
            </Style>{{end}}
           </Layer>   {{end}}
