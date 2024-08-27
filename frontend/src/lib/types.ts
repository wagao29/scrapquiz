import { z } from "zod";
import { quizSchema } from "./schemas";

export type Quiz = z.infer<typeof quizSchema>;
