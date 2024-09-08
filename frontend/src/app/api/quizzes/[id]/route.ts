import { deleteQuiz, fetchQuiz } from "@/lib/apis";
import { auth } from "@/lib/auth";
import { NextResponse } from "next/server";

export const DELETE = auth(async function DELETE(
  request,
  { params }: { params?: { id?: string } }
) {
  if (!params?.id) {
    return NextResponse.json(
      { error: "quizId is not included in the request path" },
      { status: 400 }
    );
  }

  const authUserId = request.auth?.user?.id;
  if (!authUserId) {
    return NextResponse.json({ error: "user not logged in" }, { status: 400 });
  }

  const quiz = await fetchQuiz(params.id);
  if (quiz?.userId !== authUserId) {
    return NextResponse.json(
      { error: "login user and quiz author are not the same" },
      { status: 400 }
    );
  }

  const response = await deleteQuiz(params.id);
  if (response) {
    return NextResponse.json({
      message: "quiz successfully deleted",
    });
  } else {
    return NextResponse.json(
      { error: "failed to delete quiz" },
      { status: 500 }
    );
  }
});
