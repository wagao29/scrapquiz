import Link from "next/link";
import { UserDropdown } from "./user-dropdown";

export default function Header() {
  return (
    <header className="flex justify-between px-3 py-4">
      <Link href="/" className="text-xl font-semibold">
        scrapquiz
      </Link>
      <UserDropdown
        userId="01FVSHW3SER8977QCJBYZD9HAW"
        userName="太郎"
        avatarUrl="https://lh3.googleusercontent.com/a/ACg8ocIxf8eU7MG6Gpt6LLPgV_xmTzee-ZmBJ6UV5-UpmtbYtikaswes=s288-c-no"
      />
    </header>
  );
}
