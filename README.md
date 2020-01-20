# telemetry-developer-challenge
_Context: If you're seeing this, it means we'd like to see how you would solve a real-world problem. But to be totally clear, we won't be using your code for anything beyond deciding on moving forward in the hiring process._

### Link click tracking excercise
In this exercise you will be building a fully functioning Go application packaged into a Docker Compose environment.

The Go app you will build is a link tracking service. The app will provide an API for turning a destination URL into a tracking URL. When the tracking URL is requested, the application should track the request by incrementing a counter before it redirects the user to original destination URL.

### Spec
The app should handle three endpoints:
 - `POST /url` - Receives a destination URL and returns a tracking URL
 - `GET /stats` - Returns a JSON report of URL request statistics
 - `GET /track/{id}`- The base of the tracking URL
 
 1. The `/url` endpoint should accept a JSON request containing the destination url and respond with a tracking URL
 2. Record the number of requests to each tracked URL. Data should be stored in Redis or Postgres
 3. The `stats` endpoint should return a JSON response that provides request counts for each tracking URL ordered from most requested to least requested.
 4. Package files into a fully functional `docker-compose` project. To test, we should be able to run `docker-compose up` and point our local browser to `http://localhost:5000/stats` to get the intial JSON report.
 
 ### Grading & Submission
 
Your solution will be graded on functionality, completeness, cleanliness, production readiness, and code structure. If you are unsure of a requirement just make a reasonable assumption and add a comment in your code.

Create a .zip of your project and include a README containing an overview of your solution with a self-critique explaining what you would continue iterating on if you had more time.
