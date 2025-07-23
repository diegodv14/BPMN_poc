-- Script de inicialización de la base de datos BPMN
-- Este script se ejecuta automáticamente al crear la base de datos por primera vez

-- Crear la extensión pgmq
CREATE EXTENSION IF NOT EXISTS pgmq;

CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- Crear la cola bpmn_queue
SELECT pgmq.create('bpmn_queue');

CREATE TYPE request_status AS ENUM ('pending', 'approved', 'rejected');

-- Crear la tabla request
CREATE TABLE IF NOT EXISTS request (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    status request_status DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Función para actualizar el timestamp de updated_at
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

-- Trigger para actualizar automáticamente updated_at
DROP TRIGGER IF EXISTS update_request_updated_at ON request;
CREATE TRIGGER update_request_updated_at
    BEFORE UPDATE ON request
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column(); 