import asyncio
from gql import gql
from src.models.request import Request
import os
from src.graphql.client import client as graphql_client
import aiohttp

class JobProcess:
    def __init__(self, client):
        self.semaphore = asyncio.Semaphore(1)
        self.client = client
    
    def start(self):
        if self.semaphore.locked():
            pass
        else:
            asyncio.run(self.start_async())
    
    async def start_async(self):
        with self.client.cursor() as cur:
            cur.execute(
            "SELECT msg_id, message FROM pgmq.read(%s::TEXT, %s::INTEGER, %s::INTEGER);",
            (os.getenv("QUEUE"), 30, 2)
            )
            mensajes = cur.fetchall()
            
            for mensaje in mensajes:
                print(mensaje)
                data = mensaje[1]
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
                    "name": request.name,
                    "description": request.description
                }
                response = graphql_client.execute(query, variables)
                if res := response.get("insert_request"):
                    cur.execute(
                        "SELECT pgmq.archive(%s::TEXT, %s::INTEGER);",
                        (os.getenv("QUEUE"), mensaje[0])
                    )
                    print(f"Enviando a Temporal: {res}")
                    async with aiohttp.ClientSession() as session:
                        async with session.post((os.getenv("TEMPORAL_URL")), json=res["returning"][0]) as response:
                            response_data = await response.json()
                            print(response_data)


   
        
            
            
            
            
