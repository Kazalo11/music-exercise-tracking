
export function getBackendHost() {
	return process.env.ENV === "prod" ? "https://backend-1052978901140.europe-west2.run.app": "http://localhost"
}