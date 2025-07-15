import express, { Application } from "express";
import dotenv from "dotenv"

dotenv.config()

const app: Application = express()
const port: string | undefined = process.env.PORT

app.get("/", (request, response) => { 
    response.status(200).send("Hello World");
  }); 

app.listen(port, () => {
    console.log("Listening on Port " + port)
}).on("error", (error) => {
    throw new Error(error.message);
  })


