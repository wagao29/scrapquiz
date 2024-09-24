import { QuizCard } from "@/components/quiz-card";
import QuizPagination from "@/components/quiz-pagination";
import TabMenu from "@/components/tab-menu";
import { fetchLatestQuizzes, fetchQuizCounts } from "@/lib/apis";
import { FETCH_QUIZZES_LIMIT } from "@/lib/constants";
import { sleep } from "@/lib/utils";
import { notFound } from "next/navigation";

export default async function Page({
  searchParams,
}: {
  searchParams: { [key: string]: string };
}) {
  const pageNum = Number(searchParams.page || 1);
  const [quizCounts, quizzes] = await Promise.all([
    fetchQuizCounts(),
    fetchLatestQuizzes((pageNum - 1) * FETCH_QUIZZES_LIMIT),
  ]);
  if ((pageNum - 1) * FETCH_QUIZZES_LIMIT > quizCounts || !quizzes) {
    notFound();
  }

  await sleep(1000);

  return (
    <main className="flex flex-col gap-4 items-center mx-4">
      <TabMenu selectValue="latest" className="mb-4" />
      {quizzes.quizzes.map((quiz) => (
        <QuizCard key={quiz.id} quiz={quiz} />
      ))}
      <QuizPagination
        basePath="/"
        pageNum={pageNum}
        quizCounts={quizCounts}
        className="my-5"
      />
    </main>
  );
}
