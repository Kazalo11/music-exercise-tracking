
export function getBackendHost() {
	return import.meta.env.PROD ? "https://backend-1052978901140.europe-west2.run.app": "http://localhost:8080"
}