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
import { fetchLatestQuizzes, fetchQuizCounts } from "@/lib/apis";
import { FETCH_QUIZZES_LIMIT } from "@/lib/constants";
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

  return (
    <main className="flex flex-col gap-4 items-center m-4">
      {quizzes.quizzes.map((quiz) => (
        <QuizCard quiz={quiz} />
      ))}
      <Pagination>
        <PaginationContent>
          {pageNum > 1 && (
            <>
              <PaginationItem>
                <PaginationPrevious href={`/?page=${pageNum - 1}`} />
              </PaginationItem>
              <PaginationItem>
                <PaginationEllipsis />
              </PaginationItem>
            </>
          )}
          <PaginationItem>
            <PaginationLink href={`/?page=${pageNum}`}>
              {pageNum}
            </PaginationLink>
          </PaginationItem>
          {pageNum * FETCH_QUIZZES_LIMIT < quizCounts && (
            <>
              <PaginationItem>
                <PaginationEllipsis />
              </PaginationItem>
              <PaginationItem>
                <PaginationNext href={`/?page=${pageNum + 1}`} />
              </PaginationItem>
            </>
          )}
        </PaginationContent>
      </Pagination>
    </main>
  );
}
