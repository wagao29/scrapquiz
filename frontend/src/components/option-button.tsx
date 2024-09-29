"use client";

import { cn } from "@/lib/utils";
import { CheckCircle2, CircleX } from "lucide-react";
import { Button } from "./ui/button";

type Props = React.ComponentProps<typeof Button> & {
  option: string;
  optionNum: number;
  answeredNum: number;
  answerCounts: number;
  answerCountsSum: number;
  isCorrect?: boolean;
};

export function OptionButton({
  option,
  optionNum,
  answeredNum,
  answerCounts,
  answerCountsSum,
  isCorrect,
  className,
  ...props
}: Props) {
  const answerRate = Math.floor(
    ((answerCounts + (answeredNum === optionNum ? 1 : 0)) /
      (answerCountsSum + 1)) *
      100
  );

  return (
    <Button
      variant="outline"
      className={cn(
        "h-auto w-full p-2 relative text-base",
        `${
          answeredNum === optionNum
            ? `${
                isCorrect
                  ? "border-green-700 text-green-700 hover:text-green-700"
                  : "border-red-500 text-red-500 hover:text-red-500"
              }`
            : "border-gray-400"
        }`,
        className
      )}
      {...props}
    >
      <div className="w-full break-words whitespace-pre-wrap text-start">
        {option}
      </div>
      <div
        className={cn(
          "absolute top-0 left-0 h-full bg-blue-400 opacity-20 rounded-s",
          answeredNum !== 0 && "transition-all duration-500",
          answerRate === 100 && "rounded-e"
        )}
        style={{ width: `${answeredNum === 0 ? 0 : `${answerRate}%`}` }}
      ></div>
      {answeredNum !== 0 && (
        <>
          <div className="absolute top-1/2 transform -translate-y-1/2 -left-8">
            {isCorrect ? (
              <CheckCircle2 color="green"></CheckCircle2>
            ) : (
              <CircleX color="red"></CircleX>
            )}
          </div>
          <span className="ml-1">{`${answerRate}%`}</span>
        </>
      )}
    </Button>
  );
}
