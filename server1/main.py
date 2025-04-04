from fastapi import FastAPI
from database import Base, engine
from routers import estudiantes

app = FastAPI()

# Crear tablas
Base.metadata.create_all(bind=engine)

# Incluir el router de estudiantes
app.include_router(estudiantes.router)
