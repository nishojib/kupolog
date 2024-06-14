import Image from 'next/image';
import Link from 'next/link';

import { signIn } from '@/auth';
import { Icons } from '@/components/icons';
import { Button } from '@/components/ui/button';
import { siteConfig } from '@/config/site';

export default function Page() {
  return (
    <div className="h-screen w-full lg:grid lg:grid-cols-2">
      <div>
        <Link href="/" className="flex items-center gap-2 p-4">
          <Icons.logo className="text-primary size-6 fill-current" />
          <span className="text-primary hidden font-bold sm:inline-block">
            {siteConfig.name}
          </span>
        </Link>
        <div className="flex h-[calc(100vh-8rem)] flex-col items-center justify-center">
          <div className="mx-auto grid w-[350px] gap-6">
            <div className="grid gap-2 text-center">
              <h3 className="text-3xl font-bold">Login</h3>
              <p className="text-muted-foreground text-balance">
                Try logging in with your Google or Discord account
              </p>
            </div>
            <div className="grid gap-4">
              <form
                action={async () => {
                  'use server';
                  await signIn('google', { redirectTo: '/' });
                }}
              >
                <Button className="bg-google hover:bg-google/90 w-full font-semibold">
                  <Icons.google className="mr-2 size-4" />
                  Login with Google
                </Button>
              </form>
              <form
                action={async () => {
                  'use server';
                  await signIn('discord', { redirectTo: '/' });
                }}
              >
                <Button className="bg-discord hover:bg-discord/90 w-full font-semibold">
                  <Icons.discord className="mr-2 size-4" />
                  Login with Discord
                </Button>
              </form>
            </div>
          </div>
        </div>
      </div>
      <div className="bg-muted relative hidden lg:block">
        <Image
          src="/kupoking.png"
          alt="Image"
          width="1920"
          height="1080"
          className="size-full object-cover dark:brightness-[0.8]"
        />
      </div>
      <div className="p-4; text-primary-foreground absolute bottom-0 right-0 p-4 text-sm">
        Fanart by{' '}
        <a href="https://www.reddit.com/r/ffxiv/comments/112t0ru/good_king_moggle_mog_by_grehmerl_west/">
          Wtakoh
        </a>
      </div>
    </div>
  );
}
