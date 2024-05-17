import './globals.css';

import type { Metadata } from 'next';
import { Inter } from 'next/font/google';

import { ReactQueryProvider } from '@/providers/react-query-provider';

const inter = Inter({ subsets: ['latin'] });

export const metadata: Metadata = {
  title: 'FFXIV Dailies',
  description: 'A simple app to track your daily FFXIV tasks.',
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className={inter.className}>
        <ReactQueryProvider>
          <div>{children}</div>
        </ReactQueryProvider>
      </body>
    </html>
  );
}
