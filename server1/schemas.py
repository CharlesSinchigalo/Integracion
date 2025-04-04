from pydantic import BaseModel

class EstudianteCreate(BaseModel):
    ci: str
    nombres: str
