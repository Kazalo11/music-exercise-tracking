import {
  Button,
  Drawer,
  DrawerContent,
  DrawerCloseButton,
  DrawerHeader,
  DrawerFooter,
} from "@chakra-ui/react";
import { SpotifyDrawerBody } from "./SpotifyDrawerBody";

type SpotifyDrawerProps = {
  isOpen: boolean;
  onClose: () => void;
  name: string;
  start_date: string;
  id: string;
};

export function SpotifyDrawer({
  isOpen,
  onClose,
  name,
  start_date,
  id,
}: SpotifyDrawerProps) {
  return (
    <Drawer isOpen={isOpen} placement="right" onClose={onClose}>
      <DrawerContent>
        <DrawerCloseButton />
        <DrawerHeader>
          Spotify songs listened to during {name} at {start_date}
        </DrawerHeader>

        <SpotifyDrawerBody id={id} />
        <DrawerFooter>
          <Button variant="outline" mr={3} onClick={onClose}>
            Close
          </Button>
        </DrawerFooter>
      </DrawerContent>
    </Drawer>
  );
}
