import os
from gql import gql, Client
from gql.transport.requests import RequestsHTTPTransport


url = os.getenv("ENDPOINT_GRAPHQL")
transport = RequestsHTTPTransport(url=url, verify=True)

client = Client(transport=transport, fetch_schema_from_transport=False)

