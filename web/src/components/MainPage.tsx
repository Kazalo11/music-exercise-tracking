import {
  Card,
  CardBody,
  CardHeader,
  Heading,
  StackDivider,
  useDisclosure,
  VStack,
} from "@chakra-ui/react";
import { Activity, DropDown } from "./Dropdown";
import { useEffect, useState } from "react";
import { SpotifyDrawer } from "./drawer/SpotifyDrawer";
import { useNavigate } from "react-router-dom";
import { StatusCodes } from "http-status-codes";
import UsernameInput from "./input/UserNameInput";

export function MainPage() {
  const { isOpen, onOpen, onClose } = useDisclosure();
  const [selectedActivity, setSelectedActivity] = useState<Activity | null>(
    null
  );
  const [username, setUsername] = useState<string>("");
  const [submit, setSubmit] = useState<boolean>(false);
  const navigate = useNavigate();
  const handleSelectChange = (activity: Activity) => {
    setSelectedActivity(activity);
    onOpen();
  };

  const handleUsernameChange = (username: string, isSubmit: boolean) => {
    setUsername(username);
    setSubmit(isSubmit);
  };

  useEffect(() => {
    const getCookie = async () => {
      const response = await fetch(
        "http://localhost:8080/v1/strava/access_token",
        {
          credentials: "include",
        }
      );

      if (response.status == StatusCodes.NOT_FOUND) {
        navigate("/login");
      }
    };
    getCookie();
  }, [navigate]);

  return (
    <>
      <Card>
        <CardHeader>
          <Heading size="md">
            Type in your last fm username, and choose a run to see the songs
            listened to during the run
          </Heading>
        </CardHeader>
        <CardBody>
          <VStack spacing={4} divider={<StackDivider borderColor="gray.200" />}>
            <UsernameInput onChange={handleUsernameChange} />
            {submit && <DropDown onSelectChange={handleSelectChange} />}
            {selectedActivity && (
              <SpotifyDrawer
                isOpen={isOpen}
                onClose={onClose}
                name={selectedActivity?.name}
                start_date={selectedActivity?.start_date}
                id={selectedActivity?.id}
                finish_date={selectedActivity?.finish_date}
                username={username}
              />
            )}
          </VStack>
        </CardBody>
      </Card>
    </>
  );
}
