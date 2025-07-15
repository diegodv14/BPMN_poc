from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
import uvicorn
import os
from apscheduler.schedulers.background import BackgroundScheduler
from apscheduler.triggers.interval import IntervalTrigger
from dotenv import load_dotenv
from src.job.process import start_async
from src.routes import router

load_dotenv()

app = FastAPI(
    title="BPMN Server API",
    description="API del servidor BPMN con FastAPI",
    version="1.0.0",
    docs_url="/",
    redoc_url="/redoc",
    openapi_url="/openapi.json"
)

app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)


def job():
    print("Job iniciando")
    start_async()



if __name__ == "__main__":
    
    executors = {
        'default': {'type': 'threadpool', 'max_workers': 1 }
    }

    job_defaults = {
        'coalesce': False,
        'max_instances': 1
    }
    scheduler = BackgroundScheduler(executors=executors, job_defaults=job_defaults)
    scheduler.add_job(job, IntervalTrigger(seconds=5), id="job")
    
    port = int(os.getenv("PORT", 8000))
    uvicorn.run(
        "main:app",
        host="0.0.0.0",
        port=port,
        reload=True,
        log_level="info"
    )
