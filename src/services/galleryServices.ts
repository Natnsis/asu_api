import { PrismaClient } from "@prisma/client";

const prisma = new PrismaClient();

export async function getGallery() {
  try {
    const response = await prisma.gallery.findMany({
      orderBy: { createdAt: "desc" },
    });
    return { response };
  } catch (e) {
    console.log(e);
  }
}
