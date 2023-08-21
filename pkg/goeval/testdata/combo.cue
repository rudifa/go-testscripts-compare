// ch.swisstopo.images-swissimage-dop10.metadata
// @source https://api3.geo.admin.ch/rest/services/all/MapServer/identify?geometry=678250,213000&geometryFormat=geojson&geometryType=esriGeometryPoint&imageDisplay=1391,1070,96&lang=fr&layers=all:ch.swisstopo.images-swissimage-dop10.metadata&mapExtent=312250,-77500,1007750,457500&returnGeometry=true&tolerance=5

package tiles

#CoordBbox: [ number, number, number, number]

#TileId: =~"^\\d{4}_\\d{4}$" // "2675_1211"

#AsOf: =~"^\\d{4}$" // "2021"

#Resolution: =~"^\\d{1,2}$" // "10"

#FeatureAttributes: {
	datenstand: #AsOf
	resolution: #Resolution
	label:	#TileId
}

#Feature: {
	featureId:  #TileId // "2675_1211"
	bbox:	  #CoordBbox // [675000.0, 211000.0, 676000.0, 212000.0]
	layerBodId: string   // "ch.swisstopo.images-swissimage-dop10.metadata"
	layerName:  string   // "D\u00e9coupage SWISSIMAGE 10 cm Raster"
	id:	  featureId  // "2675_1211"
	attributes: #FeatureAttributes
}

#FeatureCollection: [...#Feature]

results: #FeatureCollection

#toApp: {
  tilesConfig: {
	let firstResult = results[0] // any other result would be fine also
	layerName: firstResult.layerName
	layerBodId: firstResult.layerBodId
	for f in results {
	  "\(f.featureId)": {
		bbox:	  f.bbox
		datenstand: f.attributes.datenstand
		resolution: f.attributes.resolution
	  }
	}
  }
}
