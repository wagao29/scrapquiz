import { QuizCard } from "@/components/quiz-card";
import TabMenu from "@/components/tab-menu";
import UpdateButton from "@/components/update-button";
import { fetchRandomQuizzes } from "@/lib/apis";
import { notFound } from "next/navigation";

export default async function Page() {
  const quizzes = await fetchRandomQuizzes();
  if (!quizzes) {
    notFound();
  }

  return (
    <main className="flex flex-col gap-4 items-center mx-4">
      <TabMenu selectValue="random" className="mb-4" />
      {quizzes.quizzes.map((quiz) => (
        <QuizCard key={quiz.id} quiz={quiz} />
      ))}
      <UpdateButton className="my-5" />
    </main>
  );
}
