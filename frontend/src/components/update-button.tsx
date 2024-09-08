"use client";

import { useRouter } from "next/navigation";
import { Button } from "./ui/button";

type Props = React.ComponentProps<typeof Button>;

export default function UpdateButton({ className, ...props }: Props) {
  const router = useRouter();

  return (
    <Button
      onClick={() => {
        router.refresh();
        window.scrollTo(0, 0);
      }}
      className={className}
      {...props}
    >
      更新する
    </Button>
  );
}
