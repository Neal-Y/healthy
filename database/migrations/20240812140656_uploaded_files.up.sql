CREATE TABLE uploaded_files (
    id SERIAL PRIMARY KEY,
    file_path TEXT NOT NULL,
    upload_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    accessed_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
