import { Button, Link } from "@chakra-ui/react";

type LoginButtonProps = {
  link: string;
  text: string;
};

export function LoginButton({ link, text }: LoginButtonProps) {
  return (
    <Link href={link}>
      <Button size="lg" mt="24px" colorScheme="teal">
        {text}
      </Button>
    </Link>
  );
}
