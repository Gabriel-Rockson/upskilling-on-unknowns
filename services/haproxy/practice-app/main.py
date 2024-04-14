from fastapi import FastAPI
import os

app = FastAPI()

port = int(os.getenv("PRACTICE_APP_PORT", 8001))
current_app = str(os.getenv("CURRENT_APP", "default_app"))


def get_current_running_app(): 
    return f"Request served from {current_app}"

@app.get("/")
async def home():
    return f"This is the homepage. {get_current_running_app()}"


@app.get("/products")
async def products(): 
    return f"This is products route. {get_current_running_app()}"

@app.get("/cart")
async def cart(): 
    return f"This is cart route {get_current_running_app()}"

@app.get("/payments")
async def payments(): 
    return f"This is payments route. {get_current_running_app()}"

@app.get("/admin")
async def admin():
    return f"This is the ADMIN, not everyone should have access. {get_current_running_app()}"


if __name__ == "__main__":
    import uvicorn

    uvicorn.run(app, host='0.0.0.0', port=port)
