package main

import (
	"math"
)

func distance(p1, p2 Point) float64 {
	const R = 6371e3 // Радиус Земли в метрах
	lat1 := p1.Lat * math.Pi / 180
	lat2 := p2.Lat * math.Pi / 180
	dLat := (p2.Lat - p1.Lat) * math.Pi / 180
	dLng := (p2.Lng - p1.Lng) * math.Pi / 180

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1)*math.Cos(lat2)*math.Sin(dLng/2)*math.Sin(dLng/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return R * c
}

func perpendicularDistance(p, start, end Point) float64 {
	if start.Lat == end.Lat && start.Lng == end.Lng {
		return distance(p, start)
	}

	num := math.Abs((end.Lng-start.Lng)*(start.Lat-p.Lat) - (start.Lng-p.Lng)*(end.Lat-start.Lat))
	den := math.Sqrt(math.Pow(end.Lng-start.Lng, 2) + math.Pow(end.Lat-start.Lat, 2))
	return num / den
}

func douglasPeucker(points []Point, epsilon float64) []Point {
	if len(points) < 3 {
		return points
	}

	dmax := 0.0
	index := 0
	start := points[0]
	end := points[len(points)-1]

	for i := 1; i < len(points)-1; i++ {
		d := perpendicularDistance(points[i], start, end)
		if d > dmax {
			index = i
			dmax = d
		}
	}

	if dmax > epsilon {
		left := douglasPeucker(points[:index+1], epsilon)
		right := douglasPeucker(points[index:], epsilon)

		return append(left[:len(left)-1], right...)
	}

	return []Point{start, end}
}
