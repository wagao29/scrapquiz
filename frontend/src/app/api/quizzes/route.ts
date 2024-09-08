import { createQuiz } from "@/lib/apis";
import { auth } from "@/lib/auth";
import { NextResponse } from "next/server";

export const POST = auth(async function POST(request) {
  const authUserId = request.auth?.user?.id;
  if (!authUserId) {
    return NextResponse.json({ error: "user not logged in" }, { status: 400 });
  }

  const req = await request.json();
  const data = {
    userId: authUserId,
    ...req,
  };
  const response = await createQuiz(data);

  return NextResponse.json(response);
});
