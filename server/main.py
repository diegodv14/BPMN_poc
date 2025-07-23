from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
from contextlib import asynccontextmanager
import uvicorn
import os
from apscheduler.schedulers.background import BackgroundScheduler
from apscheduler.triggers.interval import IntervalTrigger
from dotenv import load_dotenv
from src.job.process import start_async
from src.routes import router

load_dotenv()

def job():
    print("Job iniciando")
    start_async()

executors = {
    'default': {'type': 'threadpool', 'max_workers': 1}
}
job_defaults = {
    'coalesce': False,
    'max_instances': 1
}
scheduler = BackgroundScheduler(executors=executors, job_defaults=job_defaults)
scheduler.add_job(job, IntervalTrigger(seconds=5), id="job")


@asynccontextmanager
async def lifespan(app: FastAPI):
    scheduler.start()
    print("Scheduler iniciado")
    yield
    scheduler.shutdown()
    print("Scheduler detenido")


app = FastAPI(
    title="BPMN Server API",
    description="API del servidor BPMN con FastAPI",
    version="1.0.0",
    docs_url="/",
    redoc_url="/redoc",
    openapi_url="/openapi.json",
    lifespan=lifespan
)

app.include_router(router)

app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

if __name__ == "__main__":
    port = int(os.getenv("PORT", 8000))
    uvicorn.run(
        "main:app",
        host="0.0.0.0",
        port=port,
        reload=True,
        log_level="info"
    )
