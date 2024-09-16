import "server-only";

import {
  API_KEY,
  ENDPOINT_URL,
  FETCH_ANSWER_COUNTS_REVALIDATION_SEC,
  FETCH_LATEST_QUIZZES_REVALIDATION_SEC,
  FETCH_QUIZ_COUNTS_REVALIDATION_SEC,
  FETCH_QUIZ_REVALIDATION_SEC,
  FETCH_QUIZZES_BY_USER_ID_REVALIDATION_SEC,
  FETCH_QUIZZES_LIMIT,
  FETCH_RANDOM_QUIZZES_REVALIDATION_SEC,
  FETCH_USER_REVALIDATION_SEC,
} from "./constants";
import {
  answerCountsSchema,
  quizSchema,
  quizzesSchema,
  userSchema,
} from "./schemas";
import { AnswerCounts, Quiz, Quizzes, User } from "./types";

export async function fetchUser(userId: string): Promise<User | undefined> {
  try {
    const params = {
      method: "GET",
      headers: {
        "x-api-key": API_KEY,
      },
      next: { revalidate: FETCH_USER_REVALIDATION_SEC },
    };
    const response = await fetch(`${ENDPOINT_URL}/users/${userId}`, params);
    if (!response.ok) {
      throw new Error(`[fetchUser] error status code: ${response.status}`);
    }
    const json = await response.json();
    return userSchema.parse(json);
  } catch (error) {
    console.error(error);
  }
}

export async function fetchQuiz(quizId: string): Promise<Quiz | undefined> {
  try {
    const params = {
      method: "GET",
      headers: {
        "x-api-key": API_KEY,
      },
      next: { revalidate: FETCH_QUIZ_REVALIDATION_SEC },
    };
    const response = await fetch(`${ENDPOINT_URL}/quizzes/${quizId}`, params);
    if (!response.ok) {
      throw new Error(`[fetchQuiz] error status code: ${response.status}`);
    }
    const json = await response.json();
    return quizSchema.parse(json);
  } catch (error) {
    console.error(error);
  }
}

export async function fetchQuizCounts(userId?: string): Promise<number> {
  try {
    const params = {
      method: "GET",
      headers: {
        "x-api-key": API_KEY,
      },
      next: { revalidate: FETCH_QUIZ_COUNTS_REVALIDATION_SEC },
    };
    const response = await fetch(
      `${ENDPOINT_URL}/quizzes/counts${userId ? `?user_id=${userId}` : ""}`,
      params
    );

    if (!response.ok) {
      throw new Error(
        `[fetchQuizCounts] error status code: ${response.status}`
      );
    }
    const json = await response.json();
    return json.quizCounts || 0;
  } catch (error) {
    console.error(error);
    return 0;
  }
}

export async function fetchLatestQuizzes(
  offset: number
): Promise<Quizzes | undefined> {
  try {
    const params = {
      method: "GET",
      headers: {
        "x-api-key": API_KEY,
      },
      next: { revalidate: FETCH_LATEST_QUIZZES_REVALIDATION_SEC },
    };
    const response = await fetch(
      `${ENDPOINT_URL}/quizzes?limit=${FETCH_QUIZZES_LIMIT}&offset=${offset}`,
      params
    );
    if (!response.ok) {
      throw new Error(
        `[fetchLatestQuizzes] error status code: ${response.status}`
      );
    }
    const json = await response.json();
    return quizzesSchema.parse(json);
  } catch (error) {
    console.error(error);
  }
}

export async function fetchRandomQuizzes(): Promise<Quizzes | undefined> {
  try {
    const params = {
      method: "GET",
      headers: {
        "x-api-key": API_KEY,
      },
      next: { revalidate: FETCH_RANDOM_QUIZZES_REVALIDATION_SEC },
    };
    const response = await fetch(
      `${ENDPOINT_URL}/quizzes?order=random&limit=${FETCH_QUIZZES_LIMIT}&offset=0`,
      params
    );
    if (!response.ok) {
      throw new Error(
        `[fetchRandomQuizzes] error status code: ${response.status}`
      );
    }
    const json = await response.json();
    return quizzesSchema.parse(json);
  } catch (error) {
    console.error(error);
  }
}

export async function fetchQuizzesByUserId(
  userId: string,
  offset: number
): Promise<Quizzes | undefined> {
  try {
    const params = {
      method: "GET",
      headers: {
        "x-api-key": API_KEY,
      },
      next: { revalidate: FETCH_QUIZZES_BY_USER_ID_REVALIDATION_SEC },
    };
    const response = await fetch(
      `${ENDPOINT_URL}/quizzes?user_id=${userId}&limit=${FETCH_QUIZZES_LIMIT}&offset=${offset}`,
      params
    );
    if (!response.ok) {
      throw new Error(
        `[fetchLatestQuizzes] error status code: ${response.status}`
      );
    }
    const json = await response.json();
    return quizzesSchema.parse(json);
  } catch (error) {
    console.error(error);
  }
}

export async function fetchAnswerCounts(
  quizId: string
): Promise<AnswerCounts | undefined> {
  try {
    const params = {
      method: "GET",
      headers: {
        "x-api-key": API_KEY,
      },
      next: { revalidate: FETCH_ANSWER_COUNTS_REVALIDATION_SEC },
    };
    const response = await fetch(
      `${ENDPOINT_URL}/quizzes/${quizId}/answer_counts`,
      params
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

export async function createUser(
  userId: string,
  userName: string,
  userImage: string
): Promise<User | undefined> {
  try {
    const params = {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "x-api-key": API_KEY,
      },
      body: JSON.stringify({
        id: userId,
        name: userName,
        avatarUrl: userImage,
      }),
    };
    const response = await fetch(`${ENDPOINT_URL}/users`, params);
    if (!response.ok) {
      throw new Error(`[createUser] error status code: ${response.status}`);
    }
    const json = await response.json();
    return userSchema.parse(json);
  } catch (error) {
    console.error(error);
  }
}

export async function updateUser(
  userId: string,
  userName: string,
  userImage: string
): Promise<boolean> {
  try {
    const params = {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
        "x-api-key": API_KEY,
      },
      body: JSON.stringify({
        name: userName,
        avatarUrl: userImage,
      }),
    };
    const response = await fetch(`${ENDPOINT_URL}/users/${userId}`, params);
    if (!response.ok) {
      throw new Error(`[updateUser] error status code: ${response.status}`);
    }
    return true;
  } catch (error) {
    console.error(error);
    return false;
  }
}

export async function createQuiz(data: any): Promise<any | undefined> {
  try {
    const params = {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "x-api-key": API_KEY,
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

export async function deleteQuiz(quizId: string): Promise<boolean> {
  try {
    const params = {
      method: "DELETE",
      headers: {
        "x-api-key": API_KEY,
      },
    };
    const response = await fetch(`${ENDPOINT_URL}/quizzes/${quizId}`, params);
    if (!response.ok) {
      throw new Error(`[deleteQuiz] error status code: ${response.status}`);
    }
    return true;
  } catch (error) {
    console.error(error);
    return false;
  }
}

export async function createAnswer(
  quizId: string,
  data: any
): Promise<any | undefined> {
  try {
    const params = {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "x-api-key": API_KEY,
      },
      body: JSON.stringify(data),
    };
    const response = await fetch(
      `${ENDPOINT_URL}/quizzes/${quizId}/answers`,
      params
    );
    if (!response.ok) {
      throw new Error(`[createAnswer] error status code: ${response.status}`);
    }
    return response.json();
  } catch (error) {
    console.error(error);
  }
}
