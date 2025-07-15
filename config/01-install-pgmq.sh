#!/bin/bash

# Instalar dependencias necesarias
apt-get update
apt-get install -y build-essential git postgresql-server-dev-15

# Clonar y compilar pgmq
cd /tmp
git clone https://github.com/queueup/pgmq.git
cd pgmq
make
make install

# Limpiar archivos temporales
rm -rf /tmp/pgmq

# Crear la extensión pgmq en la base de datos
psql -U postgres -d BPMN_db -c "CREATE EXTENSION IF NOT EXISTS pgmq;"

# Crear la cola bpmn_queue
psql -U postgres -d BPMN_db -c "SELECT pgmq.create('bpmn_queue');"

psql -U postgres -d BPMN_db -c "
CREATE TYPE request_status AS ENUM ('pending', 'approved', 'rejected');
CREATE TABLE IF NOT EXISTS request (
    id SERIAL PRIMARY KEY,
    status request_status DEFAULT 'pending',
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
"

psql -U postgres -d BPMN_db -c "CREATE INDEX IF NOT EXISTS idx_request_process_id ON request(process_id);"
psql -U postgres -d BPMN_db -c "CREATE INDEX IF NOT EXISTS idx_request_status ON request(status);"
psql -U postgres -d BPMN_db -c "CREATE INDEX IF NOT EXISTS idx_request_created_at ON request(created_at);"

psql -U postgres -d BPMN_db -c "
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS \$\$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
\$\$ language 'plpgsql';
"

psql -U postgres -d BPMN_db -c "DROP TRIGGER IF EXISTS update_request_updated_at ON request;"
psql -U postgres -d BPMN_db -c "
CREATE TRIGGER update_request_updated_at
    BEFORE UPDATE ON request
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();
"

echo "pgmq instalado y configurado correctamente" 