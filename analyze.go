package main

func analyze(req AnalyzeRequest) AnalyzeResponse {
	points := req.RawPoints

	var distanceSum float64
	var elevationGain float64
	var totalCalories float64

	for i := 1; i < len(points); i++ {
		distanceSum += distance(points[i-1], points[i])

		if diff := points[i].Altitude - points[i-1].Altitude; diff > 0 {
			elevationGain += diff
		}

		timeDiff := points[i].Timestamp.Sub(points[i-1].Timestamp).Hours()
		segmentDistance := distance(points[i-1], points[i]) / 1000
		speed := segmentDistance / timeDiff

		segmentCalories := (speed * 0.5) * 60 * timeDiff
		totalCalories += segmentCalories
	}

	startTime := points[0].Timestamp
	endTime := points[len(points)-1].Timestamp
	duration := endTime.Sub(startTime).Hours()
	avgSpeed := distanceSum / 1000 / duration

	points = douglasPeucker(points, 0.001)

	// 0.0005 - (50 метров) это значение, которое определяет степень упрощения. Чем меньше значение, тем больше точек останется.
	// 0.00001 - (1 метр)

	return AnalyzeResponse{
		UserID:       req.UserID,
		StartAt:      startTime,
		EndAt:        endTime,
		Distance:     distanceSum,
		Calories:     totalCalories,
		AverageSpeed: avgSpeed,
		Elevation:    elevationGain,
		Points:       points,
	}
}
