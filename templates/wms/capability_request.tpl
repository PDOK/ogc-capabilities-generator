   <Request {{ $service:= index . "service" }}>
        <GetCapabilities>
            <Format>text/xml</Format>
            <DCPType>
                <HTTP>
                    <Get>
                        <OnlineResource xmlns:xlink="http://www.w3.org/1999/xlink"
                                        xlink:href="{{escape $service.UrlWms }}?SERVICE=WMS&amp;language=dut&amp;"/>
                    </Get>
                    <Post>
                        <OnlineResource xmlns:xlink="http://www.w3.org/1999/xlink"
                                        xlink:href="{{escape $service.UrlWms }}?SERVICE=WMS&amp;language=dut&amp;"/>
                    </Post>
                </HTTP>
            </DCPType>
        </GetCapabilities>
        <GetMap>
            <Format>image/png</Format>
            <Format>image/jpeg</Format>
            <Format>image/png; mode=8bit</Format>
            <Format>image/vnd.jpeg-png</Format>
            <Format>image/vnd.jpeg-png8</Format>
            <DCPType>
                <HTTP>
                    <Get>
                        <OnlineResource xmlns:xlink="http://www.w3.org/1999/xlink"
                                        xlink:href="{{escape $service.UrlWms }}SERVICE=WMS&amp;language=dut&amp;"/>
                    </Get>
                    <Post>
                        <OnlineResource xmlns:xlink="http://www.w3.org/1999/xlink"
                                        xlink:href="{{escape $service.UrlWms }}&amp;language=dut&amp;"/>
                    </Post>
                </HTTP>
            </DCPType>
        </GetMap>
        <GetFeatureInfo>
            <Format>application/json</Format>
            <Format>application/json; subtype=geojson</Format>
            <Format>application/vnd.ogc.gml</Format>
            <Format>text/html</Format>
            <Format>text/plain</Format>
            <Format>text/xml</Format>
            <Format>text/xml; subtype=gml/3.1.1</Format>
            <DCPType>
                <HTTP>
                    <Get>
                        <OnlineResource xmlns:xlink="http://www.w3.org/1999/xlink"
                                        xlink:href="{{escape $service.UrlWms }}?SERVICE=WMS&amp;language=dut&amp;"/>
                    </Get>
                    <Post>
                        <OnlineResource xmlns:xlink="http://www.w3.org/1999/xlink"
                                        xlink:href="{{escape $service.UrlWms }}?SERVICE=WMS&amp;language=dut&amp;"/>
                    </Post>
                </HTTP>
            </DCPType>
        </GetFeatureInfo>
    </Request>