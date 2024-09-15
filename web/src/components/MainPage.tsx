import { Card, CardHeader, Heading, useDisclosure } from "@chakra-ui/react";
import { Activity, DropDown } from "./Dropdown";
import { useEffect, useState } from "react";
import { SpotifyDrawer } from "./drawer/SpotifyDrawer";
import { useNavigate } from "react-router-dom";

export function MainPage() {
  const { isOpen, onOpen, onClose } = useDisclosure();
  const [selectedActivity, setSelectedActivity] = useState<Activity | null>(
    null
  );
  const navigate = useNavigate();
  const handleSelectChange = (activity: Activity) => {
    setSelectedActivity(activity);
    onOpen();
  };

  // useEffect(() => {
  //   navigate("/login");
  // });
  return (
    <>
      <Card>
        <CardHeader>
          <Heading size="md">
            Choose a run to see the songs listened to during the run
          </Heading>
        </CardHeader>

        <DropDown onSelectChange={handleSelectChange} />
        {selectedActivity && (
          <SpotifyDrawer
            isOpen={isOpen}
            onClose={onClose}
            name={selectedActivity?.name}
            start_date={selectedActivity?.start_date}
            id={selectedActivity?.id}
            finish_date={selectedActivity?.finish_date}
          />
        )}
      </Card>
    </>
  );
}
