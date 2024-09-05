import { auth, signIn } from "@/lib/auth";
import { Avatar, AvatarImage } from "@radix-ui/react-avatar";
import { PlusSquare } from "lucide-react";
import Link from "next/link";
import { Button } from "./ui/button";

export default async function Header() {
  const session = await auth();

  return (
    <header className="flex justify-between px-3 py-4">
      <Link href="/" className="text-xl font-semibold mt-1">
        <h1>scrapquiz</h1>
      </Link>
      {session ? (
        <div className="flex items-center gap-2">
          <Button size="icon" variant="ghost" asChild>
            <Link href="/quizzes/create">
              <PlusSquare width={25} height={25} strokeWidth={1} />
            </Link>
          </Button>
          <Button
            variant="outline"
            className="rounded-full"
            size="icon"
            asChild
          >
            <Link href={`/users/${session?.user?.id}`}>
              <Avatar>
                <AvatarImage
                  width={30}
                  height={30}
                  className="rounded-full"
                  src={session?.user?.image ?? ""}
                  alt={`${session?.user?.name}のアイコン画像`}
                />
              </Avatar>
            </Link>
          </Button>
        </div>
      ) : (
        <form
          action={async () => {
            "use server";
            await signIn("google", { redirectTo: "/" });
          }}
        >
          <Button type="submit">ログイン</Button>
        </form>
      )}
    </header>
  );
}
