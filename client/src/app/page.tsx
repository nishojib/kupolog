import axios from 'axios';
import { Suspense } from 'react';

export default async function Home() {
  const data = await axios.get(`${process.env.NEXT_PUBLIC_API_URL}/health`, {
    headers: {
      'Content-Type': `application/json`,
    },
    withCredentials: true,
  });

  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      <Suspense fallback={<p>Loading...</p>}>
        {JSON.stringify(data.data, null, 2)}
      </Suspense>
    </main>
  );
}
