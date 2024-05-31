import { Suspense } from 'react';

import { getHealth } from './lib/actions';

export default async function Home() {
  const health = await getHealth();

  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      <Suspense fallback={<p>Loading...</p>}>status: {health.status} </Suspense>
    </main>
  );
}
2;
