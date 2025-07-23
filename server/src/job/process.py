import json
from gql import gql
from src.models.request import Request
from src.db.connection import DbFactory
import os
from src.graphql.client import client as graphql_client

def start_async():
    db = DbFactory()
    db.connect()
    client = db.get_client()
    
    with client.cursor() as cur:
        cur.execute("SELECT msg_id, message FROM pgmq_read(%s, %s);", (os.getenv("QUEUE"), 10))
        mensajes = cur.fetchall()
        
        for mensaje in mensajes:
            print(mensaje)
            data = json.loads(mensaje[1])
            request = Request(**data)
            print(request)
            query = gql("""
                mutation($name: String!, $description: String!) {
                    insert_request(objects: {name: $name, description: $description}) {
                        returning {
                            id
                            name
                            description
                            status
                            created_at
                            updated_at
                        }
                    }
                }
            """)
            variables = {
                "id": request.id,
                "name": request.name,
                "description": request.description,
                "status": request.status,
                "created_at": request.created_at,
                "updated_at": request.updated_at
            }
            response = graphql_client.execute(query, variables)
            print(response)
            
            
