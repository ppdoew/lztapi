package errors

var ErrorMessages = map[int]string{
	400: "Bad Request: The request is invalid due to client error.",
	401: "Unauthorized: Authentication is required.",
	403: "Forbidden: You do not have permission to view this page or perform this action.",
	429: "Too Many Requests: Too many requests in a short period.",
	500: "Internal Server Error: An internal server error occurred, try again later or report it here.",
	502: "Bad Gateway: The gateway encountered an error while processing the request, try again later.",
	503: "Service Unavailable: The server is temporarily unavailable, try again later.",
	504: "Gateway Timeout: The gateway timed out waiting for a response, try again later.",
}
