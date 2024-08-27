import { ENDPOINT_URL, FETCH_QUIZ_REVALIDATION_SEC } from "./constants";
import { quizSchema } from "./schemas";
import { Quiz } from "./types";

export async function fetchQuiz(quizId: string): Promise<Quiz | undefined> {
  try {
    const response = await fetch(`${ENDPOINT_URL}/quizzes/${quizId}`, {
      next: { revalidate: FETCH_QUIZ_REVALIDATION_SEC },
    });
    if (!response.ok) {
      throw new Error(`[fetchQuiz] error status code: ${response.status}`);
    }
    const json = await response.json();
    return quizSchema.parse(json);
  } catch (error) {
    console.error(error);
  }
}
