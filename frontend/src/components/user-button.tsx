import { cn } from "@/lib/utils";
import { Avatar, AvatarImage } from "@radix-ui/react-avatar";
import Link from "next/link";
import { Button } from "./ui/button";

type Props = React.ComponentProps<typeof Button> & {
  id: string;
  name: string;
  avatarUrl: string;
};

export function UserButton({
  id,
  name,
  avatarUrl,
  className,
  ...props
}: Props) {
  return (
    <Button variant="link" size="sm" asChild {...props}>
      <Link href={`/users/${id}`} className={cn("space-x-1.5", className)}>
        <Avatar>
          <AvatarImage
            width={30}
            height={30}
            className="rounded-full"
            src={avatarUrl}
            alt={`${name}のアイコン画像`}
          />
        </Avatar>
        <span>{name}</span>
      </Link>
    </Button>
  );
}
