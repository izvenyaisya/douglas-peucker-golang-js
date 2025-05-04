interface Point {
  lat: number;
  lng: number;
  altitude: number;
  timestamp: string;
}

interface InputData {
  userId: number;
  rawPoints: Point[];
}

interface AnalysisResult {
  userId: number;
  startAt: Date;
  endAt: Date;
  distance: number;
  calories: number;
  averageSpeed: number;
  elevation: number;
  points: Point[];
}

function distance(p1: Point, p2: Point): number {
  const R = 6371e3; // Radius of Earth in meters
  const lat1 = (p1.lat * Math.PI) / 180;
  const lat2 = (p2.lat * Math.PI) / 180;
  const dLat = ((p2.lat - p1.lat) * Math.PI) / 180;
  const dLng = ((p2.lng - p1.lng) * Math.PI) / 180;

  const a =
    Math.sin(dLat / 2) * Math.sin(dLat / 2) +
    Math.cos(lat1) * Math.cos(lat2) * Math.sin(dLng / 2) * Math.sin(dLng / 2);
  const c = 2 * Math.atan2(Math.sqrt(a), Math.sqrt(1 - a));

  return R * c;
}

function perpendicularDistance(p: Point, start: Point, end: Point): number {
  if (start.lat === end.lat && start.lng === end.lng) {
    return distance(p, start);
  }

  const num = Math.abs(
    (end.lng - start.lng) * (start.lat - p.lat) -
      (start.lng - p.lng) * (end.lat - start.lat)
  );
  const den = Math.sqrt(
    Math.pow(end.lng - start.lng, 2) + Math.pow(end.lat - start.lat, 2)
  );
  return num / den;
}

function douglasPeucker(points: Point[], epsilon: number): Point[] {
  if (points.length < 3) {
    return points;
  }

  let dmax = 0;
  let index = 0;
  const start = points[0];
  const end = points[points.length - 1];

  for (let i = 1; i < points.length - 1; i++) {
    const d = perpendicularDistance(points[i], start, end);
    if (d > dmax) {
      index = i;
      dmax = d;
    }
  }

  if (dmax > epsilon) {
    const left = douglasPeucker(points.slice(0, index + 1), epsilon);
    const right = douglasPeucker(points.slice(index), epsilon);

    return [...left.slice(0, -1), ...right];
  }

  return [start, end];
}

function analyze(data: InputData): AnalysisResult {
  const points = data.rawPoints;

  let distanceSum = 0;
  let elevationGain = 0;
  let totalCalories = 0;

  for (let i = 1; i < points.length; i++) {
    distanceSum += distance(points[i - 1], points[i]);

    const diff = points[i].altitude - points[i - 1].altitude;
    if (diff > 0) {
      elevationGain += diff;
    }

    const timeDiff =
      (new Date(points[i].timestamp).getTime() -
        new Date(points[i - 1].timestamp).getTime()) /
      (1000 * 60 * 60);
    const segmentDistance = distance(points[i - 1], points[i]) / 1000;
    const speed = segmentDistance / timeDiff;

    const segmentCalories = speed * 0.5 * 60 * timeDiff;
    totalCalories += segmentCalories;
  }

  const startTime = new Date(points[0].timestamp);
  const endTime = new Date(points[points.length - 1].timestamp);
  const duration = (endTime.getTime() - startTime.getTime()) / (1000 * 60 * 60);
  const avgSpeed = distanceSum / 1000 / duration;

  const simplifiedPoints = douglasPeucker(points, 0.001);

  return {
    userId: data.userId,
    startAt: startTime,
    endAt: endTime,
    distance: distanceSum,
    calories: totalCalories,
    averageSpeed: avgSpeed,
    elevation: elevationGain,
    points: simplifiedPoints,
  };
}

// Example data
const inputData: InputData = {
  userId: 1,
  rawPoints: [
    { lat: 45.11123944267117, lng: 41.42232545796264, altitude: 200, timestamp: "2023-01-01T10:02:20Z" },
    { lat: 45.11097688763551, lng: 41.79235402245122, altitude: 218, timestamp: "2023-01-01T10:00:30Z" },
    { lat: 45.136455694371335, lng: 41.22249020215779, altitude: 186, timestamp: "2023-01-01T10:00:15Z" },
    { lat: 45.14095212976634, lng: 41.32733010116003, altitude: 220, timestamp: "2023-01-01T10:03:30Z" },
    { lat: 45.14128978484598, lng: 41.7755221607215, altitude: 210, timestamp: "2023-01-01T10:02:55Z" },
    { lat: 45.14929815721705, lng: 41.956176529476156, altitude: 220, timestamp: "2023-01-01T10:00:40Z" },
    { lat: 45.14975564701951, lng: 41.956234050586236, altitude: 216, timestamp: "2023-01-01T10:01:45Z" },
    { lat: 45.15064898304772, lng: 41.362515872697694, altitude: 200, timestamp: "2023-01-01T10:01:25Z" },
    { lat: 45.171493278847144, lng: 41.287116118767905, altitude: 221, timestamp: "2023-01-01T10:00:05Z" },
    { lat: 45.17592127656152, lng: 41.876798325091895, altitude: 207, timestamp: "2023-01-01T10:03:45Z" },
    { lat: 45.22909041053532, lng: 41.29054622157357, altitude: 210, timestamp: "2023-01-01T10:01:50Z" },
    { lat: 45.23534580483821, lng: 41.86838346947763, altitude: 205, timestamp: "2023-01-01T10:04:20Z" },
    { lat: 45.253273133295785, lng: 41.43176717147955, altitude: 202, timestamp: "2023-01-01T10:01:15Z" },
    { lat: 45.27184894560271, lng: 41.91746030645298, altitude: 195, timestamp: "2023-01-01T10:01:40Z" },
    { lat: 45.30271287528029, lng: 41.44579444225853, altitude: 220, timestamp: "2023-01-01T10:04:00Z" },
    { lat: 45.312024117673625, lng: 41.91476521468772, altitude: 217, timestamp: "2023-01-01T10:02:10Z" },
    { lat: 45.31254213720763, lng: 41.53344568764551, altitude: 215, timestamp: "2023-01-01T10:02:00Z" },
    { lat: 45.35046349488744, lng: 41.33495339099349, altitude: 220, timestamp: "2023-01-01T10:01:00Z" },
    { lat: 45.356343079126856, lng: 41.9843474944302, altitude: 207, timestamp: "2023-01-01T10:04:10Z" },
    { lat: 45.364315368722134, lng: 41.84347631915938, altitude: 218, timestamp: "2023-01-01T10:01:35Z" },
    { lat: 45.36592449053349, lng: 41.91748443455901, altitude: 185, timestamp: "2023-01-01T10:01:10Z" },
    { lat: 45.369276968942015, lng: 41.40178875545406, altitude: 214, timestamp: "2023-01-01T10:03:25Z" },
    { lat: 45.37203596427965, lng: 41.57332344612122, altitude: 215, timestamp: "2023-01-01T10:02:50Z" },
    { lat: 45.37411764952784, lng: 41.35445224398352, altitude: 212, timestamp: "2023-01-01T10:00:45Z" },
    { lat: 45.37679238510492, lng: 41.58240109004639, altitude: 220, timestamp: "2023-01-01T10:03:15Z" },
    { lat: 45.37767072341963, lng: 41.745800231788924, altitude: 209, timestamp: "2023-01-01T10:02:45Z" },
    { lat: 45.386672302239235, lng: 41.87422698268424, altitude: 210, timestamp: "2023-01-01T10:00:25Z" },
    { lat: 45.38773011624785, lng: 41.94240569848764, altitude: 210, timestamp: "2023-01-01T10:01:30Z" },
    { lat: 45.38867902895332, lng: 41.2157950490933, altitude: 212, timestamp: "2023-01-01T10:02:35Z" },
    { lat: 45.398241794433595, lng: 41.461051431717064, altitude: 220, timestamp: "2023-01-01T10:04:15Z" },
    { lat: 45.402902567008205, lng: 41.51583132459732, altitude: 213, timestamp: "2023-01-01T10:02:25Z" },
    { lat: 45.44154730803561, lng: 41.87001136957903, altitude: 209, timestamp: "2023-01-01T10:03:55Z" },
    { lat: 45.47211249836677, lng: 41.29590538952419, altitude: 212, timestamp: "2023-01-01T10:03:35Z" },
    { lat: 45.487102800222226, lng: 41.4050135723177, altitude: 205, timestamp: "2023-01-01T10:03:10Z" },
    { lat: 45.490571772139454, lng: 41.94012686609984, altitude: 204, timestamp: "2023-01-01T10:00:10Z" },
    { lat: 45.55091297812567, lng: 41.28052958801349, altitude: 210, timestamp: "2023-01-01T10:02:15Z" },
    { lat: 45.56046239864609, lng: 41.73119323555379, altitude: 217, timestamp: "2023-01-01T10:01:05Z" },
    { lat: 45.59528583601386, lng: 41.74084923246229, altitude: 220, timestamp: "2023-01-01T10:01:20Z" },
    { lat: 45.60985450715851, lng: 41.89556210385765, altitude: 220, timestamp: "2023-01-01T10:02:05Z" },
    { lat: 45.626328799236515, lng: 41.94834649570223, altitude: 218, timestamp: "2023-01-01T10:03:40Z" },
    { lat: 45.63159212171534, lng: 41.31751533763576, altitude: 220, timestamp: "2023-01-01T10:03:00Z" },
    { lat: 45.63421876704799, lng: 41.62372652857701, altitude: 215, timestamp: "2023-01-01T10:00:50Z" },
    { lat: 45.66127467433412, lng: 41.72567984307382, altitude: 195, timestamp: "2023-01-01T10:03:05Z" },
    { lat: 45.678936494981324, lng: 41.89330885209151, altitude: 220, timestamp: "2023-01-01T10:02:30Z" },
    { lat: 45.6807811177318, lng: 41.52554794586883, altitude: 205, timestamp: "2023-01-01T10:01:55Z" },
    { lat: 45.70453709392001, lng: 41.37006473107736, altitude: 198, timestamp: "2023-01-01T10:03:50Z" },
    { lat: 45.70484390185055, lng: 41.394464851616476, altitude: 190, timestamp: "2023-01-01T10:00:20Z" },
    { lat: 45.71188103775323, lng: 41.36537452230856, altitude: 200, timestamp: "2023-01-01T10:00:35Z" },
    { lat: 45.71567996574397, lng: 41.28175912982065, altitude: 209, timestamp: "2023-01-01T10:00:55Z" },
    { lat: 45.72857732503699, lng: 41.64179751073826, altitude: 210, timestamp: "2023-01-01T10:03:20Z" },
    { lat: 45.74166976241723, lng: 41.32593945208132, altitude: 182, timestamp: "2023-01-01T10:00:00Z" },
    { lat: 45.74668798262231, lng: 41.99288470386944, altitude: 210, timestamp: "2023-01-01T10:02:40Z" },
  ],
};

console.time("analyzeExecutionTime");
const result = analyze(inputData);
console.timeEnd("analyzeExecutionTime");
const staticMapUrl = `https://maps.googleapis.com/maps/api/staticmap?size=600x400&path=color:0xff0000ff|weight:5|${result.points
  .map((p) => `${p.lat},${p.lng}`)
  .join("|")}&key=YOUR_API_KEY`;

console.log(result);
console.log("Static Map URL:", staticMapUrl);
