{{ $features:= index . "features" }}
{{ $dataset:= index . "dataset" }}
{{ $crsRef:= index . "crs" }}
{{ $boundingBoxRef:= index . "boundingbox" }}
 <FeatureTypeList  xmlns:ows="http://www.opengis.net/ows/1.1">
        <Operations>
            <Operation>Query</Operation>
        </Operations>
        {{range  $feature := $features}}
        <FeatureType>
            <Name>{{escape $feature.Name}}</Name>
            <Title>{{escape $feature.Title}}</Title>
            <Abstract>{{escape $feature.Abstract}}</Abstract>
           <ows:Keywords>{{range $keyword := $feature.Keywords}}
                 <ows:Keyword>{{ escape $keyword}}</ows:Keyword>{{end}}
               </ows:Keywords>{{range $crsIndex ,  $crs := $feature.Crs}}
               {{if isDefault $crsIndex}}<DefaultCRS>{{ (index $crsRef $crs).Value}}</DefaultCRS>{{else}}<OtherCRS>{{ (index $crsRef $crs).Value}}</OtherCRS>{{end}}{{end}}
             <OutputFormats>
                <Format>text/xml; subtype=gml/3.1.1</Format>
                <Format>application/json; subtype=geojson</Format>
                <Format>application/json</Format>
            </OutputFormats>
            <ows:WGS84BoundingBox dimensions="2">
             {{ $coords:=(index $boundingBoxRef $feature.Boundingbox).Coordinates}}<ows:LowerCorner>{{ index $coords 0}} {{index $coords 1 }}</ows:LowerCorner>
              <ows:UpperCorner>{{ index $coords 2}} {{index $coords 3 }}</ows:UpperCorner>
             </ows:WGS84BoundingBox>
            <MetadataURL format="text/plain" type="TC211">
                https://www.nationaalgeoregister.nl/geonetwork/srv/dut/xml.metadata.get?uuid={{ escape $dataset.MetadataResourceId}}
            </MetadataURL>
        </FeatureType>
        {{end}}
        </FeatureTypeList>