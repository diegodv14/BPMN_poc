import { Request, Router, Response } from "express";
import { DbFactory } from "../db/connection";

const router = Router();
/**
 * @openapi
 * /create:
 *   post:
 *     summary: Create a new request that has to ve approved on the BPMD flow, It would be sent on a pg queue
 *     requestBody:
 *       required: true
 *       content:
 *         application/json:
 *           schema:
 *             type: object
 *             required:
 *               - name
 *               - email
 *               - description
 *             properties:
 *               name:
 *                 type: string
 *                 description: The requester's name
 *               description: 
 *                  type: string
 *                  description: The request description
 *     responses:
 *       201:
 *         description: Request created and sent
 *       500:
 *          description: An error ocurred on the server
 */
router.post("/create", async (req: Request, res: Response) => {
  try {
    const data = req.body;
    console.log("Data " +  JSON.stringify(data))
    const db = DbFactory.getInstance();
    const queueName = process.env.QUEUE
    await db.connect();
    const client = db.getClient();
    const query = `SELECT * from pgmq.send($1, $2);`;
    await client.query(query, [queueName, JSON.stringify(data)]);      
    res.status(201).send("Request created and sent")
  }
  catch (err: unknown) {
    if (err instanceof Error) {
      console.error("There was an error sending to the pg queue: " + err.message);
    } else {
      console.error("An unknown error occurred:", err);
    }
    res.status(500).send("An error ocurred on the server")
  }
});

export { router };
