import { useEffect, useState } from "react";
import { ButtonGroup, Spinner, Text } from "@chakra-ui/react";
import { LoginButton } from "./LoginButton";

function LoginPage() {
  const [spotifyAuthUrl, setSpotifyAuthUrl] = useState<string | null>(null);
  const [stravaAuthUrl, setStravaAuthUrl] = useState<string | null>(null);
  const [loadingSpotify, setLoadingSpotify] = useState<boolean>(true);
  const [loadingStrava, setLoadingStrava] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    fetch("http://localhost:8080/v1/spotify/auth")
      .then((response) => response.json())
      .then((data) => {
        setSpotifyAuthUrl(data.url);
        setLoadingSpotify(false);
      })
      .catch((error) => {
        console.error("Error fetching Spotify auth URL:", error);
        setError("Error fetching Spotify auth URL.");
        setLoadingSpotify(false);
      });
  }, []);

  useEffect(() => {
    fetch("http://localhost:8080/v1/strava/auth")
      .then((response) => response.json())
      .then((data) => {
        setStravaAuthUrl(data.url);
        setLoadingStrava(false);
      })
      .catch((error) => {
        console.error("Error fetching Strava auth URL:", error);
        setError("Error fetching Strava auth URL.");
        setLoadingStrava(false);
      });
  }, []);

  return (
    <div>
      {error && <Text color="red.500">{error}</Text>}
      <ButtonGroup spacing="6" mt="24px">
        {loadingSpotify ? (
          <Spinner size="lg" />
        ) : (
          <LoginButton
            link={spotifyAuthUrl || "#"}
            text={"Login to Spotify here"}
          />
        )}
        {loadingStrava ? (
          <Spinner size="lg" />
        ) : (
          <LoginButton
            link={stravaAuthUrl || "#"}
            text={"Login to Strava here"}
          />
        )}
      </ButtonGroup>
    </div>
  );
}

export default LoginPage;
