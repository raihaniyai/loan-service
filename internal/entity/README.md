Endpoints:
- Create Loan
POST /loans
```
curl --location 'http://localhost:8080/loans' \
--header 'Authorization: Bearer 2' \
--header 'Content-Type: application/json' \
--data '{
    "principal_amount": 10000000,
    "interest_rate": 0.1,
    "return_on_investment": 0.05
}'
```

- Invest Loan
- Get Loans

Assumptions:
1. one user only has one role (e.g. admin, borrower or investor)
2. one borrower can only have one loan at a time

PostgreSQL DDL

```
CREATE TABLE loans (
    id BIGSERIAL PRIMARY KEY,
    borrower_id BIGINT NOT NULL,
    principal_amount BIGINT NOT NULL,
    interest_rate REAL NOT NULL,        -- Stored as percentage
    return_on_investment REAL NOT NULL, -- Stored as percentage
    agreement_letter TEXT,
    status INT DEFAULT 10,
    updated_by BIGINT,                  -- Stores user ID, assuming from external service
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);

-- Indexes for loans table
CREATE INDEX idx_loans_borrower_id ON loans(borrower_id);
CREATE INDEX idx_loans_status ON loans(status);
CREATE INDEX idx_loans_updated_by ON loans(updated_by);
CREATE INDEX idx_loans_created_at ON loans(created_at);

CREATE TABLE actions (
    action_id BIGSERIAL PRIMARY KEY,
    loan_id BIGINT NOT NULL REFERENCES loans(id) ON DELETE CASCADE,
    action_type INT NOT NULL,           -- Enum or predefined action types (e.g., approval, disbursement)
    action_by INT NOT NULL,
    document TEXT,
    created_by BIGINT NOT NULL,         -- Stores user ID, assuming from external service
    created_at TIMESTAMP NOT NULL
);

-- Indexes for actions table
CREATE INDEX idx_actions_loan_id ON actions(loan_id);
CREATE INDEX idx_actions_action_type ON actions(action_type);
CREATE INDEX idx_actions_action_by ON actions(action_by);
CREATE INDEX idx_actions_created_at ON actions(created_at);

CREATE TABLE investments (
    investment_id BIGSERIAL PRIMARY KEY,
    loan_id BIGINT NOT NULL REFERENCES loans(id) ON DELETE CASCADE,
    investor_id BIGINT NOT NULL,        -- Stores user ID, assuming from external service
    investment_amount BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL
);

-- Indexes for investments table
CREATE INDEX idx_investments_loan_id ON investments(loan_id);
CREATE INDEX idx_investments_investor_id ON investments(investor_id);
CREATE INDEX idx_investments_created_at ON investments(created_at);

CREATE TABLE users (
    user_id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255),
    role INT NOT NULL,                  -- Enum or predefined roles (e.g., admin=1, borrower=2, investor=3)
    email VARCHAR(255) UNIQUE,
    phone_number VARCHAR(20) UNIQUE
);

-- Indexes for users table
CREATE INDEX idx_users_role ON users(role);
```

DML users (data are dummy)
```
-- Insert 10 sample rows into the users table with Indonesian names and phone numbers (no dashes)
INSERT INTO users (name, role, email, phone_number)
VALUES 
  ('Andi Pratama', 1, 'andi.pratama@example.com', '081234567890'),
  ('Budi Santoso', 2, 'budi.santoso@example.com', '082123456789'),
  ('Citra Dewi', 3, 'citra.dewi@example.com', '083123456789'),
  ('Dina Sari', 2, 'dina.sari@example.com', '084234567890'),
  ('Eka Putri', 1, 'eka.putri@example.com', '085345678901'),
  ('Fajar Hidayat', 3, 'fajar.hidayat@example.com', '085656789012'),
  ('Gina Rahayu', 2, 'gina.rahayu@example.com', '087767890123'),
  ('Hendra Wijaya', 1, 'hendra.wijaya@example.com', '089878901234'),
  ('Ika Lestari', 3, 'ika.lestari@example.com', '081789012345'),
  ('Joko Susilo', 2, 'joko.susilo@example.com', '082390123456');
```