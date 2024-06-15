import { getHealth } from '@/actions/health';
import { auth } from '@/auth';

export default async function Page() {
  const health = await getHealth();
  const session = await auth();

  return (
    <div className="container">
      <div className="mt-6 flex flex-col items-center justify-center">
        <h1 className="text-primary text-4xl font-bold">
          Hello, {session?.user?.name ?? 'Kupo'}!
        </h1>
        <p className="mt-2 flex items-center gap-1 capitalize">
          <span>status:</span>
          <span className="text-xs">🟢</span>
          <span>{health.status}</span>
        </p>
      </div>
    </div>
  );
}
