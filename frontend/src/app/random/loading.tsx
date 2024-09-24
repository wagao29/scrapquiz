import { SkeletonQuizCard } from "@/components/quiz-card";
import TabMenu from "@/components/tab-menu";
import UpdateButton from "@/components/update-button";

export default function Loading() {
  return (
    <main className="flex flex-col gap-4 items-center mx-4">
      <TabMenu selectValue="random" className="mb-4" />
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
      <UpdateButton className="my-5" />
    </main>
  );
}
