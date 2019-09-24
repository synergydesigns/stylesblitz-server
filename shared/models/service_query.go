package models

import "fmt"

func searchServiceQuery(lat *float64, lng *float64, name string, rating *SortRating, price *SortPrice) string {
	var priceQuery string
	var selects string
	var joins string
	var assignments string

	queryLocation := func() bool {
		if lng != nil && lat != nil {
			return true
		} else {
			return false
		}
	}()

	if price != nil {
		if *price == SortPriceHighest {
			priceQuery = "DESC"
		}

		if *price == SortPriceLowest {
			priceQuery = "ASC"
		}
	}

	if queryLocation {
		selects = ", vendor_id, geog, city"
		joins = `
			LEFT JOIN vendor_address ON vendor_address.vendor_id = services.vendor_id
			LEFT JOIN address ON address.id = vendor_address.address_id
			GROUP BY services.id, geog, city
		`
		assignments = `
			,address.geog as geog,
			address.city as city
		`
	}

	query := fmt.Sprintf(`SELECT id, name, duration, duration_type, price, trending, vendor_id, tvs %s
		FROM (
				SELECT
					services.id as id,
					services.name as name,
					services.vendor_id as vendor_id,
					services.duration as duration,
					services.duration_type as duration_type,
					services.price as price,
					services.trending as trending,
					services.tsv as tsv,
					%s
				FROM services
				%s
				ORDER BY services.price %s
		) search
		WHERE search.tsv @@ plainto_tsquery('%s')
	`, selects, assignments, joins, priceQuery, name)

	if queryLocation {
		query = fmt.Sprintf(`%s AND ST_DWithin(search.geog, ST_Point(%f, %f)::geography, 2000);`, query, *lng, *lat)
	}

	return query
}
