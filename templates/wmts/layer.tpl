{{ $layers:= index . "layers" }}{{ $boundingBoxRef:= index . "boundingbox" }}{{range  $layer := $layers}} <Layer xmlns:ows="http://www.opengis.net/ows/1.1">
<ows:Title>{{ escape $layer.Name}}</ows:Title>
<ows:Abstract></ows:Abstract>
<ows:WGS84BoundingBox> {{ $coords:=(index $boundingBoxRef $layer.Boundingbox).Coordinates}}
<ows:LowerCorner>{{ index $coords 0}} {{index $coords 1 }}</ows:LowerCorner>
<ows:UpperCorner>{{ index $coords 2}} {{index $coords 3 }}</ows:UpperCorner>
</ows:WGS84BoundingBox>
<ows:Identifier>{{escape $layer.Name}}</ows:Identifier>
<Style>
 <ows:Identifier>default</ows:Identifier>
</Style>
<Format>image/png</Format>
<TileMatrixSetLink>
 <TileMatrixSet>EPSG:28992</TileMatrixSet>
</TileMatrixSetLink>
<TileMatrixSetLink>
 <TileMatrixSet>EPSG:25831</TileMatrixSet>
 </TileMatrixSetLink>
<TileMatrixSetLink>
   <TileMatrixSet>EPSG:3857</TileMatrixSet>
</TileMatrixSetLink>
</Layer>{{end}}
