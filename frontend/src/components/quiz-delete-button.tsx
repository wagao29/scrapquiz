"use client";

import { Trash2 } from "lucide-react";

import { useRouter } from "next/navigation";
import { toastError, toastSuccess } from "./toasts";
import { Button } from "./ui/button";

type Props = React.ComponentProps<typeof Button> & {
  quizId: string;
};

export function QuizDeleteButton({ quizId, className, ...props }: Props) {
  const router = useRouter();
  return (
    <Button
      size="icon"
      variant="ghost"
      onClick={async () => {
        const response = await fetch(`/api/quizzes/${quizId}`, {
          method: "DELETE",
        });
        if (response.ok) {
          router.refresh();
          toastSuccess("クイズを削除しました");
        } else {
          toastError("クイズの削除に失敗しました");
        }
      }}
      {...props}
    >
      <Trash2 color="red" />
    </Button>
  );
}
