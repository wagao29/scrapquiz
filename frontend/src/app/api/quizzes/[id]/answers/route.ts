import { createAnswer } from "@/lib/apis";
import { auth } from "@/lib/auth";
import { NextResponse } from "next/server";

export const POST = auth(async function POST(
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
  const req = await request.json();
  const data = {
    userId: authUserId || req.anonymousUserId,
    answerNum: req.answerNum,
  };
  const response = await createAnswer(params.id, data);

  return NextResponse.json(response);
});
