from fastapi import APIRouter, Depends, HTTPException
from sqlalchemy.orm import Session
from database import SessionLocal
from models import Estudiante
from schemas import EstudianteCreate

router = APIRouter(prefix="/estudiantes", tags=["Estudiantes"])

# Dependencia para obtener la DB
def get_db():
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()

# Obtener todos los estudiantes
@router.get("/")
def listar_estudiantes(db: Session = Depends(get_db)):
    return db.query(Estudiante).all()

# Crear un nuevo estudiante
@router.post("/")
def crear_estudiante(est: EstudianteCreate, db: Session = Depends(get_db)):
    db_est = Estudiante(ci=est.ci, nombres=est.nombres)
    db.add(db_est)
    db.commit()
    db.refresh(db_est)
    return db_est
