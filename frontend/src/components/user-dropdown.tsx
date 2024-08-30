"use client";

import { Avatar, AvatarImage } from "@radix-ui/react-avatar";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@radix-ui/react-dropdown-menu";

import Link from "next/link";
import { useRouter } from "next/navigation";
import { Button } from "./ui/button";

type Props = {
  userId: string;
  userName: string;
  avatarUrl: string;
};

export function UserDropdown({ userId, userName, avatarUrl }: Props) {
  const router = useRouter();

  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <Button variant="ghost" size="icon">
          <Avatar>
            <AvatarImage
              width={35}
              height={35}
              className="rounded-full"
              src={avatarUrl}
              alt={`${userName}のアイコン画像`}
            />
          </Avatar>
        </Button>
      </DropdownMenuTrigger>
      <DropdownMenuContent
        align="end"
        sideOffset={5}
        className="border rounded-sm bg-white text-sm shadow-lg"
      >
        <DropdownMenuItem asChild>
          <Link
            href={`/users/${userId}`}
            className="flex justify-center px-8 py-2 border-b"
          >
            マイページ
          </Link>
        </DropdownMenuItem>
        <DropdownMenuItem asChild>
          <Link
            href={"/quizzes/create"}
            className="flex justify-center px-8 py-2 border-b"
          >
            クイズ作成
          </Link>
        </DropdownMenuItem>
        <DropdownMenuItem
          className="flex justify-center px-8 py-2 cursor-pointer"
          onSelect={() => console.log("ログアウト")}
        >
          ログアウト
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  );
}
