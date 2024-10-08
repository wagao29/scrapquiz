import { signOut } from "@/lib/auth";
import { Avatar, AvatarImage } from "@radix-ui/react-avatar";
import { Button } from "./ui/button";
type Props = {
  userName: string;
  userImage: string;
  quizCounts: number;
  isLoginUser: boolean;
};

export default async function Profile({
  userName,
  userImage,
  quizCounts,
  isLoginUser,
}: Props) {
  return (
    <div className="flex justify-between w-full max-w-[850px]">
      <div className="flex items-center gap-1">
        <Avatar>
          <AvatarImage
            width={80}
            height={80}
            className="rounded-full border"
            src={userImage}
            alt={`${userName}のアイコン画像`}
          />
        </Avatar>
        <div className="flex flex-col gap-1">
          <span className="text-lg font-semibold">{userName}</span>
          <span className="text-sm">作成数: {quizCounts}</span>
        </div>
      </div>
      {isLoginUser && (
        <form
          action={async () => {
            "use server";
            await signOut({ redirectTo: "/" });
          }}
        >
          <Button
            size={"sm"}
            variant={"outline"}
            className="text-xs rounded-full"
            type="submit"
          >
            ログアウト
          </Button>
        </form>
      )}
    </div>
  );
}
