"use client";

import { ShareIcon } from "lucide-react";
import { Button } from "./ui/button";

type Props = React.ComponentProps<typeof Button> & {
  quizId: string;
};

export function QuizShareButton({ quizId, className, ...props }: Props) {
  const onClickButton = async () => {
    try {
      await window.navigator.share({
        title: "scrapquiz にアクセスしてクイズに答えよう！",
        url: `${window.location.origin}/quizzes/${quizId}`,
      });
    } catch {
      console.warn("window.navigator.share is not supported");
    }
  };

  return (
    <Button size="icon" variant="ghost" onClick={onClickButton} {...props}>
      <ShareIcon color="black"></ShareIcon>
    </Button>
  );
}
