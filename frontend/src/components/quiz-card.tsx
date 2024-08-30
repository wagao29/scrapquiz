import { fetchAnswerCounts } from "@/lib/apis";
import { Quiz } from "@/lib/types";
import { cn } from "@/lib/utils";
import { ShareIcon, Trash2 } from "lucide-react";
import { QuizContent } from "./quiz-content";
import { Button } from "./ui/button";
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

  return (
    <Card
      className={cn("min-w-[350px] max-w-[850px] w-full", className)}
      {...props}
    >
      <CardHeader className="flex py-2 px-0 items-start">
        <UserButton
          id={quiz.userId}
          name={quiz.userName}
          avatarUrl={quiz.userAvatarUrl}
        />
        <span className="text-xs ml-10 !mt-0">{answerCountsSum}人が回答</span>
      </CardHeader>
      <CardContent className="py-2">
        <QuizContent
          content={quiz.content}
          options={quiz.options}
          correctNum={quiz.correctNum}
          answerCounts={answerCounts}
          answerCountsSum={answerCountsSum}
          explanation={quiz.explanation}
        ></QuizContent>
      </CardContent>
      <CardFooter className="flex gap-0.5 justify-end p-2">
        {deletable && (
          <Button size="icon" variant="ghost">
            <Trash2 color="red"></Trash2>
          </Button>
        )}
        <Button size="icon" variant="ghost">
          <ShareIcon color="black"></ShareIcon>
        </Button>
      </CardFooter>
    </Card>
  );
}
