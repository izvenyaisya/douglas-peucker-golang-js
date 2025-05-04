package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func analyzeHandler(w http.ResponseWriter, r *http.Request) {
	var req AnalyzeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println("Received request", req)
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	log.Println("Received request", req)

	resp := analyze(req)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func mustParseTime(value string) time.Time {
	t, err := time.Parse(time.RFC3339, value)
	if err != nil {
		log.Fatalf("failed to parse time: %v", err)
	}
	return t
}

func main() {

	rawPoints := []Point{
		{Lat: 45.11123944267117, Lng: 41.42232545796264, Altitude: 200, Timestamp: mustParseTime("2023-01-01T10:02:20Z")},
		{Lat: 45.11097688763551, Lng: 41.79235402245122, Altitude: 218, Timestamp: mustParseTime("2023-01-01T10:00:30Z")},
		{Lat: 45.136455694371335, Lng: 41.22249020215779, Altitude: 186, Timestamp: mustParseTime("2023-01-01T10:00:15Z")},
		{Lat: 45.14095212976634, Lng: 41.32733010116003, Altitude: 220, Timestamp: mustParseTime("2023-01-01T10:03:30Z")},
		{Lat: 45.14128978484598, Lng: 41.7755221607215, Altitude: 210, Timestamp: mustParseTime("2023-01-01T10:02:55Z")},
		{Lat: 45.14929815721705, Lng: 41.956176529476156, Altitude: 220, Timestamp: mustParseTime("2023-01-01T10:00:40Z")},
		{Lat: 45.14975564701951, Lng: 41.956234050586236, Altitude: 216, Timestamp: mustParseTime("2023-01-01T10:01:45Z")},
		{Lat: 45.15064898304772, Lng: 41.362515872697694, Altitude: 200, Timestamp: mustParseTime("2023-01-01T10:01:25Z")},
		{Lat: 45.171493278847144, Lng: 41.287116118767905, Altitude: 221, Timestamp: mustParseTime("2023-01-01T10:00:05Z")},
		{Lat: 45.17592127656152, Lng: 41.876798325091895, Altitude: 207, Timestamp: mustParseTime("2023-01-01T10:03:45Z")},
		{Lat: 45.22909041053532, Lng: 41.29054622157357, Altitude: 210, Timestamp: mustParseTime("2023-01-01T10:01:50Z")},
		{Lat: 45.23534580483821, Lng: 41.86838346947763, Altitude: 205, Timestamp: mustParseTime("2023-01-01T10:04:20Z")},
		{Lat: 45.253273133295785, Lng: 41.43176717147955, Altitude: 202, Timestamp: mustParseTime("2023-01-01T10:01:15Z")},
		{Lat: 45.27184894560271, Lng: 41.91746030645298, Altitude: 195, Timestamp: mustParseTime("2023-01-01T10:01:40Z")},
		{Lat: 45.30271287528029, Lng: 41.44579444225853, Altitude: 220, Timestamp: mustParseTime("2023-01-01T10:04:00Z")},
		{Lat: 45.312024117673625, Lng: 41.91476521468772, Altitude: 217, Timestamp: mustParseTime("2023-01-01T10:02:10Z")},
		{Lat: 45.31254213720763, Lng: 41.53344568764551, Altitude: 215, Timestamp: mustParseTime("2023-01-01T10:02:00Z")},
		{Lat: 45.35046349488744, Lng: 41.33495339099349, Altitude: 220, Timestamp: mustParseTime("2023-01-01T10:01:00Z")},
		{Lat: 45.356343079126856, Lng: 41.9843474944302, Altitude: 207, Timestamp: mustParseTime("2023-01-01T10:04:10Z")},
		{Lat: 45.364315368722134, Lng: 41.84347631915938, Altitude: 218, Timestamp: mustParseTime("2023-01-01T10:01:35Z")},
		{Lat: 45.36592449053349, Lng: 41.91748443455901, Altitude: 185, Timestamp: mustParseTime("2023-01-01T10:01:10Z")},
		{Lat: 45.369276968942015, Lng: 41.40178875545406, Altitude: 214, Timestamp: mustParseTime("2023-01-01T10:03:25Z")},
		{Lat: 45.37203596427965, Lng: 41.57332344612122, Altitude: 215, Timestamp: mustParseTime("2023-01-01T10:02:50Z")},
		{Lat: 45.37411764952784, Lng: 41.35445224398352, Altitude: 212, Timestamp: mustParseTime("2023-01-01T10:00:45Z")},
		{Lat: 45.37679238510492, Lng: 41.58240109004639, Altitude: 220, Timestamp: mustParseTime("2023-01-01T10:03:15Z")},
		{Lat: 45.37767072341963, Lng: 41.745800231788924, Altitude: 209, Timestamp: mustParseTime("2023-01-01T10:02:45Z")},
		{Lat: 45.386672302239235, Lng: 41.87422698268424, Altitude: 210, Timestamp: mustParseTime("2023-01-01T10:00:25Z")},
		{Lat: 45.38773011624785, Lng: 41.94240569848764, Altitude: 210, Timestamp: mustParseTime("2023-01-01T10:01:30Z")},
		{Lat: 45.38867902895332, Lng: 41.2157950490933, Altitude: 212, Timestamp: mustParseTime("2023-01-01T10:02:35Z")},
		{Lat: 45.398241794433595, Lng: 41.461051431717064, Altitude: 220, Timestamp: mustParseTime("2023-01-01T10:04:15Z")},
		{Lat: 45.402902567008205, Lng: 41.51583132459732, Altitude: 213, Timestamp: mustParseTime("2023-01-01T10:02:25Z")},
		{Lat: 45.44154730803561, Lng: 41.87001136957903, Altitude: 209, Timestamp: mustParseTime("2023-01-01T10:03:55Z")},
		{Lat: 45.47211249836677, Lng: 41.29590538952419, Altitude: 212, Timestamp: mustParseTime("2023-01-01T10:03:35Z")},
		{Lat: 45.487102800222226, Lng: 41.4050135723177, Altitude: 205, Timestamp: mustParseTime("2023-01-01T10:03:10Z")},
		{Lat: 45.490571772139454, Lng: 41.94012686609984, Altitude: 204, Timestamp: mustParseTime("2023-01-01T10:00:10Z")},
		{Lat: 45.55091297812567, Lng: 41.28052958801349, Altitude: 210, Timestamp: mustParseTime("2023-01-01T10:02:15Z")},
		{Lat: 45.56046239864609, Lng: 41.73119323555379, Altitude: 217, Timestamp: mustParseTime("2023-01-01T10:01:05Z")},
		{Lat: 45.59528583601386, Lng: 41.74084923246229, Altitude: 220, Timestamp: mustParseTime("2023-01-01T10:01:20Z")},
		{Lat: 45.60985450715851, Lng: 41.89556210385765, Altitude: 220, Timestamp: mustParseTime("2023-01-01T10:02:05Z")},
		{Lat: 45.626328799236515, Lng: 41.94834649570223, Altitude: 218, Timestamp: mustParseTime("2023-01-01T10:03:40Z")},
		{Lat: 45.63159212171534, Lng: 41.31751533763576, Altitude: 220, Timestamp: mustParseTime("2023-01-01T10:03:00Z")},
		{Lat: 45.63421876704799, Lng: 41.62372652857701, Altitude: 215, Timestamp: mustParseTime("2023-01-01T10:00:50Z")},
		{Lat: 45.66127467433412, Lng: 41.72567984307382, Altitude: 195, Timestamp: mustParseTime("2023-01-01T10:03:05Z")},
		{Lat: 45.678936494981324, Lng: 41.89330885209151, Altitude: 220, Timestamp: mustParseTime("2023-01-01T10:02:30Z")},
		{Lat: 45.6807811177318, Lng: 41.52554794586883, Altitude: 205, Timestamp: mustParseTime("2023-01-01T10:01:55Z")},
		{Lat: 45.70453709392001, Lng: 41.37006473107736, Altitude: 198, Timestamp: mustParseTime("2023-01-01T10:03:50Z")},
		{Lat: 45.70484390185055, Lng: 41.394464851616476, Altitude: 190, Timestamp: mustParseTime("2023-01-01T10:00:20Z")},
		{Lat: 45.71188103775323, Lng: 41.36537452230856, Altitude: 200, Timestamp: mustParseTime("2023-01-01T10:00:35Z")},
		{Lat: 45.71567996574397, Lng: 41.28175912982065, Altitude: 209, Timestamp: mustParseTime("2023-01-01T10:00:55Z")},
		{Lat: 45.72857732503699, Lng: 41.64179751073826, Altitude: 210, Timestamp: mustParseTime("2023-01-01T10:03:20Z")},
		{Lat: 45.74166976241723, Lng: 41.32593945208132, Altitude: 182, Timestamp: mustParseTime("2023-01-01T10:00:00Z")},
		{Lat: 45.74668798262231, Lng: 41.99288470386944, Altitude: 210, Timestamp: mustParseTime("2023-01-01T10:02:40Z")},
	}

	startTime := time.Now()

	inputData := AnalyzeRequest{
		UserID:    1,
		RawPoints: rawPoints,
	}

	log.Println("Starting analysis...")

	resp := analyze(inputData)

	elapsedTime := time.Since(startTime)

	log.Println("analyzeExecutionTime: start")
	output := map[string]interface{}{
		"userId":       resp.UserID,
		"startAt":      resp.StartAt.Format(time.RFC3339),
		"endAt":        resp.EndAt.Format(time.RFC3339),
		"distance":     resp.Distance,
		"calories":     resp.Calories,
		"averageSpeed": resp.AverageSpeed,
		"elevation":    resp.Elevation,
		"points":       resp.Points,
		"elapsedTime":  elapsedTime,
	}

	log.Println("analyzeExecutionTime:", elapsedTime)
	log.Println("analyzeExecutionTime: end")

	jsonOutput, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		log.Fatalf("failed to marshal output: %v", err)
	}

	log.Println("Analysis result:")
	log.Println(string(jsonOutput))

	staticMapURL := "https://maps.googleapis.com/maps/api/staticmap?size=600x400&path=color:0xff0000ff|weight:5|"
	for i, point := range resp.Points {
		if i > 0 {
			staticMapURL += "|"
		}
		staticMapURL += fmt.Sprintf("%f,%f", point.Lat, point.Lng)
	}
	staticMapURL += "&key=YOUR_API_KEY"

	log.Println("Static Map URL:", staticMapURL)
}

/*
func main() {
	port := flag.String("port", "3010", "port to listen on")
	flag.Parse()

	http.HandleFunc("/analyze", analyzeHandler)
	log.Printf("Listening on localhost:%s...", *port)
	log.Fatal(http.ListenAndServe("localhost:"+*port, nil))
}
*/
