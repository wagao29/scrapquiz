import { SkeletonQuizCard } from "@/components/quiz-card";
import QuizPagination from "@/components/quiz-pagination";
import { Skeleton } from "@/components/ui/skeleton";

export default function Loading() {
  return (
    <main className="flex flex-col gap-4 items-center mx-4 mt-4">
      <div className="flex justify-between w-full max-w-[850px]">
        <div className="flex items-center gap-2">
          <Skeleton className="h-20 w-20 rounded-full" />
          <div className="flex flex-col gap-2">
            <Skeleton className="h-6 w-20" />
            <Skeleton className="h-4 w-28" />
          </div>
        </div>
      </div>
      <div className="flex flex-col gap-4 mt-8 min-w-[350px] max-w-[850px] w-full">
        <span className="font-semibold">作成したクイズ</span>
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
      </div>
      <QuizPagination
        basePath="/"
        pageNum={1}
        quizCounts={1}
        className="my-5"
      />
    </main>
  );
}
