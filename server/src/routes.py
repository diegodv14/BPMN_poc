from fastapi import APIRouter, HTTPException
from typing import Dict, Any

router = APIRouter()

@router.get("/", 
    summary="Obtener API",
    description="Retorna la API",
    tags=["BPMN"],
    response_model=Dict[str, Any])
async def get_api():
    return {
        "success": True,
        "message": "API obtenida exitosamente",
        "data": "API"
    }