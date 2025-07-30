import { PrismaClient } from "@prisma/client";

const prisma = new PrismaClient();

export async function getAbout() {
  try {
    const response = await prisma.about.findMany({
      orderBy: { createdAt: "desc" },
    });
    return { response };
  } catch (e) {
    console.log(e);
  }
}
