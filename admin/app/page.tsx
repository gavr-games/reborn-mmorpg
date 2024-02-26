import Image from "next/image";

export default function Home() {
  return (
    <main className="flex flex-col">
      <h1>Admin site</h1>
      <Image
        src="/admin/logo.jpg"
        width={760}
        height={760}
        className="hidden rounded-2xl md:block"
        alt="Reborn MMORPG Logo"
      />
    </main>
  );
}
