"use client";

import { useState } from "react";
import { OptionButton } from "./option-button";

type Props = {
  quizId: string;
  content: string;
  options: string[];
  correctNum: number;
  explanation: string;
  answerCounts: number[];
  answerCountsSum: number;
};

export function QuizContent({
  quizId,
  content,
  explanation,
  options,
  correctNum,
  answerCounts,
  answerCountsSum,
}: Props) {
  const [answeredNum, setAnsweredNum] = useState(0);

  return (
    <div className="flex flex-col px-4">
      <div>{content}</div>
      <div className="mt-2.5">
        {options.map((opt, index) => {
          return (
            <OptionButton
              key={index}
              option={opt}
              optionNum={index + 1}
              isCorrect={index + 1 === correctNum}
              answerCounts={answerCounts[index]}
              answerCountsSum={answerCountsSum}
              answeredNum={answeredNum}
              className="my-1.5"
              onClick={async () => {
                if (answeredNum === 0) {
                  setAnsweredNum(index + 1);
                  const params = {
                    method: "POST",
                    headers: {
                      "Content-Type": "application/json",
                    },
                    body: JSON.stringify({
                      answerNum: index + 1,
                    }),
                  };
                  await fetch(`/api/quizzes/${quizId}/answers`, params);
                }
              }}
            />
          );
        })}
      </div>
      {answeredNum !== 0 && explanation && (
        <div className="bg-gray-100 p-2.5 mt-2.5">
          <p>解説</p>
          <p>{explanation}</p>
        </div>
      )}
    </div>
  );
}
