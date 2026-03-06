CREATE TABLE appointments (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    patient_id UUID NOT NULL,
    doctor_id UUID NOT NULL,
    appointment_date TIMESTAMP NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'PENDING',
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    

    CONSTRAINT fk_patient FOREIGN KEY (patient_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_doctor FOREIGN KEY (doctor_id) REFERENCES users(id) ON DELETE CASCADE
);


CREATE INDEX idx_appointments_patient ON appointments(patient_id);
CREATE INDEX idx_appointments_doctor ON appointments(doctor_id);