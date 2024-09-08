import {
  Pagination,
  PaginationContent,
  PaginationEllipsis,
  PaginationItem,
  PaginationLink,
  PaginationNext,
  PaginationPrevious,
} from "@/components/ui/pagination";
import { FETCH_QUIZZES_LIMIT } from "@/lib/constants";

type Props = React.ComponentProps<typeof Pagination> & {
  basePath: string;
  pageNum: number;
  quizCounts: number;
};

export default function QuizPagination({
  basePath,
  pageNum,
  quizCounts,
  ...props
}: Props) {
  return (
    <Pagination {...props}>
      <PaginationContent>
        {pageNum > 1 && (
          <>
            <PaginationItem>
              <PaginationPrevious href={`${basePath}?page=${pageNum - 1}`} />
            </PaginationItem>
            <PaginationItem>
              <PaginationEllipsis />
            </PaginationItem>
          </>
        )}
        <PaginationItem>
          <PaginationLink href={`${basePath}?page=${pageNum}`}>
            {pageNum}
          </PaginationLink>
        </PaginationItem>
        {pageNum * FETCH_QUIZZES_LIMIT < quizCounts && (
          <>
            <PaginationItem>
              <PaginationEllipsis />
            </PaginationItem>
            <PaginationItem>
              <PaginationNext href={`${basePath}?page=${pageNum + 1}`} />
            </PaginationItem>
          </>
        )}
      </PaginationContent>
    </Pagination>
  );
}
