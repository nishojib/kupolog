import { useRouter } from 'next/navigation';
import { useCallback, useEffect, useState } from 'react';

export const useSSEEvents = (url: string) => {
  const [sseConnection, setSSEConnection] = useState<EventSource | null>(null);
  const router = useRouter();

  const listenToSSEUpdates = useCallback(() => {
    const eventSource = new EventSource(url, {
      withCredentials: true,
    });

    eventSource.onopen = () => {
      console.log('sse connection opened');
    };

    eventSource.onmessage = (event) => {
      const data = event.data;
      console.log('message received', data);
      router.refresh();
    };

    eventSource.onerror = (error) => {
      console.error('Error:', error);
    };

    setSSEConnection(eventSource);
    return eventSource;
  }, [router, url]);

  useEffect(() => {
    listenToSSEUpdates();

    return () => {
      if (sseConnection) {
        sseConnection.close();
      }
    };
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [listenToSSEUpdates]);

  useEffect(() => {
    // Add beforeunload event listener to close SSE connection when navigating away
    const handleBeforeUnload = () => {
      console.dir(sseConnection);
      if (sseConnection) {
        console.info('closing sse connection before unloading the page');
        sseConnection.close();
      }
    };

    window.addEventListener('beforeunload', handleBeforeUnload);

    return () => {
      window.removeEventListener('beforeunload', handleBeforeUnload);
    };
  }, [sseConnection]);
};
