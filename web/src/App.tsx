import { useState, useEffect } from "react";
import "./App.css";
import { Heading, Text } from "@chakra-ui/react";
import LoginPage from "./components/login/LoginPage";
import { MainPage } from "./components/MainPage";

function App() {
  const [isAuthenticated, setIsAuthenticated] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);

  //now with using last_fm just need to check if the strava auth is successful
  useEffect(() => {
    const checkAuthStatus = () => {
      fetch("http://localhost:8080/v1/auth")
        .then((response) => response.json())
        .then((data) => {
          setIsAuthenticated(data.isAuthenticated);
        })
        .catch((error) => {
          console.error("Error checking authentication status:", error);
          setError("Error checking authentication status.");
        });
    };
    checkAuthStatus();
  }, []);

  return (
    <div className="App">
      <Heading>Welcome to Spotify + Strava Tracking</Heading>
      {error && <Text color="red.500">{error}</Text>}
      {isAuthenticated ? <MainPage /> : <LoginPage />}
    </div>
  );
}

export default App;
