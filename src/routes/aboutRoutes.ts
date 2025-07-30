import express from "express";
import { handleGetAbout } from "../controllers/aboutController";

const router = express.Router();

router.get("/", handleGetAbout);

export default router;
