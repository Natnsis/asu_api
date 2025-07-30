import { PrismaClient } from "@prisma/client";

const prisma = new PrismaClient();

const getAllNews = async () => {
  try {
    const response = await prisma.news.findMany({
      orderBy: {
        createdAt: "desc",
      },
    });
    return response;
  } catch (e) {
    console.log(e);
  }
};

export { getAllNews };
