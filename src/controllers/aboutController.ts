import { Request, Response } from "express";
import { getAbout } from "../services/aboutServices";

export async function handleGetAbout(req: Request, res: Response) {
  try {
    const about = await getAbout();
    res.json(about);
  } catch (e) {
    console.log(e);
  }
}
