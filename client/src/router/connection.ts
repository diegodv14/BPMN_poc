import { Client } from "pg"

export class DbFactory {
    private static instance: DbFactory;
    private client: Client;
    private isConnected: boolean = false;

    private constructor() {
        this.client = new Client({
            host: process.env.POSTGRES_HOST,
            port: process.env.POSTGRES_PORT ? Number(process.env.POSTGRES_PORT) : 5432,
            database: process.env.POSTGRES_DB,
            user: process.env.POSTGRES_USER,
            password: process.env.POSTGRES_PASS
        })
    }
     
    public static getInstance(): DbFactory {
        if (!DbFactory.instance) {
            DbFactory.instance = new DbFactory();
        }
        return DbFactory.instance;
    }

    async connect() {
        try {
            if (!this.isConnected) {
                await this.client.connect();
                this.isConnected = true;
                console.log('Conectado a la base de datos PostgreSQL');
            }
        } catch (err) {
            console.error('Error conectando a la base de datos:', err);
            throw err;
        }
    }

    async disconnect() {
        try {
            if (this.isConnected) {
                await this.client.end();
                this.isConnected = false;
                console.log('🔌 Desconectado de la base de datos');
            }
        } catch (err) {
            console.error('❌ Error desconectando de la base de datos:', err);
            throw err;
        }
    }

    getClient(): Client {
        if (!this.isConnected) {
            throw new Error('La conexión a la base de datos no está establecida. Llama a connect() primero.');
        }
        return this.client;
    }

    isConnectedToDb(): boolean {
        return this.isConnected;
    }
}
