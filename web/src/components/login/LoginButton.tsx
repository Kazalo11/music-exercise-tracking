import { Button, Link } from "@chakra-ui/react";

type LoginButtonProps = {
  link: string;
  text: string;
  onClick: () => void;
};

export function LoginButton({ link, text, onClick }: LoginButtonProps) {
  return (
    <Link href={link} isExternal>
      <Button size="lg" mt="24px" colorScheme="teal" onClick={onClick}>
        {text}
      </Button>
    </Link>
  );
}
