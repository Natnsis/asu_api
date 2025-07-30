import express from "express";
import { handleGetCurriculum } from "../controllers/curriculumController";

const router = express.Router();

router.get("/", handleGetCurriculum);

export default router;
