import { fetchAnswerCounts } from "@/lib/apis";
import { Quiz } from "@/lib/types";
import { cn, formatDate } from "@/lib/utils";
import { QuizContent } from "./quiz-content";
import { QuizDeleteButton } from "./quiz-delete-button";
import { QuizShareButton } from "./quiz-share-button";
import { Card, CardContent, CardFooter, CardHeader } from "./ui/card";
import { UserButton } from "./user-button";

type Props = React.ComponentProps<typeof Card> & {
  quiz: Quiz;
  deletable?: boolean;
};

export async function QuizCard({
  quiz,
  deletable,
  className,
  ...props
}: Props) {
  const res = await fetchAnswerCounts(quiz.id);
  const answerCounts = Object.values(res || {});
  const answerCountsSum = answerCounts.reduce((sum, ac) => sum + ac, 0);
  const correctAnswerRate = Math.floor(
    (answerCounts[quiz.correctNum - 1] / answerCountsSum) * 100
  );

  return (
    <Card
      className={cn("min-w-[350px] max-w-[850px] w-full", className)}
      {...props}
    >
      <CardHeader className="flex py-2 px-0">
        <div className="flex justify-between">
          <UserButton
            id={quiz.userId}
            name={quiz.userName}
            avatarUrl={quiz.userAvatarUrl}
          />
          <span className="text-sm mt-2 mr-3">
            正答率 {correctAnswerRate || 0}%
          </span>
        </div>
        <div className="flex gap-1.5 ml-4 text-gray-600">
          <span className="text-xs !mt-0">{formatDate(quiz.createdAt)}</span>
          <span className="text-xs !mt-0">{answerCountsSum}人が回答</span>
        </div>
      </CardHeader>
      <CardContent className="py-2">
        <QuizContent
          quizId={quiz.id}
          content={quiz.content}
          options={quiz.options}
          correctNum={quiz.correctNum}
          answerCounts={answerCounts}
          answerCountsSum={answerCountsSum}
          explanation={quiz.explanation}
        ></QuizContent>
      </CardContent>
      <CardFooter className="flex gap-0.5 justify-end p-2">
        {deletable && <QuizDeleteButton quizId={quiz.id} />}
        <QuizShareButton quizId={quiz.id} />
      </CardFooter>
    </Card>
  );
}
