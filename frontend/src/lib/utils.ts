import { type ClassValue, clsx } from "clsx";
import { twMerge } from "tailwind-merge";

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

export function formatDate(datetime: string) {
  const from = new Date(datetime);
  const diff = new Date().getTime() - from.getTime();
  const elapsed = new Date(diff);

  if (elapsed.getUTCFullYear() - 1970) {
    return elapsed.getUTCFullYear() - 1970 + "年前";
  } else if (elapsed.getUTCMonth()) {
    return elapsed.getUTCMonth() + "ヶ月前";
  } else if (elapsed.getUTCDate() - 1) {
    return elapsed.getUTCDate() - 1 + "日前";
  } else if (elapsed.getUTCHours()) {
    return elapsed.getUTCHours() + "時間前";
  } else if (elapsed.getUTCMinutes()) {
    return elapsed.getUTCMinutes() + "分前";
  } else {
    return "たった今";
  }
}

export const getAnonymousUserId = () => {
  const anonymousUserId = localStorage.getItem("anonymousUserId");
  if (anonymousUserId) return anonymousUserId;

  const digits = "0123456789";
  let result = "a";
  for (let i = 0; i < 20; i++) {
    result += digits.charAt(Math.floor(Math.random() * digits.length));
  }
  localStorage.setItem("anonymousUserId", result);
  return result;
};
