from datetime import datetime
from enum import Enum
from pydantic import BaseModel, Field
from uuid import uuid4

class Status(Enum):
    pending = "pending"
    in_progress = "in_progress"
    completed = "completed"
    failed = "failed"

class Request(BaseModel):
    id: str = Field(default=str(uuid4()))
    name: str = Field(default="")
    description: str = Field(default="")
    status: Status = Field(default=Status.pending.value)
    created_at: datetime = Field(default=datetime.now())
    updated_at: datetime = Field(default=datetime.now())
    
    
    
    
    
    