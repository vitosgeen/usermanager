UPDATE "users" SET "user_role" = 'user' WHERE "nickname" NOT IN ('root', 'admin');