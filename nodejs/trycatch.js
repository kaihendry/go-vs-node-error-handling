var express = require("express");

var app = express();

app.use(express.json());

app.post("/", function (request, response) {
  // age in ten years
  try {
    older = request.body.human.age + 10;
    response.status(200).json({ older });
  } catch (error) {
    response.status(400).send(error.message);
  }
});

app.listen(3000);
