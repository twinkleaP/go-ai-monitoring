from fastapi import FastAPI
from pydantic import BaseModel
from model import AnomalyModel
#from app.model import AnomalyModel
import numpy as np
from sklearn.ensemble import IsolationForest



app = FastAPI()

# Request schema
class Metrics(BaseModel):
    cpu: float
    memory: float
    disk: float

# Load model
model = AnomalyModel()

@app.get("/health")
def health():
    return {"status": "ok"}

@app.post("/predict")
def predict(metrics: Metrics):
    data = [[metrics.cpu, metrics.memory, metrics.disk]]
    result = model.predict(data)
    return {"prediction": bool(result[0])}
