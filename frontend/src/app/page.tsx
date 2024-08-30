import { QuizCard } from "@/components/quiz-card";
import { fetchQuiz } from "@/lib/apis";
import { notFound } from "next/navigation";

export default async function Page() {
  const quiz = await fetchQuiz("01J62PPM6FDXWE3R1NFFN1DA94");
  if (!quiz) notFound();

  return (
    <div className="flex flex-col gap-4 items-center m-4">
      <QuizCard quiz={quiz} />
      <QuizCard quiz={quiz} />
      <QuizCard quiz={quiz} />
      <QuizCard quiz={quiz} deletable />
      <QuizCard quiz={quiz} deletable />
    </div>
  );
}
