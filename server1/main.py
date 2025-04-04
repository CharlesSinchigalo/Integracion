from fastapi import FastAPI
from routers import estudiantes, asignaturas

app = FastAPI(
    title="API Estudiantes - Servidor 1",
    description="Maneja estudiantes y asignaturas con SQLite",
    version="1.0.0"
)

# Incluir las rutas
app.include_router(estudiantes.router, prefix="/api", tags=["Estudiantes"])
app.include_router(asignaturas.router, prefix="/api", tags=["Asignaturas"])
