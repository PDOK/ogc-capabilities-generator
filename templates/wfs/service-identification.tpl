<ows:ServiceIdentification xmlns:ows="http://www.opengis.net/ows/1.1">
<ows:Title>{{escape .Title}}</ows:Title>
<ows:Abstract>{{escape  .Abstract}}</ows:Abstract>
<ows:Keywords>{{range $i, $keyword := $.Keywords}}
  <ows:Keyword>{{escape $keyword}}</ows:Keyword>{{end}}
</ows:Keywords>
<ows:ServiceType codeSpace="OGC">WFS</ows:ServiceType>
<ows:ServiceTypeVersion>{{escape .Version}}</ows:ServiceTypeVersion>
<ows:Fees>NONE</ows:Fees>
<ows:AccessConstraints>http://creativecommons.org/publicdomain/zero/1.0/deed.nl</ows:AccessConstraints>
</ows:ServiceIdentification>