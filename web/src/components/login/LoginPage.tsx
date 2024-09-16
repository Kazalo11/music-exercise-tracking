import { ButtonGroup, Text } from "@chakra-ui/react";
import { StatusCodes } from "http-status-codes";
import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { loadConfig } from "../../../config/Config";
import { LoginButton } from "./LoginButton";

export default function LoginPage() {
  const [stravaAuthUrl, setStravaAuthUrl] = useState<string | null>(null);
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  const [error, _setError] = useState<string | null>(null);
  const navigate = useNavigate();

  useEffect(() => {
    const getCookie = async () => {
      const config = await loadConfig("../../config/Config");
      const response = await fetch(
        `http://${config.server.host}/v1/strava/access_token`,
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
      const config = await loadConfig("../../config/Config");
      const response = await fetch(
        `http://${config.server.host}/v1/strava/auth`
      );

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
