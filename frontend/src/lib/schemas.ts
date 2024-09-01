import { z } from "zod";

export const quizSchema = z.object({
  id: z.string().ulid(),
  content: z.string(),
  options: z.array(z.string()).min(2).max(4),
  correctNum: z.number().min(1).max(4),
  explanation: z.string(),
  userId: z.string(),
  userName: z.string(),
  userAvatarUrl: z.string().url(),
});

export const quizzesSchema = z.object({
  quizzes: z.array(quizSchema),
});

export const answerCountsSchema = z.object({
  answerNum1Count: z.number().min(0),
  answerNum2Count: z.number().min(0),
  answerNum3Count: z.number().min(0),
  answerNum4Count: z.number().min(0),
});
