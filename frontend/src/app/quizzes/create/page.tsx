import { QuizForm } from "@/components/quiz-form";

export default async function Page() {
  return (
    <main className="flex flex-col items-center">
      <h1 className="text-lg">クイズ作成</h1>
      <QuizForm />
    </main>
  );
}
