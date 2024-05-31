import { Suspense } from 'react';

import { getHealth } from './lib/actions';

export default async function Home() {
  const health = await getHealth();

  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      <div>Hello, Kupo!</div>
      <Suspense fallback={<p>Loading...</p>}>
        <div>API Status: 🟢 {health.status}</div>
      </Suspense>
    </main>
  );
}
