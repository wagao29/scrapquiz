import { fetchQuiz } from "@/lib/apis";
import { notFound } from "next/navigation";

type Props = {
  quizId: string;
};

export async function QuizCard({ quizId }: Props) {
  const quiz = await fetchQuiz(quizId);
  if (!quiz) notFound();

  return (
    <div>
      <p>id: {quiz.id}</p>
      <p>content: {quiz.content}</p>
      {quiz.options.map((opt, index) => {
        return (
          <p key={index}>
            option{index + 1}: {opt}
          </p>
        );
      })}
      <p>correctNum: {quiz.correctNum}</p>
      <p>explanation: {quiz.explanation}</p>
      <p>userId: {quiz.userId}</p>
      <p>userName: {quiz.userName}</p>
      <p>userAvatarUrl: {quiz.userAvatarUrl}</p>
    </div>
  );
}
