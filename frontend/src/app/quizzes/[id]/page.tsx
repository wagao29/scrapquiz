import { QuizCard } from "@/components/quiz-card";
import { Button } from "@/components/ui/button";
import { fetchQuiz } from "@/lib/apis";
import Link from "next/link";
import { notFound } from "next/navigation";

export default async function Page({ params }: { params: { id: string } }) {
  const quiz = await fetchQuiz(params.id);
  if (!quiz) notFound();

  return (
    <main className="flex flex-col gap-4 items-center mt-10 px-4">
      <QuizCard key={quiz.id} quiz={quiz} />
      <Button className="mt-10" asChild>
        <Link href="/">トップへ戻る</Link>
      </Button>
    </main>
  );
}
