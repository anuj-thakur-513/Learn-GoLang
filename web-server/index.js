const express = require("express");

const app = express();
const PORT = process.env.PORT || 8000;

app.use(express.json());
app.use(express.urlencoded({ extended: true }));

app.get("/", (req, res) => {
  res.status(200).send("Welcome to LearnCodeonline server");
});

app.get("/get", (req, res) => {
  res.status(200).json({ message: "Hello from learnCodeonline.in" });
});

app.post("/post", (req, res) => {
  let myJson = req.body; // your JSON

  res.status(200).send(myJson);
});

app.post("/postform", (req, res) => {
  res.status(200).send(JSON.stringify(req.body));
});

app.listen(PORT, () => {
  console.log(`server started on port ${PORT}`);
});
