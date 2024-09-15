import { useEffect, useState } from "react";
import { ButtonGroup, Text } from "@chakra-ui/react";
import { LoginButton } from "./LoginButton";
import { useNavigate } from "react-router-dom";

export default function LoginPage() {
  const [stravaAuthUrl, setStravaAuthUrl] = useState<string | null>(null);
  const [error, setError] = useState<string | null>(null);
  const navigate = useNavigate();

  useEffect(() => {
    const checkLoginStatus = () => {
      const isLoggedIn = localStorage.getItem("isLoggedIn");
      if (isLoggedIn == "true") {
        navigate("/main");
      }
    };

    checkLoginStatus();
  }, [navigate]);

  return (
    <div>
      {error && <Text color="red.500">{error}</Text>}
      <ButtonGroup spacing="6" mt="24px">
        <LoginButton
          link={stravaAuthUrl || "#"}
          text={"Login to Strava here"}
          onClick={() => localStorage.setItem("isLoggedIn", "true")}
        />
      </ButtonGroup>
    </div>
  );
}
