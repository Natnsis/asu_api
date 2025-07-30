import express from "express";
import { handleGetSchedule } from "../controllers/scheduleController";
const router = express.Router();

router.get("/", handleGetSchedule);

export default router;
