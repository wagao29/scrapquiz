import { Button } from "@/components/ui/button";
import { auth, signOut } from "@/lib/auth";
import { redirect } from "next/navigation";

export default async function Page({ params }: { params: { id: string } }) {
  const session = await auth();

  return (
    <main className="flex flex-col gap-4 items-center p-4">
      <h1>{`ユーザー詳細画面 (id: ${params.id})`}</h1>
      {session && (
        <form
          action={async () => {
            "use server";
            await signOut();
            redirect("/");
          }}
        >
          <Button type="submit">ログアウト</Button>
        </form>
      )}
    </main>
  );
}
