import { PrismaClient } from "@prisma/client";

const prisma = new PrismaClient();

export async function getSchedule() {
  try {
    const response = await prisma.schedule.findMany({
      orderBy: { createdAt: "desc" },
    });
    return { response };
  } catch (e) {
    console.log(e);
  }
}
