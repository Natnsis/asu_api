import { getAllCurriculum } from "../services/curricululmService";
import { Request, Response } from "express";
export async function handleGetCurriculum(req: Request, res: Response) {
  try {
    const curriculum = await getAllCurriculum();
    res.json(curriculum);
  } catch (e) {
    console.log(e);
  }
}
