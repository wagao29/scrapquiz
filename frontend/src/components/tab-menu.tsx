import { Tabs, TabsList, TabsTrigger } from "@/components/ui/tabs";
import Link from "next/link";

type Props = React.ComponentProps<typeof Tabs> & {
  selectValue: "latest" | "random";
};

export default function TabMenu({ selectValue, ...props }: Props) {
  return (
    <Tabs defaultValue={selectValue} {...props}>
      <TabsList>
        <TabsTrigger value="latest" asChild>
          <Link href={"/"}>新着</Link>
        </TabsTrigger>
        <TabsTrigger value="random" asChild>
          <Link href={"/random"}>ランダム</Link>
        </TabsTrigger>
      </TabsList>
    </Tabs>
  );
}
