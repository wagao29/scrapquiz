import Profile from "@/components/profile";
import { QuizCard } from "@/components/quiz-card";
import {
  Pagination,
  PaginationContent,
  PaginationEllipsis,
  PaginationItem,
  PaginationLink,
  PaginationNext,
  PaginationPrevious,
} from "@/components/ui/pagination";
import { fetchQuizCounts, fetchQuizzesByUserId } from "@/lib/apis";
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
  const [quizCounts, quizzes] = await Promise.all([
    fetchQuizCounts(params.id),
    fetchQuizzesByUserId(params.id, (pageNum - 1) * FETCH_QUIZZES_LIMIT),
  ]);
  if ((pageNum - 1) * FETCH_QUIZZES_LIMIT > quizCounts || !quizzes) {
    notFound();
  }

  const session = await auth();

  return (
    <main className="flex flex-col gap-4 items-center p-4">
      <Profile
        userName={session?.user?.name || ""}
        userImage={session?.user?.image || ""}
        quizCounts={quizCounts}
        isLoginUser={session?.user?.id === params.id}
      />
      <div className="flex flex-col gap-2 mt-8 min-w-[350px] max-w-[850px] w-full">
        <span className="font-semibold">作成したクイズ</span>
        {quizzes.quizzes.map((quiz) => (
          <QuizCard quiz={quiz} deletable />
        ))}
      </div>
      <Pagination>
        <PaginationContent>
          {pageNum > 1 && (
            <>
              <PaginationItem>
                <PaginationPrevious
                  href={`/users/${params.id}/?page=${pageNum - 1}`}
                />
              </PaginationItem>
              <PaginationItem>
                <PaginationEllipsis />
              </PaginationItem>
            </>
          )}
          <PaginationItem>
            <PaginationLink href={`/users/${params.id}/?page=${pageNum}`}>
              {pageNum}
            </PaginationLink>
          </PaginationItem>
          {pageNum * FETCH_QUIZZES_LIMIT < quizCounts && (
            <>
              <PaginationItem>
                <PaginationEllipsis />
              </PaginationItem>
              <PaginationItem>
                <PaginationNext
                  href={`/users/${params.id}/?page=${pageNum + 1}`}
                />
              </PaginationItem>
            </>
          )}
        </PaginationContent>
      </Pagination>
    </main>
  );
}
