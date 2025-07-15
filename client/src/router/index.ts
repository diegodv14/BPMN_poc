import { Request, Router, Response } from "express";
import { DbFactory } from "../db/connection";

const router = Router();
/**
 * @swagger
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
    const db = DbFactory.getInstance();
    const queueName = process.env.QUEUE
    await db.connect();
    const client = db.getClient();
    const query = `SELECT pgmq_send($1, $2);`;
    const resmq = await client.query(query, [queueName, JSON.stringify(data)]);
    console.log('Mensaje enviado con ID:', resmq.rows[0].pgmq_send);
  
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
