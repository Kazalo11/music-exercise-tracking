import {
  Card,
  CardBody,
  CardHeader,
  Heading,
  Image,
  ListItem,
  UnorderedList,
} from "@chakra-ui/react";
import { formatUnixTimestamp } from "../../Date";

export type MusicCardProps = {
  name: string;
  artist: {
    "#text": string;
  };
  album: {
    "#text": string;
  };
  image: {
    "#text": string;
  }[];
  "@attr": {
    nowplaying: string;
  };
  url: string;
  date: {
    "#text": string;
    uts: string;
  };
};

export function MusicCard(props: MusicCardProps) {
  return (
    props["@attr"]["nowplaying"] !== "true" && (
      <Card>
        <CardHeader>
          <Heading size="md">{props.name}</Heading>
        </CardHeader>
        <CardBody>
          <Image
            objectFit="cover"
            maxW={{ base: "100%", sm: "200px" }}
            src={props.image[props.image.length - 1]["#text"]}
            alt="Music Album"
          />
          <UnorderedList>
            <ListItem>Album: {props.album["#text"]}</ListItem>
            <ListItem>Artist: {props.artist["#text"]}</ListItem>
            <ListItem>
              Listened to at: {formatUnixTimestamp(props.date.uts)}
            </ListItem>
          </UnorderedList>
        </CardBody>
      </Card>
    )
  );
}
