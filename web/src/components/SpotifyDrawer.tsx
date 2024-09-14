import {
  Button,
  Drawer,
  DrawerContent,
  DrawerCloseButton,
  DrawerHeader,
  DrawerBody,
  DrawerFooter,
} from "@chakra-ui/react";

type SpotifyDrawerProps = {
  isOpen: boolean;
  onClose: () => void;
  name: string;
  start_date: string;
};

export function SpotifyDrawer({
  isOpen,
  onClose,
  name,
  start_date,
}: SpotifyDrawerProps) {
  return (
    <Drawer isOpen={isOpen} placement="right" onClose={onClose}>
      <DrawerContent>
        <DrawerCloseButton />
        <DrawerHeader>
          Spotify songs listened to during {name} at {start_date}
        </DrawerHeader>

        <DrawerBody>{/* Additional drawer content can go here */}</DrawerBody>

        <DrawerFooter>
          <Button variant="outline" mr={3} onClick={onClose}>
            Close
          </Button>
        </DrawerFooter>
      </DrawerContent>
    </Drawer>
  );
}
