import { LockIcon, UnlockIcon } from "@chakra-ui/icons";
import {
  IconButton,
  Input,
  InputGroup,
  InputRightElement,
  VStack,
} from "@chakra-ui/react";
import React, { useState } from "react";

interface UserNameInputProps {
  onChange: (username: string, isSubmit: boolean) => void;
}

export function UserNameInput({ onChange }: UserNameInputProps) {
  const [username, setUsername] = useState("");
  const [isEditable, setIsEditable] = useState(true);

  const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setUsername(event.target.value);
    onChange("", false);
  };

  const handleSubmit = () => {
    setIsEditable(!isEditable);
    onChange(username, isEditable);
  };

  return (
    <VStack spacing={4}>
      <InputGroup>
        <Input
          value={username}
          onChange={handleChange}
          isReadOnly={!isEditable}
          placeholder="Enter username"
          backgroundColor={!isEditable ? "gray.200" : "white"}
          borderColor={!isEditable ? "gray.300" : "gray.300"}
          _placeholder={{ color: !isEditable ? "gray.500" : "gray.400" }}
        />
        <InputRightElement>
          <IconButton
            aria-label="Save username"
            icon={isEditable ? <UnlockIcon /> : <LockIcon />}
            onClick={handleSubmit}
          />
        </InputRightElement>
      </InputGroup>
    </VStack>
  );
}
