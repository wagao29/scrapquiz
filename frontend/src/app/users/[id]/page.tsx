import Profile from "@/components/profile";
import { QuizCard } from "@/components/quiz-card";
import QuizPagination from "@/components/quiz-pagination";
import { fetchQuizCounts, fetchQuizzesByUserId, fetchUser } from "@/lib/apis";
import { auth } from "@/lib/auth";
import { FETCH_QUIZZES_LIMIT } from "@/lib/constants";
import { notFound } from "next/navigation";

export default async function Page({
  params,
  searchParams,
}: {
  params: { id: string };
  searchParams: { [key: string]: string };
}) {
  const pageNum = Number(searchParams.page || 1);
  const [user, quizCounts, quizzes] = await Promise.all([
    fetchUser(params.id),
    fetchQuizCounts(params.id),
    fetchQuizzesByUserId(params.id, (pageNum - 1) * FETCH_QUIZZES_LIMIT),
  ]);
  if (!user) {
    notFound();
  }

  const session = await auth();

  return (
    <main className="flex flex-col gap-4 items-center p-4">
      <Profile
        userName={user.name}
        userImage={user.avatarUrl}
        quizCounts={quizCounts}
        isLoginUser={user.id === session?.user?.id}
      />
      <div className="flex flex-col gap-4 mt-8 min-w-[350px] max-w-[850px] w-full">
        <span className="font-semibold">作成したクイズ</span>
        {quizzes?.quizzes.map((quiz) => (
          <QuizCard
            key={quiz.id}
            quiz={quiz}
            deletable={quiz.userId === session?.user?.id}
          />
        ))}
      </div>
      <QuizPagination
        basePath={`/users/${params.id}/`}
        pageNum={pageNum}
        quizCounts={quizCounts}
      />
    </main>
  );
}
