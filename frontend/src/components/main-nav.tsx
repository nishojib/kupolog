'use client';

import Link from 'next/link';
import { usePathname } from 'next/navigation';
import * as React from 'react';

import { Icons } from '@/components/icons';
import { siteConfig } from '@/config/site';
import { cn } from '@/lib/utils';

export function MainNav() {
  const pathname = usePathname();

  return (
    <div className="mr-4 hidden md:flex">
      <Link href="/" className="mr-6 flex items-center space-x-2">
        <Icons.logo className="text-primary size-6 fill-current" />
        <span className="text-primary hidden font-bold sm:inline-block">
          {siteConfig.name}
        </span>
      </Link>
      <nav className="flex items-center gap-4 text-sm lg:gap-6">
        <Link
          href="/profile"
          className={cn(
            'hover:text-foreground/80 transition-colors',
            pathname === '/profile' ? 'text-foreground' : 'text-foreground/60',
          )}
        >
          Profile
        </Link>
        <Link
          href="/dailies"
          className={cn(
            'hover:text-foreground/80 transition-colors',
            pathname?.startsWith('/dailies')
              ? 'text-foreground'
              : 'text-foreground/60',
          )}
        >
          Dailies
        </Link>
      </nav>
    </div>
  );
}
