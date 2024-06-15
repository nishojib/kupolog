import Link from 'next/link';

import { auth, signOut } from '@/auth';
import { MainNav } from '@/components/main-nav';
import { MobileNav } from '@/components/mobile-nav';
import { ModeToggle } from '@/components/mode-toggle';
import { Button, buttonVariants } from '@/components/ui/button';

export async function SiteHeader() {
  const session = await auth();

  return (
    <header className="border-border/40 bg-background/95 supports-[backdrop-filter]:bg-background/60 sticky top-0 z-50 w-full border-b backdrop-blur">
      <div className="container flex h-14 max-w-screen-2xl items-center">
        <MainNav />
        <MobileNav />
        <nav className="flex flex-1 items-center justify-end">
          {session?.user ? (
            <form
              action={async () => {
                'use server';
                await signOut({ redirectTo: '/' });
              }}
            >
              <Button variant="ghost">Sign Out</Button>
            </form>
          ) : (
            <Link
              href="/auth/login"
              className={buttonVariants({ variant: 'ghost' })}
            >
              Login
            </Link>
          )}
          <ModeToggle />
        </nav>
      </div>
    </header>
  );
}
