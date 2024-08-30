import { z } from "zod";
import { answerCountsSchema, quizSchema } from "./schemas";

export type Quiz = z.infer<typeof quizSchema>;
export type AnswerCounts = z.infer<typeof answerCountsSchema>;
