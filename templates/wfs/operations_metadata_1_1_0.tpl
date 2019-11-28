{{ $service:= index . "service" }}{{ $dataset:= index . "dataset" }}<ows:OperationsMetadata>
        <ows:Operation name="GetCapabilities">
            <ows:DCP>
                <ows:HTTP>
                    <ows:Get xlink:type="simple" xlink:href="{{ $service.UrlWfs }}?SERVICE=WFS&amp"/>
                    <ows:Post xlink:type="simple" xlink:href="{{ $service.UrlWfs }}?SERVICE=WFS&amp;"/>
                </ows:HTTP>
            </ows:DCP>
            <ows:Parameter name="service">
                <ows:Value>WFS</ows:Value>
            </ows:Parameter>
            <ows:Parameter name="AcceptVersions">
                <ows:Value>1.1.0</ows:Value>
            </ows:Parameter>
            <ows:Parameter name="AcceptFormats">
                <ows:Value>text/xml</ows:Value>
            </ows:Parameter>
        </ows:Operation>
        <ows:Operation name="DescribeFeatureType">
            <ows:DCP>
                <ows:HTTP>
                    <ows:Get xlink:type="simple" xlink:href="{{ $service.UrlWfs }}?SERVICE=WFS&amp;"/>
                    <ows:Post xlink:type="simple" xlink:href="{{ $service.UrlWfs }}?SERVICE=WFS&amp;"/>
                </ows:HTTP>
            </ows:DCP>
            <ows:Parameter name="outputFormat">
                <ows:Value>XMLSCHEMA</ows:Value>
                <ows:Value>text/xml; subtype=gml/2.1.2</ows:Value>
                <ows:Value>text/xml; subtype=gml/3.1.1</ows:Value>
            </ows:Parameter>
        </ows:Operation>
        <ows:Operation name="GetFeature">
            <ows:DCP>
                <ows:HTTP>
                    <ows:Get xlink:type="simple" xlink:href="{{ $service.UrlWfs }}?SERVICE=WFS&amp;"/>
                    <ows:Post xlink:type="simple" xlink:href="{{ $service.UrlWfs }}?SERVICE=WFS&amp;"/>
                </ows:HTTP>
            </ows:DCP>
            <ows:Parameter name="resultType">
                <ows:Value>results</ows:Value>
                <ows:Value>hits</ows:Value>
            </ows:Parameter>
            <ows:Parameter name="outputFormat">
                <ows:Value>text/xml; subtype=gml/3.1.1</ows:Value>
                <ows:Value>application/json; subtype=geojson</ows:Value>
                <ows:Value>application/json</ows:Value>
            </ows:Parameter>
        </ows:Operation>
        <ows:Constraint name="DefaultMaxFeatures">
            <ows:Value>1000</ows:Value>
        </ows:Constraint>
    </ows:OperationsMetadata>