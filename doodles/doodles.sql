SELECT ST_AsText(ST_POINT(1,2));

SELECT ST_AsText( ST_MakeLine(ST_Point(1,2), ST_Point(3,4)) );

SELECT ST_AsEWKT( ST_MakeLine(ST_MakePoint(1,2,3), ST_MakePoint(3,4,5) ));

SELECT ST_AsText( ST_MakeLine(ST_MakePoint(1,2,3), ST_MakePoint(3,4,5) ));

select ST_AsText( ST_MakeLine( 'LINESTRING(0 0, 1 1)', 'LINESTRING(2 2, 3 3)' ) );

CREATE TABLE korean_restaurants (
    id SERIAL PRIMARY KEY,
    name VARCHAR(64),
    location geography(POINT,4326)
)

-- https://postgis.net/docs/manual-3.3/using_postgis_dbmanagement.html#spatial_ref_sys
-- https://postgis.net/workshops/postgis-intro/knn.html

-- tasty kalby restauramt
SELECT 
		kr.name,
		kr.description, 
		kr.the_geog <-> 'SRID=4326;POINT(126.9711902547677 37.22769003295954)'::geography AS dist
	FROM
		korean_restaurants kr
		ORDER BY
		dist
		LIMIT 3;

-- jajjang rest
-- tasty kalby restauramt
SELECT 
		kr.name,
		kr.description, 
		kr.the_geog <-> 'SRID=4326;POINT(126.9745752913237 37.22214152223785)'::geography AS dist,
        ST_AsText(kr.the_geog)
	FROM
		korean_restaurants kr
	ORDER BY
		dist
	LIMIT 3;

--ST_AsEKWT
--SRID=4326;POINT(126.9747364367602 37.22265010668237 43.97082937402935)

--ST_AsText
--POINT Z (126.9775201550173 37.22450239990378 35.55338264945428)


SELECT 
		kr.name,
		kr.description, 
		kr.the_geog <-> 'SRID=4326;POINT(126.9745752913237 37.22214152223785)'::geography AS dist,
        ST_AsText(kr.the_geog)
	FROM
		korean_restaurants kr
    WHERE kr.the_geog <-> 'SRID=4326;POINT(126.9745752913237 37.22214152223785)'::geography < 1000
    ORDER BY
		dist
	LIMIT 3;