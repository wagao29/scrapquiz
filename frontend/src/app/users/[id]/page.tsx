export default async function Page({ params }: { params: { id: string } }) {
  const id = params.id;

  return (
    <main className="flex flex-col gap-4 items-center mt-4">{`ユーザー詳細画面 (id: ${id})`}</main>
  );
}
