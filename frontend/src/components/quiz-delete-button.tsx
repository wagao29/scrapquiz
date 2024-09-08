"use client";

import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
} from "@/components/ui/alert-dialog";
import { Trash2 } from "lucide-react";
import { useRouter } from "next/navigation";
import { useState } from "react";
import { toastError, toastSuccess } from "./toasts";
import { Button } from "./ui/button";

type Props = React.ComponentProps<typeof Button> & {
  quizId: string;
};

export function QuizDeleteButton({ quizId, className, ...props }: Props) {
  const router = useRouter();
  const [openDialog, setOpenDialog] = useState(false);

  const onClickDelete = async () => {
    const response = await fetch(`/api/quizzes/${quizId}`, {
      method: "DELETE",
    });
    if (response.ok) {
      router.refresh();
      toastSuccess("クイズを削除しました");
    } else {
      toastError("クイズの削除に失敗しました");
    }
    setOpenDialog(false);
  };

  return (
    <>
      <Button
        size="icon"
        variant="ghost"
        onClick={() => {
          setOpenDialog(true);
        }}
        {...props}
      >
        <Trash2 color="red" />
      </Button>
      <AlertDialog open={openDialog}>
        <AlertDialogContent>
          <AlertDialogHeader>
            <AlertDialogTitle>
              本当にこのクイズを削除しますか？
            </AlertDialogTitle>
            <AlertDialogDescription>
              この操作は取り消せません。
            </AlertDialogDescription>
          </AlertDialogHeader>
          <AlertDialogFooter>
            <AlertDialogCancel
              onClick={() => {
                setOpenDialog(false);
              }}
            >
              キャンセル
            </AlertDialogCancel>
            <AlertDialogAction
              className="bg-white text-red-600 border border-red-600 hover:bg-red-600 hover:text-white"
              onClick={onClickDelete}
            >
              削除する
            </AlertDialogAction>
          </AlertDialogFooter>
        </AlertDialogContent>
      </AlertDialog>
    </>
  );
}
