import numpy as np
from tensorflow.keras.models import load_model

def train(model_path, dataset_path, epochs, batch_size, optimizer, learning_rate):
    model = load_model(model_path)

    model.compile(optimizer=optimizer, loss='categorical_crossentropy', metrics=['accuracy'])

    # datasert upl
    data = np.load(dataset_path)
    x_train = data['x_train']
    y_train = data['y_train']

    print(f"Обучение начато... {epochs} эпох")
    
    history = model.fit(x_train, y_train, 
                        epochs=epochs, 
                        batch_size=batch_size)

    model.save(model_path)

    result = {
        "model": os.path.basename(model_path),
        "dataset": os.path.basename(dataset_path),
        "epochs": epochs,
        "batch_size": batch_size,
        "optimizer": optimizer,
        "learning_rate": learning_rate,
        "final_loss": float(history.history['loss'][-1]),
        "final_accuracy": float(history.history['accuracy'][-1]),
        "status": "completed",
        "timestamp": int(time.time())
    }

    print(f"RESULT:{json.dumps(result)}")
    return result