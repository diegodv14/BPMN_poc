import psycopg
from dotenv import load_dotenv
import os
load_dotenv()
class DbFactory:
    _instance = None

    def __new__(cls):
        if cls._instance is None:
            cls._instance = super().__new__(cls)
            cls._instance.client = None
            cls._instance.is_connected = False
        return cls._instance

    def __init__(self):
        self.config = {
            'host': os.getenv('POSTGRES_HOST'),
            'port': int(os.getenv('POSTGRES_PORT', 5432)),
            'dbname': os.getenv('POSTGRES_DB'),
            'user': os.getenv('POSTGRES_USER'),
            'password': os.getenv('POSTGRES_PASSWORD'),
        }

    def connect(self):
        if not self.is_connected:
            self.client = psycopg.connect(**self.config)
            self.is_connected = True
            print("Conectado a la base de datos PostgreSQL")

    def disconnect(self):
        if self.is_connected and self.client:
            self.client.close()
            self.is_connected = False
            print("Desconectado de la base de datos")

    def get_client(self):
        if not self.is_connected:
            raise RuntimeError("La conexión a la base de datos no está establecida. Llama a connect() primero.")
        return self.client

    def is_connected_to_db(self):
        return self.is_connected