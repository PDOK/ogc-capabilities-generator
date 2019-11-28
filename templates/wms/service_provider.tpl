  <Service>{{ $identification:= index . "identification" }}{{ $organization:= index . "organization" }}
        <Name>WMS</Name>
        <Title>{{ escape $identification.Title}}</Title>
        <Abstract>{{ escape $identification.Abstract}}</Abstract>
        <KeywordList>{{range $i, $keyword := $identification.Keywords}}
         <Keyword>{{escape $keyword}}</Keyword>{{end}}
         <Keyword vocabulary="ISO">infoMapAccessService</Keyword>
        </KeywordList>
        <OnlineResource xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="https://geodata.nationaalgeoregister.nl"/>
        <ContactInformation>
            <ContactPersonPrimary>
                <ContactPerson>{{escape $organization.Individual}}</ContactPerson>
                <ContactOrganization>{{escape $organization.Name}}</ContactOrganization>
            </ContactPersonPrimary>
            <ContactPosition>pointOfContact</ContactPosition>
            <ContactAddress>
                <AddressType>Work</AddressType>
                <Address></Address>
                <City>{{escape $organization.City}}</City>
                <StateOrProvince></StateOrProvince>
                <PostCode></PostCode>
                <Country>{{ escape $organization.Country}}</Country>
            </ContactAddress>
            <ContactVoiceTelephone></ContactVoiceTelephone>
            <ContactFacsimileTelephone></ContactFacsimileTelephone>
            <ContactElectronicMailAddress>{{escape $organization.Email}}</ContactElectronicMailAddress>
        </ContactInformation>
        <Fees>NONE</Fees>
        <AccessConstraints>http://creativecommons.org/publicdomain/zero/1.0/deed.nl</AccessConstraints>
        <MaxWidth>10000</MaxWidth>
        <MaxHeight>10000</MaxHeight>
    </Service>
