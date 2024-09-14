import { DrawerBody } from "@chakra-ui/react";
import { useEffect, useState } from "react";

type SpotifyDrawerBodyProps = {
  id: string;
  start_date: string;
  finish_date: string;
};
type Song = {
  name: string;
  artists: string[];
};

export function SpotifyDrawerBody({
  id,
  start_date,
  finish_date,
}: SpotifyDrawerBodyProps) {
  const [song, setSong] = useState<Song | null>(null);
  useEffect(() => {
    const getActivityInfo = async (id: string) => {
      const response = await fetch("http://localhost:8080/v1/spotify/songs", {
        method: "POST",
        headers: {
          Accept: "application/json",
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ id, start: start_date, end: finish_date }),
      });

      setSong(await response.json());
    };

    getActivityInfo(id);
  }, [id, start_date, finish_date]);
  return <DrawerBody>{song?.name}</DrawerBody>;
}
