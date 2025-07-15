from src.db.connection import DbFactory
import os

def start_async():
    db = DbFactory()
    db.connect()
    client = db.get_client()
    
    with client.cursor() as cur:
        cur.execute("SELECT msg_id, message FROM pgmq_read(%s, %s);", (os.getenv("QUEUE"), 10))
        mensajes = cur.fetchall()
        
        for mensaje in mensajes:
            print(mensaje)
