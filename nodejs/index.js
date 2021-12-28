var express = require("express");

var app = express();

app.use(express.json());

app.post("/", function (request, response) {
  older = request.body.human.age + 10;
  response.status(200).json({ older });
});

app.listen(3000);
