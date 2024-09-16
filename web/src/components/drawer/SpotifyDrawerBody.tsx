import { DrawerBody } from "@chakra-ui/react";
import { useEffect, useState } from "react";
import { getBackendHost } from "../../../config/Config";
import { MusicCard, MusicCardProps } from "../card/MusicCard";

type SpotifyDrawerBodyProps = {
  user_name: string;
  start_date: string;
  finish_date: string;
};

export function SpotifyDrawerBody({
  user_name,
  start_date,
  finish_date,
}: SpotifyDrawerBodyProps) {
  const [songs, setSongs] = useState<MusicCardProps[]>([]);
  useEffect(() => {
    const getActivityInfo = async () => {
      const response = await fetch(`${getBackendHost()}/v1/lastfm/tracks`, {
        method: "POST",
        headers: {
          Accept: "application/json",
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          user_name,
          start: start_date,
          end: finish_date,
        }),
      });
      const songsData = await response.json();
      console.log(songsData);

      setSongs(songsData.tracks);
    };

    getActivityInfo();
  }, [user_name, start_date, finish_date]);
  return (
    <DrawerBody>
      {songs.map((song, index) => (
        <MusicCard key={index} {...song} />
      ))}
    </DrawerBody>
  );
}
