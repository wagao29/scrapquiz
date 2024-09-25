import { z } from "zod";

export const userSchema = z.object({
  id: z.string(),
  name: z.string(),
  avatarUrl: z.string().url(),
});

export const quizSchema = z.object({
  id: z.string().ulid(),
  content: z.string(),
  options: z.array(z.string()).min(2).max(4),
  correctNum: z.number().min(1).max(4),
  explanation: z.string(),
  userId: z.string(),
  userName: z.string(),
  userAvatarUrl: z.string().url(),
  createdAt: z.string().datetime(),
});

export const quizzesSchema = z.object({
  quizzes: z.array(quizSchema),
});

export const answerCountsSchema = z.object({
  answerCounts: z.array(z.number()).length(4),
});
