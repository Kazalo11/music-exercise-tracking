// components/DropDown.tsx
import { Box, Select, Spinner } from "@chakra-ui/react";
import { ChangeEvent, useEffect, useState } from "react";
import { formatISOString } from "../Date";

export type Activity = {
  id: string;
  name: string;
  start_date: string;
  finish_date: string;
};

type DropDownProps = {
  onSelectChange: (activity: Activity) => void;
};

export function DropDown({ onSelectChange }: DropDownProps) {
  const [options, setOptions] = useState<Activity[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string>("");

  useEffect(() => {
    const fetchOptions = async () => {
      try {
        const response = await fetch(
          "http://localhost:8080/v1/strava/activities"
        );
        if (!response.ok) {
          throw new Error("Network response was not ok");
        }
        const data = await response.json();
        setOptions(data.activities);
      } catch (error: unknown) {
        setError((error as Error).message);
      } finally {
        setLoading(false);
      }
    };

    fetchOptions();
  }, []);

  if (loading) return <Spinner />;
  if (error) return <Box color="red.500">{`Error: ${error}`}</Box>;

  const handleSelectChange = (e: ChangeEvent<HTMLSelectElement>) => {
    const selectedActivity = options.find(
      (option) => option.id == e.target.value
    );
    if (selectedActivity) {
      onSelectChange(selectedActivity);
    }
  };

  return (
    <Select placeholder="Select option" onChange={handleSelectChange}>
      {options.map((option) => (
        <option key={option.id} value={option.id}>
          {option.name}, with start: {formatISOString(option.start_date)}
        </option>
      ))}
    </Select>
  );
}
