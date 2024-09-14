import { Card, CardHeader, Heading, useDisclosure } from "@chakra-ui/react";
import { Activity, DropDown } from "./Dropdown";
import { useState } from "react";
import { SpotifyDrawer } from "./drawer/SpotifyDrawer";
export function MainPage() {
  const { isOpen, onOpen, onClose } = useDisclosure();
  const [selectedActivity, setSelectedActivity] = useState<Activity | null>(
    null
  );
  const handleSelectChange = (activity: Activity) => {
    setSelectedActivity(activity);
    onOpen();
  };
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
          />
        )}
      </Card>
    </>
  );
}
