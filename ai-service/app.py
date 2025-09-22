from fastapi import FastAPI
from pydantic import BaseModel
import numpy as np
from sklearn.ensemble import IsolationForest

app = FastAPI()

# Example model trained on dummy data
model = IsolationForest(contamination=0.1, random_state=42)
model.fit(np.random.normal(50, 10, (100, 3)))  # CPU/Memory/Disk

class Metrics(BaseModel):
    cpu: float
    memory: float
    disk: float

@app.post("/predict")
def predict(m: Metrics):
    X = np.array([[m.cpu, m.memory, m.disk]])
    pred = model.predict(X)[0]       # -1 = anomaly, 1 = normal
    # Convert to Python bool
    is_anomaly = True if pred == -1 else False
    return {"anomaly": is_anomaly}
