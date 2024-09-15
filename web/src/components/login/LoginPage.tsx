import { useEffect, useState } from "react";
import { ButtonGroup, Text } from "@chakra-ui/react";
import { LoginButton } from "./LoginButton";
import { useNavigate } from "react-router-dom";
import { StatusCodes } from "http-status-codes";

export default function LoginPage() {
  const [stravaAuthUrl, setStravaAuthUrl] = useState<string | null>(null);
  const [error, setError] = useState<string | null>(null);
  const navigate = useNavigate();

  useEffect(() => {
    const getCookie = async () => {
      const response = await fetch(
        "http://localhost:8080/v1/strava/access_token",
        {
          credentials: "include",
        }
      );

      if (response.status == StatusCodes.OK) {
        navigate("/");
      }
    };
    getCookie();
  }, [navigate]);

  useEffect(() => {
    const getAuthUrl = async () => {
      const response = await fetch("http://localhost:8080/v1/strava/auth");

      const authUrl = await response.json();
      setStravaAuthUrl(authUrl.url);
    };
    getAuthUrl();
  });

  return (
    <div>
      {error && <Text color="red.500">{error}</Text>}
      <ButtonGroup spacing="6" mt="24px">
        <LoginButton
          link={stravaAuthUrl || "#"}
          text={"Login to Strava here"}
        />
      </ButtonGroup>
    </div>
  );
}
