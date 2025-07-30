import express from "express";
import { handleGetGallery } from "../controllers/galleryController";

const router = express.Router();

router.get("/", handleGetGallery);

export default router;
