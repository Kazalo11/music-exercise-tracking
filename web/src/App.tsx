import { useEffect, useState } from "react";
import "./App.css";

function App() {
  const [authUrl, setAuthUrl] = useState<string | null>(null);

  useEffect(() => {
    fetch("http://localhost:8080/auth")
      .then((response) => response.json())
      .then((data) => setAuthUrl(data.url))
      .catch((error) => console.error("Error fetching auth URL:", error));
  }, []);

  useEffect(() => {
    if (authUrl) {
      window.location.href = authUrl;
    }
  }, [authUrl]);

  return (
    <div className="App">
      <h1>Authentication</h1>
      {authUrl ? (
        <a href={authUrl} target="_blank" rel="noopener noreferrer">
          Click here to authenticate
        </a>
      ) : (
        <p>Loading authentication link...</p>
      )}
    </div>
  );
}

export default App;
