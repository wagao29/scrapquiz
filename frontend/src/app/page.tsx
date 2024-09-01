import { QuizCard } from "@/components/quiz-card";
import { fetchLatestQuizzes } from "@/lib/apis";
import { notFound } from "next/navigation";

export default async function Page() {
  const latestQuizzes = await fetchLatestQuizzes(0);
  if (!latestQuizzes) notFound();

  return (
    <main className="flex flex-col gap-4 items-center m-4">
      {latestQuizzes.quizzes.map((quiz) => (
        <QuizCard quiz={quiz} />
      ))}
    </main>
  );
}
