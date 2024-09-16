CREATE TABLE tender (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL,                  
    description TEXT,                           
    service_type VARCHAR(50) NOT NULL,             
    status VARCHAR(20) NOT NULL,                 
    organization_id UUID NOT NULL REFERENCES organization(id) ON DELETE CASCADE, 
    creator_id UUID NOT NULL REFERENCES employee(id) ON DELETE SET NULL,       
    version INT DEFAULT 1,                        
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  
    creator_username VARCHAR(50),                
    CONSTRAINT chk_status CHECK (status IN ('CREATED', 'PUBLISHED', 'CLOSED'))
);