{{ $service:= index . "service" }}{{ $dataset:= index . "dataset" }}<ows:OperationsMetadata xmlns="http://www.opengis.net/wfs/2.0"
                        xmlns:ows="http://www.opengis.net/ows/1.1"
                        xmlns:wfs="http://www.opengis.net/wfs/2.0"
                        xmlns:inspire_common="http://inspire.ec.europa.eu/schemas/common/1.0"
                        xmlns:inspire_dls="http://inspire.ec.europa.eu/schemas/inspire_dls/1.0">
    <ows:Operation name="GetCapabilities">
        <ows:DCP>
            <ows:HTTP>
                <ows:Get xlink:type="simple" xlink:href="{{ $service.UrlWfs }}?SERVICE=WFS&amp;language=eng&amp;"/>
                <ows:Post xlink:type="simple" xlink:href="{{ $service.UrlWfs }}?SERVICE=WFS&amp;language=eng&amp;"/>
            </ows:HTTP>
        </ows:DCP>
        <ows:Parameter name="AcceptVersions">
            <ows:AllowedValues>
                <ows:Value>2.0.0</ows:Value>
                <ows:Value>1.1.0</ows:Value>
            </ows:AllowedValues>
        </ows:Parameter>
        <ows:Parameter name="AcceptFormats">
            <ows:AllowedValues>
                <ows:Value>text/xml</ows:Value>
            </ows:AllowedValues>
        </ows:Parameter>
        <ows:Parameter name="Sections">
            <ows:AllowedValues>
                <ows:Value>ServiceIdentification</ows:Value>
                <ows:Value>ServiceProvider</ows:Value>
                <ows:Value>OperationsMetadata</ows:Value>
                <ows:Value>FeatureTypeList</ows:Value>
                <ows:Value>Filter_Capabilities</ows:Value>
            </ows:AllowedValues>
        </ows:Parameter>
    </ows:Operation>
    <ows:Operation name="DescribeFeatureType">
        <ows:DCP>
            <ows:HTTP>
                <ows:Get xlink:type="simple" xlink:href="{{ $service.UrlWfs }}?SERVICE=WFS&amp;language=eng&amp;"/>
                <ows:Post xlink:type="simple" xlink:href="{{ $service.UrlWfs }}?SERVICE=WFS&amp;language=eng&amp;"/>
            </ows:HTTP>
        </ows:DCP>
        <ows:Parameter name="outputFormat">
            <ows:AllowedValues>
                <ows:Value>application/gml+xml; version=3.2</ows:Value>
                <ows:Value>text/xml; subtype=gml/3.2.1</ows:Value>
                <ows:Value>text/xml; subtype=gml/3.1.1</ows:Value>
                <ows:Value>text/xml; subtype=gml/2.1.2</ows:Value>
                <ows:Value>application/json; subtype=geojson</ows:Value>
                <ows:Value>application/json</ows:Value>
                <ows:Value>text/xml</ows:Value>
            </ows:AllowedValues>
        </ows:Parameter>
    </ows:Operation>
    <ows:Operation name="GetFeature">
        <ows:DCP>
            <ows:HTTP>
                <ows:Get xlink:type="simple" xlink:href="{{ $service.UrlWfs }}?SERVICE=WFS&amp;language=eng&amp;"/>
                <ows:Post xlink:type="simple" xlink:href="{{ $service.UrlWfs }}?SERVICE=WFS&amp;language=eng&amp;"/>
            </ows:HTTP>
        </ows:DCP>
        <ows:Parameter name="outputFormat">
            <ows:AllowedValues>
                <ows:Value>application/gml+xml; version=3.2</ows:Value>
                <ows:Value>text/xml; subtype=gml/3.2.1</ows:Value>
                <ows:Value>text/xml; subtype=gml/3.1.1</ows:Value>
                <ows:Value>text/xml; subtype=gml/2.1.2</ows:Value>
                <ows:Value>application/json; subtype=geojson</ows:Value>
                <ows:Value>application/json</ows:Value>
            </ows:AllowedValues>
        </ows:Parameter>
    </ows:Operation>
    <ows:Operation name="GetPropertyValue">
        <ows:DCP>
            <ows:HTTP>
                <ows:Get xlink:type="simple" xlink:href="{{ $service.UrlWfs }}?SERVICE=WFS&amp;language=eng&amp;"/>
                <ows:Post xlink:type="simple" xlink:href="{{ $service.UrlWfs }}?SERVICE=WFS&amp;language=eng&amp;"/>
            </ows:HTTP>
        </ows:DCP>
        <ows:Parameter name="outputFormat">
            <ows:AllowedValues>
                <ows:Value>application/gml+xml; version=3.2</ows:Value>
                <ows:Value>text/xml; subtype=gml/3.2.1</ows:Value>
                <ows:Value>text/xml; subtype=gml/3.1.1</ows:Value>
                <ows:Value>text/xml; subtype=gml/2.1.2</ows:Value>
                <ows:Value>application/json; subtype=geojson</ows:Value>
                <ows:Value>application/json</ows:Value>
            </ows:AllowedValues>
        </ows:Parameter>
    </ows:Operation>
{{if $service.Inspire}}
    <ows:Operation name="ListStoredQueries">
        <ows:DCP>
            <ows:HTTP>
                <ows:Get xlink:type="simple" xlink:href="{{ escape $service.UrlWfs }}?SERVICE=WFS&amp;language=eng&amp;"/>
                <ows:Post xlink:type="simple" xlink:href="{{ escape $service.UrlWfs }}?SERVICE=WFS&amp;language=eng&amp;"/>
            </ows:HTTP>
        </ows:DCP>
    </ows:Operation>
    <ows:Operation name="DescribeStoredQueries">
        <ows:DCP>
            <ows:HTTP>
                <ows:Get xlink:type="simple" xlink:href="{{ $service.UrlWfs }}?SERVICE=WFS&amp;language=eng&amp;"/>
                <ows:Post xlink:type="simple" xlink:href="{{ $service.UrlWfs }}?SERVICE=WFS&amp;language=eng&amp;"/>
            </ows:HTTP>
        </ows:DCP>
    </ows:Operation>
{{ end }}
    <ows:Parameter name="version">
        <ows:AllowedValues>
            <ows:Value>2.0.0</ows:Value>
            <ows:Value>1.1.0</ows:Value>
        </ows:AllowedValues>
    </ows:Parameter>
    <ows:Constraint name="ImplementsBasicWFS">
        <ows:NoValues/>
        <ows:DefaultValue>TRUE</ows:DefaultValue>
    </ows:Constraint>
    <ows:Constraint name="ImplementsTransactionalWFS">
        <ows:NoValues/>
        <ows:DefaultValue>FALSE</ows:DefaultValue>
    </ows:Constraint>
    <ows:Constraint name="ImplementsLockingWFS">
        <ows:NoValues/>
        <ows:DefaultValue>FALSE</ows:DefaultValue>
    </ows:Constraint>
    <ows:Constraint name="KVPEncoding">
        <ows:NoValues/>
        <ows:DefaultValue>TRUE</ows:DefaultValue>
    </ows:Constraint>
    <ows:Constraint name="XMLEncoding">
        <ows:NoValues/>
        <ows:DefaultValue>TRUE</ows:DefaultValue>
    </ows:Constraint>
    <ows:Constraint name="SOAPEncoding">
        <ows:NoValues/>
        <ows:DefaultValue>FALSE</ows:DefaultValue>
    </ows:Constraint>
    <ows:Constraint name="ImplementsInheritance">
        <ows:NoValues/>
        <ows:DefaultValue>FALSE</ows:DefaultValue>
    </ows:Constraint>
    <ows:Constraint name="ImplementsRemoteResolve">
        <ows:NoValues/>
        <ows:DefaultValue>FALSE</ows:DefaultValue>
    </ows:Constraint>
    <ows:Constraint name="ImplementsResultPaging">
        <ows:NoValues/>
        <ows:DefaultValue>TRUE</ows:DefaultValue>
    </ows:Constraint>
    <ows:Constraint name="ImplementsStandardJoins">
        <ows:NoValues/>
        <ows:DefaultValue>FALSE</ows:DefaultValue>
    </ows:Constraint>
    <ows:Constraint name="ImplementsSpatialJoins">
        <ows:NoValues/>
        <ows:DefaultValue>FALSE</ows:DefaultValue>
    </ows:Constraint>
    <ows:Constraint name="ImplementsTemporalJoins">
        <ows:NoValues/>
        <ows:DefaultValue>FALSE</ows:DefaultValue>
    </ows:Constraint>
    <ows:Constraint name="ImplementsFeatureVersioning">
        <ows:NoValues/>
        <ows:DefaultValue>FALSE</ows:DefaultValue>
    </ows:Constraint>
    <ows:Constraint name="ManageStoredQueries">
        <ows:NoValues/>
        <ows:DefaultValue>FALSE</ows:DefaultValue>
    </ows:Constraint>
    <ows:Constraint name="PagingIsTransactionSafe">
        <ows:NoValues/>
        <ows:DefaultValue>FALSE</ows:DefaultValue>
    </ows:Constraint>
    <ows:Constraint name="CountDefault">
        <ows:NoValues/>
        <ows:DefaultValue>1000</ows:DefaultValue>
    </ows:Constraint>
    <ows:Constraint name="QueryExpressions">
        <ows:AllowedValues>
            <ows:Value>wfs:Query</ows:Value>
            <ows:Value>wfs:StoredQuery</ows:Value>
        </ows:AllowedValues>
    </ows:Constraint>
{{if $service.Inspire}}
    <ows:ExtendedCapabilities>
        <inspire_dls:ExtendedCapabilities>
            <inspire_common:MetadataUrl xsi:type="inspire_common:resourceLocatorType">
                <inspire_common:URL>http://www.nationaalgeoregister.nl/geonetwork/srv/eng/csw?service=CSW&amp;version=2.0.2&amp;request=GetRecordById&amp;outputschema=http://www.isotc211.org/2005/gmd&amp;elementsetname=full&amp;id={{ $service.MetadataResourceId }}</inspire_common:URL>
                <inspire_common:MediaType>application/vnd.ogc.csw.GetRecordByIdResponse_xml</inspire_common:MediaType>
            </inspire_common:MetadataUrl>
            <inspire_common:SupportedLanguages>
                <inspire_common:DefaultLanguage>
                    <inspire_common:Language>eng</inspire_common:Language>
                </inspire_common:DefaultLanguage>
            </inspire_common:SupportedLanguages>
            <inspire_common:ResponseLanguage>
                <inspire_common:Language>eng</inspire_common:Language>
            </inspire_common:ResponseLanguage>
            <inspire_dls:SpatialDataSetIdentifier>
                <inspire_common:Code>{{escape $dataset.MetadataResourceId }}</inspire_common:Code>
            </inspire_dls:SpatialDataSetIdentifier>
        </inspire_dls:ExtendedCapabilities>
    </ows:ExtendedCapabilities>
{{- end}}
</ows:OperationsMetadata>