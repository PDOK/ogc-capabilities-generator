  <Layer queryable="1" opaque="0" cascaded="0">
    <Name>{{ .Layer.Name}}</Name>
    <Title>{{ .Layer.Title}}</Title>
    <Abstract>{{ .Layer.Abstract}}</Abstract>
    <KeywordList>{{range $keyword := .Layer.Keywords}}
        <Keyword>{{ $keyword}}</Keyword>{{end}}
        <Keyword vocabulary="ISO">infoMapAccessService</Keyword>
    </KeywordList>{{range $crs := $.Layer.Crs}}
    <CRS>{{$crs}}</CRS>{{end}}
   <EX_GeographicBoundingBox>{{ $boundingBoxRef:= index .ServiceDef.Boundingbox .Layer.Boundingbox }}{{ $coords:= $boundingBoxRef.Coordinates}}
        <westBoundLongitude>{{index $coords 0 }}</westBoundLongitude>
        <eastBoundLongitude>{{index $coords 2}}</eastBoundLongitude>
        <southBoundLatitude>{{index $coords 1}}</southBoundLatitude>
        <northBoundLatitude>{{index $coords 3}}</northBoundLatitude>
    </EX_GeographicBoundingBox>{{ $srsRef:= index  .ServiceDef.Srs }}{{range $srs := .Layer.Srs}}
    <BoundingBox {{ (index $srsRef $srs).Value}}/>{{end}}
    <AuthorityURL name="Kadaster">
     <OnlineResource xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="http://www.kadaster.nl"/>
    </AuthorityURL>
    <Identifier authority="Kadaster">nl.kad.4</Identifier>
        <MetadataURL type="TC211">
       <Format>text/plain</Format>{{ $metaDataRef:= index .ServiceDef.Datasets .Service.Dataset }}{{ $uuid:= $metaDataRef.MetadataResourceId}}
      <OnlineResource xmlns:xlink="http://www.w3.org/1999/xlink" xlink:type="simple" xlink:href="https://www.nationaalgeoregister.nl/geonetwork/srv/dut/xml.metadata.get?uuid={{$uuid }}"/>
       </MetadataURL>
       <Style>{{range  $style := .Layer.Styles}}
       <Name>{{escape $style.Name}}</Name>
       <Title>{{escape $style.Title}}</Title>
       <LegendURL width="183" height="169">
          <Format>image/png</Format>
          <OnlineResource xmlns:xlink="http://www.w3.org/1999/xlink" xlink:type="simple" xlink:href="https://geodata.nationaalgeoregister.nl/kadastralekaart/wms/v4_0?SERVICE=WMS&amp;language=dut&amp;version=1.3.0&amp;service=WMS&amp;request=GetLegendGraphic&amp;sld_version=1.1.0&amp;layer=kadastralekaart&amp;format=image/png&amp;STYLE={{escape $style.Name}}"/>
       </LegendURL>{{end}}
    </Style>
   {{.Layer.Nested}}
 </Layer>