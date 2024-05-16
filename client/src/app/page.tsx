'use client';

import { useQuery } from '@tanstack/react-query';
import axios from 'axios';

export default function Home() {
  const { data, error } = useQuery({
    queryKey: ['health'],
    queryFn: async () => {
      const response = await axios.get(
        `${process.env.NEXT_PUBLIC_API_URL}/health`,
        {
          headers: {
            'Content-Type': `application/json`,
          },
          withCredentials: true,
        }
      );

      console.log(response.data);
      return response.data;
    },
  });

  if (error) {
    return <div>Error: {error.message}</div>;
  }

  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      {JSON.stringify(data, null, 2)}
    </main>
  );
}
