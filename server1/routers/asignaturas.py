from fastapi import APIRouter # type: ignore
from models import Asignatura
from database import get_connection

router = APIRouter()

@router.post("/asignaturas")
def registrar_asignatura(asig: Asignatura):
    conn = get_connection()
    cursor = conn.cursor()
    cursor.execute("INSERT INTO asignaturas (id, nombre, nivel) VALUES (?, ?, ?)", (asig.id, asig.nombre, asig.nivel))
    conn.commit()
    conn.close()
    return {"mensaje": "Asignatura registrada"}
