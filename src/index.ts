import express from "express";
import newsRoutes from "./routes/newsRoutes";
import curriculumRoutes from "./routes/curriculumRoutes";
import scheduleRoutes from "./routes/scheduleRoutes";

const app = express();
app.use(express.json());

// Routes
app.use("/news", newsRoutes);
app.use("/curriculums", curriculumRoutes);
app.use("/schedules", scheduleRoutes);

app.get("/", (req, res) => {
  res.send("welcome to the backend");
});

app.listen(3000, () => {
  console.log("Server is running on http://localhost:3000");
});
