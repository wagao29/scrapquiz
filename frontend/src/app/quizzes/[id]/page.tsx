import { QuizCard } from "@/components/quiz-card";
import { fetchQuiz } from "@/lib/apis";
import { notFound } from "next/navigation";

export default async function Page({ params }: { params: { id: string } }) {
  const quiz = await fetchQuiz(params.id);
  if (!quiz) notFound();

  return (
    <main className="flex flex-col gap-4 items-center mt-10 px-4">
      <QuizCard key={quiz.id} quiz={quiz} />
    </main>
  );
}
