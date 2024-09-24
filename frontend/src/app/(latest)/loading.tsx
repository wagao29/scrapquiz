import { SkeletonQuizCard } from "@/components/quiz-card";
import QuizPagination from "@/components/quiz-pagination";

import TabMenu from "@/components/tab-menu";

export default function Loading() {
  return (
    <main className="flex flex-col gap-4 items-center mx-4">
      <TabMenu selectValue="latest" className="mb-4" />
      <SkeletonQuizCard />
      <SkeletonQuizCard />
      <SkeletonQuizCard />
      <SkeletonQuizCard />
      <SkeletonQuizCard />
      <SkeletonQuizCard />
      <SkeletonQuizCard />
      <SkeletonQuizCard />
      <SkeletonQuizCard />
      <SkeletonQuizCard />
      <QuizPagination
        basePath="/"
        pageNum={1}
        quizCounts={1}
        className="my-5"
      />
    </main>
  );
}
