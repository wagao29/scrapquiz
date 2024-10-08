import { z } from "zod";
import {
  answerCountsSchema,
  quizSchema,
  quizzesSchema,
  userSchema,
} from "./schemas";

export type User = z.infer<typeof userSchema>;
export type Quiz = z.infer<typeof quizSchema>;
export type Quizzes = z.infer<typeof quizzesSchema>;
export type AnswerCounts = z.infer<typeof answerCountsSchema>;
