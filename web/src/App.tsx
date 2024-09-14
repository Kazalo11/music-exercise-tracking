import { useEffect, useState } from "react";
import "./App.css";
import { Heading, Text } from "@chakra-ui/react";
import LoginPage from "./components/LoginPage";

function App() {
  const [isAuthenticated, setIsAuthenticated] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    fetch("http://localhost:8080/v1/auth")
      .then((response) => response.json())
      .then((data) => {
        setIsAuthenticated(data.isAuthenticated);
      })
      .catch((error) => {
        console.error("Error checking authentication status:", error);
        setError("Error checking authentication status.");
      });
  }, []);

  return (
    <div className="App">
      <Heading>Welcome to Spotify + Strava Tracking</Heading>
      {error && <Text color="red.500">{error}</Text>}
      {isAuthenticated ? (
        <Text>Welcome back! You are already authenticated.</Text>
      ) : (
        <LoginPage />
      )}
    </div>
  );
}

export default App;
