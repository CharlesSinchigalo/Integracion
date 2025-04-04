import sqlite3

# Conexión a la base de datos
conn = sqlite3.connect("database.db")
cursor = conn.cursor()

# Crear tabla de estudiantes
cursor.execute('''
CREATE TABLE IF NOT EXISTS estudiantes (
    ci TEXT PRIMARY KEY,
    nombres TEXT NOT NULL
)
''')

# Crear tabla de asignaturas
cursor.execute('''
CREATE TABLE IF NOT EXISTS asignaturas (
    id INTEGER PRIMARY KEY,
    nombre TEXT NOT NULL,
    nivel INTEGER NOT NULL
)
''')

# Confirmar cambios y cerrar conexión
conn.commit()
conn.close()

print("✅ Tablas creadas correctamente en database.db")
