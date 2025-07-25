import express, { Application } from "express";
import swaggerUi from "swagger-ui-express";
import swaggerJsdoc from "swagger-jsdoc";
import cors from "cors";
import dotenv from "dotenv";
import { router } from "./router";
import bodyParser from "body-parser";

dotenv.config();

const app: Application = express();

app.use(cors());
app.use(express.json());
app.use(bodyParser.json()); 
app.use(express.urlencoded({ extended: true }));

const swaggerOptions = {
  definition: {
    openapi: "3.0.0",
    info: {
      title: "BPMN Client API",
      version: "1.0.0",
      description: "API del cliente BPMN con Express",
    },
    servers: [
      {
        url: "/api",
      },
    ],
  },
  apis: ["./src/router/*.ts"],
  
};

const swaggerSpec = swaggerJsdoc(swaggerOptions);

app.use("/api", router);

app.use("/", swaggerUi.serve, swaggerUi.setup(swaggerSpec));

app.use(
  (
    err: any,
    req: express.Request,
    res: express.Response,
    next: express.NextFunction
  ) => {
    console.error(err.stack);
    res.status(500).json({ error: "Algo salió mal!" });
  }
);

export default app;
