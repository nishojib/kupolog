import { Suspense } from 'react';

import { getHealth } from './lib/actions';

export default async function Home() {
  const health = await getHealth();

  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      Hello, Kupo!
      <Suspense fallback={<p>Loading...</p>}>
        API Status: 🟢 {health.status}
      </Suspense>
    </main>
  );
}
2;
