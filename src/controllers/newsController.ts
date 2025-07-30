import { getAllNews } from "../services/newsService";
import {Response, Request} from "express";

const handleGetAllNews = async (req: Request, res: Response) => {
  try {
    const news = await getAllNews();
    res.json(news);
  } catch (err) {
    console.log(err);
  }
};

export { handleGetAllNews };
