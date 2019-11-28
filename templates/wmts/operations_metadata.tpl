{{ $service:= index . "service" }}    <ows:OperationsMetadata>
        <ows:Operation name="GetCapabilities">
            <ows:DCP>
                <ows:HTTP>
                    <ows:Get xlink:href="{{ $service.UrlWmts }}">
                        <ows:Constraint name="GetEncoding">
                            <ows:AllowedValues>
                                <ows:Value>KVP</ows:Value>
                            </ows:AllowedValues>
                        </ows:Constraint>
                    </ows:Get>
                </ows:HTTP>
            </ows:DCP>
        </ows:Operation>
        <ows:Operation name="GetTile">
            <ows:DCP>
                <ows:HTTP>
                    <ows:Get xlink:href="{{ $service.UrlWmts }}">
                        <ows:Constraint name="GetEncoding">
                            <ows:AllowedValues>
                                <ows:Value>KVP</ows:Value>
                            </ows:AllowedValues>
                        </ows:Constraint>
                    </ows:Get>
                </ows:HTTP>
            </ows:DCP>
        </ows:Operation>
    </ows:OperationsMetadata>