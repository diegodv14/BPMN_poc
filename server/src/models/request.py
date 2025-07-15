from datetime import datetime
from enum import Enum
from pydantic import BaseModel, Field
from uuid import uuid4

enum_status = Enum("status", ["pending", "in_progress", "completed", "failed"])

class Request(BaseModel):
    id: str = Field(default=str(uuid4()))
    name: str = Field(default="")
    description: str = Field(default="")
    status: enum_status = Field(default=enum_status.pending)
    created_at: datetime = Field(default=datetime.now())
    updated_at: datetime = Field(default=datetime.now())
    
    
    
    
    
    