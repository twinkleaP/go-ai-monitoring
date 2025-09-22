import numpy as np
from sklearn.ensemble import IsolationForest

class AnomalyModel:
    def __init__(self):
        # For demo: train on synthetic "normal" data
        normal_data = np.random.normal(50, 10, (100, 3))  # fake baseline
        self.model = IsolationForest(contamination=0.1)
        self.model.fit(normal_data)

    def predict(self, X):
        return self.model.predict(X)  # -1 = anomaly, 1 = normal
