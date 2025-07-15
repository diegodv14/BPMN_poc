import app from './src/app';
import dotenv from 'dotenv';

dotenv.config();

const port: string | number = process.env.PORT || 3001;

app.listen(port, async () => {
  console.log(`Client listening on http://localhost:${port}`);
}).on("error", (error) => {
  console.error('Error al iniciar el servidor:', error.message);
  process.exit(1);
});


