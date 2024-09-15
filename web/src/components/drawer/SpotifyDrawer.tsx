import {
  Button,
  Drawer,
  DrawerCloseButton,
  DrawerContent,
  DrawerFooter,
  DrawerHeader,
} from "@chakra-ui/react";
import { formatISOString } from "../../Date";
import { SpotifyDrawerBody } from "./SpotifyDrawerBody";

type SpotifyDrawerProps = {
  isOpen: boolean;
  onClose: () => void;
  name: string;
  start_date: string;
  id: string;
  finish_date: string;
  username: string;
};

export function SpotifyDrawer({
  isOpen,
  onClose,
  name,
  start_date,
  finish_date,
  username,
}: SpotifyDrawerProps) {
  return (
    <Drawer isOpen={isOpen} placement="right" onClose={onClose}>
      <DrawerContent>
        <DrawerCloseButton />
        <DrawerHeader>
          Spotify songs listened to during {name} at{" "}
          {formatISOString(start_date)}
        </DrawerHeader>

        <SpotifyDrawerBody
          user_name={username}
          start_date={start_date}
          finish_date={finish_date}
        />
        <DrawerFooter>
          <Button variant="outline" mr={3} onClick={onClose}>
            Close
          </Button>
        </DrawerFooter>
      </DrawerContent>
    </Drawer>
  );
}
