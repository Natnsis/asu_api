import { Request, Response } from "express";
import { getGallery } from "../services/galleryServices";

export async function handleGetGallery(req: Request, res: Response) {
  try {
    const gallery = await getGallery();
    res.json(gallery);
  } catch (e) {
    console.log(e);
  }
}
