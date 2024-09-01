import "server-only";

import {
  ENDPOINT_URL,
  FETCH_ANSWER_COUNTS_REVALIDATION_SEC,
  FETCH_QUIZ_REVALIDATION_SEC,
} from "./constants";
import { answerCountsSchema, quizSchema } from "./schemas";
import { AnswerCounts, Quiz } from "./types";

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

export async function fetchAnswerCounts(
  quizId: string
): Promise<AnswerCounts | undefined> {
  try {
    const response = await fetch(
      `${ENDPOINT_URL}/quizzes/${quizId}/answer_counts`,
      {
        next: { revalidate: FETCH_ANSWER_COUNTS_REVALIDATION_SEC },
      }
    );
    if (!response.ok) {
      throw new Error(
        `[fetchAnswerCounts] error status code: ${response.status}`
      );
    }
    const json = await response.json();
    return answerCountsSchema.parse(json);
  } catch (error) {
    console.error(error);
  }
}

export async function createQuiz(data: any): Promise<any | undefined> {
  try {
    const params = {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    };
    const response = await fetch(`${ENDPOINT_URL}/quizzes`, params);
    if (!response.ok) {
      throw new Error(`[createQuiz] error status code: ${response.status}`);
    }
    return response.json();
  } catch (error) {
    console.error(error);
  }
}
