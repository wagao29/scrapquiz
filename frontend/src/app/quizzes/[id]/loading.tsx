import { SkeletonQuizCard } from "@/components/quiz-card";
import { Button } from "@/components/ui/button";
import Link from "next/link";

export default function Loading() {
  return (
    <main className="flex flex-col gap-4 items-center mt-10 px-4">
      <SkeletonQuizCard />
      <Button className="mt-10" asChild>
        <Link href="/">トップへ戻る</Link>
      </Button>
    </main>
  );
}
