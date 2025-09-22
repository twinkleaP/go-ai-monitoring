from fastapi import FastAPI
from pydantic import BaseModel
from app.model import AnomalyModel

app = FastAPI(title="AI Anomaly Detection Service")

# Request schema
class Metrics(BaseModel):
    cpu: float
    memory: float
    disk: float

# Load model
model = AnomalyModel()

@app.post("/predict")
def predict(metrics: Metrics):
    data = [[metrics.cpu, metrics.memory, metrics.disk]]
    result = model.predict(data)
    return {"prediction": bool(result[0])}
