import { Request, Response } from "express";
import { getSchedule } from "../services/scheduleServices";

export async function handleGetSchedule(req: Request, res: Response) {
  try {
    const schedule = await getSchedule();
    res.json(schedule);
  } catch (e) {
    console.log(e);
  }
}
