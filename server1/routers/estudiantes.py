from fastapi import APIRouter
from models import Estudiante
from database import get_connection

router = APIRouter()

@router.post("/estudiantes")
def registrar_estudiante(est: Estudiante):
    conn = get_connection()
    cursor = conn.cursor()
    cursor.execute("INSERT INTO estudiantes (ci, nombres) VALUES (?, ?)", (est.ci, est.nombres))
    conn.commit()
    conn.close()
    return {"mensaje": "Estudiante registrado"}


@router.get("/estudiantes")
def obtener_estudiantes():
    conn = get_connection()
    cursor = conn.cursor()
    cursor.execute("SELECT * FROM estudiantes")
    estudiantes = cursor.fetchall()
    conn.close()
    return [dict(row) for row in estudiantes]
