"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { z } from "zod";

import { Button } from "@/components/ui/button";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import {
  MAX_QUIZ_CONTENT,
  MAX_QUIZ_EXPLANATION,
  MAX_QUIZ_OPTION,
} from "@/lib/constants";
import { useRouter } from "next/navigation";
import { Input } from "./ui/input";
import { RadioGroup, RadioGroupItem } from "./ui/radio-group";
import { Textarea } from "./ui/textarea";

const formSchema = z.object({
  content: z
    .string()
    .trim()
    .min(1, { message: "入力してください" })
    .max(MAX_QUIZ_CONTENT, {
      message: `${MAX_QUIZ_CONTENT}字以内で入力してください`,
    }),
  correctNum: z.enum(["1", "2", "3", "4"]),
  option1: z
    .string()
    .trim()
    .min(1, { message: "入力してください" })
    .max(MAX_QUIZ_OPTION, {
      message: `${MAX_QUIZ_OPTION}字以内で入力してください`,
    }),
  option2: z
    .string()
    .trim()
    .min(1, { message: "入力してください" })
    .max(MAX_QUIZ_OPTION, {
      message: `${MAX_QUIZ_OPTION}字以内で入力してください`,
    }),
  option3: z
    .string()
    .trim()
    .min(1, { message: "入力してください" })
    .max(MAX_QUIZ_OPTION, {
      message: `${MAX_QUIZ_OPTION}字以内で入力してください`,
    }),
  option4: z
    .string()
    .trim()
    .min(1, { message: "入力してください" })
    .max(MAX_QUIZ_OPTION, {
      message: `${MAX_QUIZ_OPTION}字以内で入力してください`,
    }),
  explanation: z
    .string()
    .trim()
    .max(MAX_QUIZ_EXPLANATION, {
      message: `${MAX_QUIZ_EXPLANATION}字以内で入力してください`,
    }),
});

export const QuizForm = () => {
  const router = useRouter();

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      content: "",
      correctNum: "1",
      option1: "",
      option2: "",
      option3: "",
      option4: "",
      explanation: "",
    },
  });

  async function onSubmit(values: z.infer<typeof formSchema>) {
    const options: string[] = [];
    for (const [key, value] of Object.entries(values)) {
      if (key.startsWith("option")) {
        options.push(value);
      }
    }
    const data = {
      content: values.content,
      correctNum: Number(values.correctNum),
      options: options,
      explanation: values.explanation,
    };
    const params = {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    };
    const response = await fetch("/api/quizzes/create", params);
    const json = await response.json();
    router.push(`/quizzes/${json.id}`);
  }

  return (
    <Form {...form}>
      <form className="w-full px-4 py-8 min-w-[350px] max-w-[850px] flex flex-col items-center">
        <FormField
          control={form.control}
          name="content"
          render={({ field }) => (
            <FormItem className="w-full">
              <FormLabel className="text-black">問題文</FormLabel>
              <FormControl>
                <Textarea
                  placeholder="ここに問題文を入力してください"
                  className="resize-y h-[200px]"
                  {...field}
                />
              </FormControl>
              <FormMessage className="absolute" />
            </FormItem>
          )}
        />
        <div className="relative w-full mt-8">
          <FormField
            control={form.control}
            name="correctNum"
            render={({ field }) => (
              <FormItem>
                <FormLabel>正解 / 選択肢</FormLabel>
                <FormControl>
                  <RadioGroup
                    onValueChange={field.onChange}
                    defaultValue={field.value}
                    className="flex flex-col space-y-10"
                  >
                    <FormItem className="flex items-center mt-3">
                      <FormControl>
                        <RadioGroupItem value="1" />
                      </FormControl>
                    </FormItem>
                    <FormItem className="flex items-center">
                      <FormControl>
                        <RadioGroupItem value="2" />
                      </FormControl>
                    </FormItem>
                    <FormItem className="flex items-center">
                      <FormControl>
                        <RadioGroupItem value="3" />
                      </FormControl>
                    </FormItem>
                    <FormItem className="flex items-center">
                      <FormControl>
                        <RadioGroupItem value="4" />
                      </FormControl>
                    </FormItem>
                  </RadioGroup>
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <div className="absolute top-8 left-8 space-y-6 w-[70%] min-w-[310px]">
            <FormField
              control={form.control}
              name="option1"
              render={({ field }) => (
                <FormItem>
                  <FormControl>
                    <FormItem>
                      <FormControl>
                        <Input placeholder="選択肢1" {...field} />
                      </FormControl>
                    </FormItem>
                  </FormControl>
                  <FormMessage className="absolute !m-0" />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="option2"
              render={({ field }) => (
                <FormItem>
                  <FormControl>
                    <FormItem>
                      <FormControl>
                        <Input placeholder="選択肢2" {...field} />
                      </FormControl>
                    </FormItem>
                  </FormControl>
                  <FormMessage className="absolute !m-0" />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="option3"
              render={({ field }) => (
                <FormItem>
                  <FormControl>
                    <FormItem>
                      <FormControl>
                        <Input placeholder="選択肢3" {...field} />
                      </FormControl>
                    </FormItem>
                  </FormControl>
                  <FormMessage className="absolute !m-0" />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="option4"
              render={({ field }) => (
                <FormItem>
                  <FormControl>
                    <FormItem>
                      <FormControl>
                        <Input placeholder="選択肢4" {...field} />
                      </FormControl>
                    </FormItem>
                  </FormControl>
                  <FormMessage className="absolute !m-0" />
                </FormItem>
              )}
            />
          </div>
        </div>
        <FormField
          control={form.control}
          name="explanation"
          render={({ field }) => (
            <FormItem className="w-full mt-12">
              <FormLabel className="text-black">解説文 (省略可)</FormLabel>
              <FormControl>
                <Textarea
                  placeholder="ここに解説文を入力してください"
                  className="resize-y"
                  {...field}
                />
              </FormControl>
              <FormMessage className="absolute" />
            </FormItem>
          )}
        />
        <Button
          type="button"
          onClick={form.handleSubmit(onSubmit)}
          className="mt-10"
        >
          作成する
        </Button>
      </form>
    </Form>
  );
};
