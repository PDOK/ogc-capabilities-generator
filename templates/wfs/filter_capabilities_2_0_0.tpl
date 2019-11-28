<fes:Filter_Capabilities>
<fes:Conformance>
  <fes:Constraint name="ImplementsQuery">
    <ows:NoValues/>
    <ows:DefaultValue>TRUE</ows:DefaultValue>
  </fes:Constraint>
  <fes:Constraint name="ImplementsAdHocQuery">
    <ows:NoValues/>
    <ows:DefaultValue>TRUE</ows:DefaultValue>
  </fes:Constraint>
  <fes:Constraint name="ImplementsFunctions">
    <ows:NoValues/>
    <ows:DefaultValue>FALSE</ows:DefaultValue>
  </fes:Constraint>
  <fes:Constraint name="ImplementsResourceId">
    <ows:NoValues/>
    <ows:DefaultValue>TRUE</ows:DefaultValue>
  </fes:Constraint>
  <fes:Constraint name="ImplementsMinStandardFilter">
    <ows:NoValues/>
    <ows:DefaultValue>TRUE</ows:DefaultValue>
  </fes:Constraint>
  <fes:Constraint name="ImplementsStandardFilter">
    <ows:NoValues/>
    <ows:DefaultValue>TRUE</ows:DefaultValue>
  </fes:Constraint>
  <fes:Constraint name="ImplementsMinSpatialFilter">
    <ows:NoValues/>
    <ows:DefaultValue>TRUE</ows:DefaultValue>
  </fes:Constraint>
  <fes:Constraint name="ImplementsSpatialFilter">
    <ows:NoValues/>
    <ows:DefaultValue>FALSE</ows:DefaultValue>
  </fes:Constraint>
  <fes:Constraint name="ImplementsMinTemporalFilter">
    <ows:NoValues/>
    <ows:DefaultValue>TRUE</ows:DefaultValue>
  </fes:Constraint>
  <fes:Constraint name="ImplementsTemporalFilter">
    <ows:NoValues/>
    <ows:DefaultValue>FALSE</ows:DefaultValue>
  </fes:Constraint>
  <fes:Constraint name="ImplementsVersionNav">
    <ows:NoValues/>
    <ows:DefaultValue>FALSE</ows:DefaultValue>
  </fes:Constraint>
  <fes:Constraint name="ImplementsSorting">
    <ows:NoValues/>
    <ows:DefaultValue>TRUE</ows:DefaultValue>
  </fes:Constraint>
  <fes:Constraint name="ImplementsExtendedOperators">
    <ows:NoValues/>
    <ows:DefaultValue>FALSE</ows:DefaultValue>
  </fes:Constraint>
  <fes:Constraint name="ImplementsMinimumXPath">
    <ows:NoValues/>
    <ows:DefaultValue>TRUE</ows:DefaultValue>
  </fes:Constraint>
  <fes:Constraint name="ImplementsSchemaElementFunc">
    <ows:NoValues/>
    <ows:DefaultValue>FALSE</ows:DefaultValue>
  </fes:Constraint>
</fes:Conformance>
<fes:Id_Capabilities>
  <fes:ResourceIdentifier name="fes:ResourceId"/>
</fes:Id_Capabilities>
<fes:Scalar_Capabilities>
  <fes:LogicalOperators/>
  <fes:ComparisonOperators>
    <fes:ComparisonOperator name="PropertyIsEqualTo"/>
    <fes:ComparisonOperator name="PropertyIsNotEqualTo"/>
    <fes:ComparisonOperator name="PropertyIsLessThan"/>
    <fes:ComparisonOperator name="PropertyIsGreaterThan"/>
    <fes:ComparisonOperator name="PropertyIsLessThanOrEqualTo"/>
    <fes:ComparisonOperator name="PropertyIsGreaterThanOrEqualTo"/>
    <fes:ComparisonOperator name="PropertyIsLike"/>
    <fes:ComparisonOperator name="PropertyIsBetween"/>
  </fes:ComparisonOperators>
</fes:Scalar_Capabilities>
<fes:Spatial_Capabilities>
  <fes:GeometryOperands>
    <fes:GeometryOperand name="gml:Point"/>
    <fes:GeometryOperand name="gml:MultiPoint"/>
    <fes:GeometryOperand name="gml:LineString"/>
    <fes:GeometryOperand name="gml:MultiLineString"/>
    <fes:GeometryOperand name="gml:Curve"/>
    <fes:GeometryOperand name="gml:MultiCurve"/>
    <fes:GeometryOperand name="gml:Polygon"/>
    <fes:GeometryOperand name="gml:MultiPolygon"/>
    <fes:GeometryOperand name="gml:Surface"/>
    <fes:GeometryOperand name="gml:MultiSurface"/>
    <fes:GeometryOperand name="gml:Box"/>
    <fes:GeometryOperand name="gml:Envelope"/>
  </fes:GeometryOperands>
  <fes:SpatialOperators>
    <fes:SpatialOperator name="Equals"/>
    <fes:SpatialOperator name="Disjoint"/>
    <fes:SpatialOperator name="Touches"/>
    <fes:SpatialOperator name="Within"/>
    <fes:SpatialOperator name="Overlaps"/>
    <fes:SpatialOperator name="Crosses"/>
    <fes:SpatialOperator name="Intersects"/>
    <fes:SpatialOperator name="Contains"/>
    <fes:SpatialOperator name="DWithin"/>
    <fes:SpatialOperator name="Beyond"/>
    <fes:SpatialOperator name="BBOX"/>
  </fes:SpatialOperators>
</fes:Spatial_Capabilities>
<fes:Temporal_Capabilities>
  <fes:TemporalOperands>
    <fes:TemporalOperand name="gml:TimePeriod"/>
    <fes:TemporalOperand name="gml:TimeInstant"/>
  </fes:TemporalOperands>
  <fes:TemporalOperators>
    <fes:TemporalOperator name="During"/>
  </fes:TemporalOperators>
</fes:Temporal_Capabilities>
</fes:Filter_Capabilities>