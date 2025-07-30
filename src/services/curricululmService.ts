import { PrismaClient } from "@prisma/client";

const prisma = new PrismaClient();

export async function getAllCurriculum() {
  try {
    const response = await prisma.curriculum.findMany({
      orderBy: { createdAt: "desc" },
    });
    return { response };
  } catch (e) {
    console.log(e);
  }
}
