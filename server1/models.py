from pydantic import BaseModel

class Estudiante(BaseModel):
    ci: str
    nombres: str

class Asignatura(BaseModel):
    id: int
    nombre: str
    nivel: int
