import { createUser, fetchUser, updateUser } from "@/lib/apis";
import type { NextAuthConfig } from "next-auth";
import Google from "next-auth/providers/google";

export const authConfig: NextAuthConfig = {
  providers: [Google],
  callbacks: {
    async signIn({ user, profile }) {
      if (profile?.sub) {
        user.id = profile.sub;
        const fetchedUser = await fetchUser(user.id);
        if (fetchedUser) {
          if (
            fetchedUser.name !== user.name ||
            fetchedUser.avatarUrl !== user.image
          ) {
            await updateUser(user.id, user.name ?? "unknown", user.image ?? "");
          }
        } else {
          await createUser(user.id, user.name ?? "unknown", user.image ?? "");
        }
      }
      return true;
    },
    async session({ token, session }) {
      if (token.sub && session.user) {
        session.user.id = token.sub;
      }
      return session;
    },
  },
};
