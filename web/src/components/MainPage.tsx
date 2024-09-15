import {
  Card,
  CardBody,
  CardHeader,
  Heading,
  StackDivider,
  useDisclosure,
  VStack,
} from "@chakra-ui/react";
import { StatusCodes } from "http-status-codes";
import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { SpotifyDrawer } from "./drawer/SpotifyDrawer";
import { Activity, DropDown } from "./Dropdown";
import { UserNameInput } from "./input/UserNameInput";

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
          <Heading size="md">Type in your last fm username and submit</Heading>
        </CardHeader>
        <CardBody>
          <VStack spacing={4} divider={<StackDivider borderColor="gray.200" />}>
            <UserNameInput onChange={handleUsernameChange} />
            {submit && (
              <VStack spacing={5}>
                <Heading size="md">
                  Select a run to see the songs listened to during it
                </Heading>
                <DropDown onSelectChange={handleSelectChange} />
              </VStack>
            )}
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
