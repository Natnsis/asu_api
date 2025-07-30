import express from "express";
import { handleGetAllNews } from "../controllers/newsController";
const router = express.Router();

router.get("/", handleGetAllNews);

export default router;
