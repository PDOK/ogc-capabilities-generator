 <ows:ServiceProvider xmlns="http://www.opengis.net/wfs/2.0" xmlns:ows="http://www.opengis.net/ows/1.1">
   <ows:ProviderName>{{ .Name}}</ows:ProviderName>
    <ows:ServiceContact>
        <ows:IndividualName>{{ .Individual}}</ows:IndividualName>
        <ows:PositionName>{{ .Position}}</ows:PositionName>
        <ows:ContactInfo>
            <ows:Phone>
                <ows:Voice/>
                <ows:Facsimile/>
            </ows:Phone>
            <ows:Address>
                <ows:DeliveryPoint/>
                <ows:City>{{ .City}}</ows:City>
                <ows:AdministrativeArea/>
                <ows:PostalCode/>
                <ows:Country>{{ .Country}}</ows:Country>
                <ows:ElectronicMailAddress>{{ .Email}}</ows:ElectronicMailAddress>
            </ows:Address>
        </ows:ContactInfo>
         <ows:Role></ows:Role>
    </ows:ServiceContact>
 </ows:ServiceProvider>